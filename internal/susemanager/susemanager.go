// Package susemanager - SUSE Manager api call and support functions
package susemanager

import (
	"encoding/json"
	"fmt"
	"strings"

	sumamodels "SUSE-Manager-Tools-V2/internal/models/susemanager"
	util "SUSE-Manager-Tools-V2/internal/util/contains"
	"SUSE-Manager-Tools-V2/internal/util/rest"
	returnCodes "SUSE-Manager-Tools-V2/internal/util/returnCodes"

	"github.com/pkg/errors"

	"go.uber.org/zap"
)

// SuseManager - general information
type SuseManager struct {
	proxy  IProxy
	cfg    *SumanConfig
	logger *zap.Logger
}

// NewSuseManager - open SUSE Manager
//
// param: proxy
// param: cfg
// param: logger
// return:
func NewSuseManager(proxy IProxy, cfg *SumanConfig, logger *zap.Logger) ISuseManager {
	return &SuseManager{
		proxy:  proxy,
		cfg:    cfg,
		logger: logger,
	}
}

// GetSystemGroupName - generate the systemgroup name to be used for a POD
//
// param: negName
// return:
func (s *SuseManager) GetSystemGroupName(negName string) string {
	location := strings.ToLower(strings.ReplaceAll(negName, "/", "-"))
	return "a4-loc-" + location
}

// SetK3sDetails - set k3s formular
//
// param: auth
// param: systemgroupName
// param: k3sconfiginput
func (s *SuseManager) SetK3sDetails(auth AuthParams, systemgroupName string, k3sconfiginput map[string]interface{}) error {
	var formulaName = "dtag-k3s-pod"
	var section = "k3sconfig"
	returnValue := setformulaatgrouplevel(s, auth, systemgroupName, formulaName, section, k3sconfiginput)
	if returnValue != nil {
		return returnValue
	}
	return nil
}

// setformulaatgrouplevel - at formular to systemgroup
//
// param: s
// param: auth
// param: systemgroupName
// param: formulaName
// param: section
// param: input
func setformulaatgrouplevel(s *SuseManager, auth AuthParams, systemgroupName string, formulaName string, section string, input map[string]interface{}) error {
	data, err := s.proxy.SystemGroupGetDetails(auth, systemgroupName)
	if err != nil {
		return err
	}
	k3full, err := s.proxy.GetGroupFormulaData(auth, data.ID, formulaName)
	if err != nil {
		return err
	}
	k3sdata := k3full.(map[string]interface{})

	for item, k3sinner := range k3sdata {
		if item == section {
			for inputkey, inputdata := range input {
				k3sinner.(map[string]interface{})[inputkey] = inputdata
			}
			break
		}
	}
	success, err := s.proxy.SetGroupFormulaData(auth, data.ID, formulaName, k3sdata)
	if err != nil {
		return err
	}

	if success != 1 {
		return fmt.Errorf(fmt.Sprintf("not able to update %s data in susemanager %s formula at system group %s", section, formulaName, systemgroupName))
	}
	return nil
}

// ChangeChannels - change assigned channels to system
//
// param: auth
// param: systemID
// param: targetedVersion
func (s *SuseManager) ChangeChannels(auth AuthParams, systemID int, targetedVersion string) error {
	channels, err := s.proxy.ChannelListSoftwareChannels(auth)
	if err != nil {
		return err
	}
	baseChannelLabel := ""
	for i := range channels {
		if strings.Contains(channels[i].ParentLabel, targetedVersion) {
			baseChannelLabel = channels[i].ParentLabel
			break
		}
	}
	if baseChannelLabel == "" {
		return fmt.Errorf("base channel for targetd version %s not available", targetedVersion)
	}
	childChannels, err := s.proxy.ChannelSoftwareListChildren(auth, baseChannelLabel)
	if err != nil {
		return err
	}
	err = s.proxy.SystemScheduleChangeChannels(auth, systemID, baseChannelLabel, childChannels)
	if err != nil {
		return fmt.Errorf("error while updating the channels")
	}
	s.logger.Info("Channel change is completed", zap.Any("systemID", systemID))

	return nil
}

// InstallPackages - install the mentoined packages to a system
//
// param: auth
// param: systemID
// param: pkgs
// param: timeout
func (s *SuseManager) InstallPackages(auth AuthParams, systemID int, pkgs []string, timeout int) error {
	s.logger.Debug("Inside Install pkgs function", zap.Any("pkgs", pkgs))
	// Getting Installed Pkgs
	installedPkgs, err := s.proxy.SystemListInstalledPackages(auth, systemID)
	if err != nil {
		return err
	}
	var installedPkgsName []string
	for i := range installedPkgs {
		installedPkgsName = append(installedPkgsName, installedPkgs[i].Name)
	}
	// Getting list of pkgs that are to be installed
	var installpkg []string
	for i := range pkgs {
		if !util.Contains(installedPkgsName, pkgs[i]) {
			installpkg = append(installpkg, pkgs[i])
		}
	}
	// Getting list of installable pkgs
	installablePkgs, err := s.proxy.ListLatestInstallablePackages(auth, systemID)
	if err != nil {
		return err
	}
	var installablePkgsName []string
	for i := range installablePkgs {
		installablePkgsName = append(installablePkgsName, installablePkgs[i].Name)
	}

	// Getting list of pkgs that are actually can be installed
	for i := range installpkg {
		if !util.Contains(installablePkgsName, installpkg[i]) {
			return fmt.Errorf("%s package is not present to intall", installpkg[i])
		}
	}

	// run script to install pkgs
	installpkgs := strings.Join(installpkg, " ")
	script := "#!/bin/bash\ntransactional-update -c -n pkg install " + installpkgs
	err = s.proxy.ScheduleScriptRun(auth, systemID, timeout, script)
	if err != nil {
		return err
	}

	return nil
}

// HandleSuseManagerResponse - handle API response
//
// param: body
// return:
func HandleSuseManagerResponse(body []byte) (interface{}, error) {
	var resp sumamodels.RespAPI
	err := json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	if resp.Success {
		return resp.Result, nil
	}
	if resp.Message != "" {
		return false, errors.New(resp.Message)
	}
	if resp.Messages != nil && len(resp.Messages) > 0 {
		return false, errors.New(resp.Messages[0])
	}
	return false, errors.New("Unexpected error")
}

// AuthParams - authentication key
type AuthParams struct {
	SessionKey string
	Host       string
}

// GetAuth - get authentication key
//
// param: sessionKey
// return:
func (s *SuseManager) GetAuth(sessionKey string) (*AuthParams, error) {
	auth := AuthParams{
		SessionKey: sessionKey,
		Host:       s.cfg.Host,
	}
	return &auth, nil
}

// GetHost - get host
//
// param: negName
// param: sessionKey
// return:
func (s *SuseManager) GetHost(negName string, sessionKey string) (*AuthParams, error) {

	auth := AuthParams{
		SessionKey: sessionKey,
		Host:       s.cfg.Host,
	}
	// List system based on group
	podMembers, err := s.proxy.SystemGroupListSystemsMinimal(auth, s.GetSystemGroupName(negName))
	if err != nil {
		return nil, err
	}

	// Get slaves session key
	if podMembers == nil || err != nil {
		slaves, err := s.proxy.GetSlaves(auth.SessionKey)
		if err != nil {
			return nil, err
		}
		for _, slave := range slaves {
			system, err := s.proxy.SystemGetID(auth, slave.Label)
			if err != nil {
				s.logger.Error("Error while getting systemID", zap.Any("error", err.Error()))
				return nil, err
			}
			resp, err := s.proxy.GetSystemFormulaData(auth, system[0].ID, "uyunihub")
			if err != nil {
				s.logger.Error("Error while unmarshaling data from SystemFormula", zap.Any("error", err.Error()))
				return nil, err
			}
			var slaveformuladata sumamodels.Uyunihub
			byteArray, _ := json.Marshal(resp)
			err = json.Unmarshal(byteArray, &slaveformuladata)
			if err != nil {
				s.logger.Error("Error while unmarshaling data from Uyunihub", zap.Any("error", err.Error()))
				return nil, err
			}
			reqBody, _ := json.Marshal(map[string]interface{}{
				"login":    slaveformuladata.Hub.ServerUserName,
				"password": slaveformuladata.Hub.ServerPassword})
			slavesessionKey, err := s.proxy.GetSessionKey(reqBody, slave.Label)
			if err != nil {
				s.logger.Error("Error while login to suse slave", zap.Any("error", err.Error()))
			}
			auth = AuthParams{
				SessionKey: slavesessionKey,
				Host:       slave.Label,
			}
			podMembers, err = s.proxy.SystemGroupListSystemsMinimal(auth, s.GetSystemGroupName(negName))
			if podMembers != nil {
				return &auth, nil
			}
			if err != nil || podMembers == nil {
				err := s.proxy.SumanLogout(auth)
				if err != nil {
					s.logger.Error("Error while perform logout against SUSE Manager", zap.Any("error", err.Error()))
					return nil, err
				}
				continue
			}
		}
	} else {
		auth := AuthParams{
			SessionKey: auth.SessionKey,
			Host:       s.cfg.Host,
		}
		return &auth, nil
	}
	return nil, nil
}

// CheckResponseProgress - check response from api call
//
// param: auth
// param: response
// param: timeOut
// param: systemID
// param: funcName
func (p *Proxy) CheckResponseProgress(auth AuthParams, response *rest.HTTPHelperStruct, timeOut int, systemID int, funcName string) error {
	var actionID int
	if response.StatusCode == 200 {
		resp, err := HandleSuseManagerResponse(response.Body)
		if err != nil {
			p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrHandlingSuseManagerResponse, err))
			return fmt.Errorf(returnCodes.ErrHandlingSuseManagerResponse)
		}
		byteArray, err := json.Marshal(resp)
		if err != nil {
			p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedMarshalling, err))
			return fmt.Errorf(returnCodes.ErrFailedMarshalling)
		}
		err = json.Unmarshal(byteArray, &actionID)
		if err != nil {
			p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedUnMarshalling, err))
			return fmt.Errorf(returnCodes.ErrFailedUnMarshalling)
		}
		_, err = p.CheckProgress(auth, actionID, timeOut, funcName, systemID)
		if err != nil {
			return err
		}
	} else {
		p.logger.Error(fmt.Sprintf("running %v Failed. Http StatusCode: %v Http Body: %v", funcName, response.StatusCode, response.Body))
		return fmt.Errorf(returnCodes.ErrProcessingData)
	}
	return nil
}
