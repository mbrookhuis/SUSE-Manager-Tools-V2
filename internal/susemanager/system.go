// Package susemanager api call for SUSE Manager related to system
package susemanager

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	sumamodels "ecp-golang-cm/pkg/models/susemanager"
	returnCodes "ecp-golang-cm/pkg/util/returnCodes"

	"go.uber.org/zap"
)

// SystemGetID - get systemID from the given server
//
// param: requestID
// param: auth
// param: systemName
// return: []sumamodels.System, error
func (p *Proxy) SystemGetID(requestID string, auth AuthParams, systemName string) ([]sumamodels.System, error) {
	var systeminfo []sumamodels.System
	body, err := json.Marshal(map[string]string{"name": systemName})
	if err != nil {
		p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedMarshalling, err), zap.Any("requestID", requestID))
		return nil, fmt.Errorf(returnCodes.ErrFailedMarshalling)
	}
	path := "system/getId"
	response, err := p.suse.SuseManagerCall(body, http.MethodGet, auth.Host, path, auth.SessionKey)
	if err != nil {
		p.logger.Error(fmt.Sprintf("unable to get system info: %s", response.Body))
		return nil, errors.New(returnCodes.ErrSystemNotFound)
	}
	if response.StatusCode == 200 {
		resp, err := HandleSuseManagerResponse(response.Body)
		if err != nil {
			p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrHandlingSuseManagerResponse, err), zap.Any("requestID", requestID))
			return nil, fmt.Errorf(returnCodes.ErrHandlingSuseManagerResponse)
		}
		byteArray, err := json.Marshal(resp)
		if err != nil {
			p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedMarshalling, err), zap.Any("requestID", requestID))
			return nil, fmt.Errorf(returnCodes.ErrFailedMarshalling)
		}
		err = json.Unmarshal(byteArray, &systeminfo)
		if err != nil {
			p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedUnMarshalling, err), zap.Any("requestID", requestID))
			return nil, fmt.Errorf(returnCodes.ErrFailedUnMarshalling)
		}
	} else {
		p.logger.Error(fmt.Sprintf("unable to get system info: %s", response.Body))
		return nil, errors.New(returnCodes.ErrSystemNotFound)
	}
	return systeminfo, nil
}

// SchedulePackageRefresh - schedule get a list of installed packages
//
// param: requestID
// param: auth
// param: systemID
func (p *Proxy) SchedulePackageRefresh(requestID string, auth AuthParams, systemID int) error {
	p.logger.Debug("Scheduling package refresh called", zap.Any("requestID", requestID))
	body, err := json.Marshal(map[string]interface{}{"sid": systemID, "earliestOccurrence": time.Now()})
	if err != nil {
		p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedMarshalling, err), zap.Any("requestID", requestID))
		return fmt.Errorf(returnCodes.ErrFailedMarshalling)
	}
	path := "system/schedulePackageRefresh"
	response, err := p.suse.SuseManagerCall(body, http.MethodPost, auth.Host, path, auth.SessionKey)
	if err != nil {
		p.logger.Error(fmt.Sprintf("error while scheduling package refresh err: %s", response.Body))
		return errors.New(returnCodes.ErrProcessingData)
	}
	return p.CheckResponseProgress(requestID, auth, response, 12000, systemID, "SchedulePackageRefresh")
}

// ScheduleScriptRun - schedule running a script on the given server
//
// param: requestID
// param: auth
// param: systemID
// param: timeout
// param: script
func (p *Proxy) ScheduleScriptRun(requestID string, auth AuthParams, systemID int, timeout int, script string) error {
	p.logger.Info("script run scheduled", zap.Any("script", script), zap.Any("requestID", requestID), zap.Any("systemID", systemID))
	body, err := json.Marshal(map[string]interface{}{"sid": systemID,
		"username":           "root",
		"groupname":          "root",
		"timeout":            timeout,
		"script":             script,
		"earliestOccurrence": time.Now()})
	if err != nil {
		p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedMarshalling, err), zap.Any("requestID", requestID))
		return fmt.Errorf(returnCodes.ErrFailedMarshalling)
	}
	path := "system/scheduleScriptRun"
	response, err := p.suse.SuseManagerCall(body, http.MethodPost, auth.Host, path, auth.SessionKey)
	if err != nil {
		p.logger.Error("Error message recieved from suse-manger", zap.Any("suma-host", auth.Host), zap.Any("requestID", requestID), zap.Any("error", err.Error()))
		return fmt.Errorf(returnCodes.ErrProcessingData)
	}
	return p.CheckResponseProgress(requestID, auth, response, timeout, systemID, "ScheduleScriptRun")
}

// SystemGetScriptResult - get results from script run
//
// param: requestID
// param: auth
// param: actionID
// param: resultCompleted
// return: string, error
func (p *Proxy) SystemGetScriptResult(requestID string, auth AuthParams, actionID int, resultCompleted int) (string, error) {
	body, err := json.Marshal(map[string]interface{}{"actionId": actionID})
	if err != nil {
		p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedMarshalling, err), zap.Any("requestID", requestID))
		return "", fmt.Errorf(returnCodes.ErrFailedMarshalling)
	}
	path := "system/getScriptResults"
	response, err := p.suse.SuseManagerCall(body, http.MethodPost, auth.Host, path, auth.SessionKey)
	if err != nil {
		p.logger.Error("Error message recieved from suse-manger", zap.Any("requestID", requestID), zap.Any("error", err.Error()))
		return "", fmt.Errorf(returnCodes.ErrProcessingData)
	}
	var scriptResults []sumamodels.ScriptResult
	var output string
	if response.StatusCode == 200 {
		resp, err := HandleSuseManagerResponse(response.Body)
		if err != nil {
			p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrHandlingSuseManagerResponse, err), zap.Any("requestID", requestID))
			return "", fmt.Errorf(returnCodes.ErrHandlingSuseManagerResponse)
		}
		byteArray, err := json.Marshal(resp)
		if err != nil {
			p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedMarshalling, err), zap.Any("requestID", requestID))
			return "", fmt.Errorf(returnCodes.ErrFailedMarshalling)
		}
		err = json.Unmarshal(byteArray, &scriptResults)
		if err != nil {
			p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedUnMarshalling, err), zap.Any("requestID", requestID))
			return "", fmt.Errorf(returnCodes.ErrFailedUnMarshalling)
		}
		for i := range scriptResults {
			output = scriptResults[i].Output
		}
		if resultCompleted != 1 {
			p.logger.Error("Error from script run", zap.Any("requestID", requestID), zap.Any("Script message", output))
			return "", fmt.Errorf(returnCodes.ErrProcessingData)
		}
	}
	return output, nil
}

// SystemScheduleReboot - reboot the given system
//
// param: requestID
// param: auth
// param: systemID
// param: timeout
func (p *Proxy) SystemScheduleReboot(requestID string, auth AuthParams, systemID int, timeout int) error {
	body, err := json.Marshal(map[string]interface{}{"sid": systemID, "earliestOccurrence": time.Now()})
	if err != nil {
		p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedMarshalling, err), zap.Any("requestID", requestID))
		return fmt.Errorf(returnCodes.ErrFailedMarshalling)
	}
	path := "system/scheduleReboot"
	response, err := p.suse.SuseManagerCall(body, http.MethodPost, auth.Host, path, auth.SessionKey)
	if err != nil {
		p.logger.Error("Error message recieved from suse-manger", zap.Any("requestID", requestID), zap.Any("error", err.Error()))
		return fmt.Errorf(returnCodes.ErrProcessingData)
	}
	return p.CheckResponseProgress(requestID, auth, response, timeout, systemID, "SystemScheduleReboot")
}

// ListInprogressSystem - list systems with running actions started from SUSE Manager
//
// param: requestID
// param: auth
// param: actionID
// return:
func (p *Proxy) ListInprogressSystem(requestID string, auth AuthParams, actionID int) ([]interface{}, error) {
	body, err := json.Marshal(map[string]interface{}{"actionId": actionID})
	if err != nil {
		p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedMarshalling, err), zap.Any("requestID", requestID))
		return nil, fmt.Errorf(returnCodes.ErrFailedMarshalling)
	}
	path := "schedule/listInProgressSystems"
	response, err := p.suse.SuseManagerCall(body, http.MethodGet, auth.Host, path, auth.SessionKey)
	if err != nil {
		p.logger.Error("Error message recieved from suse-manger", zap.Any("requestID", requestID), zap.Any("error", err.Error()))
		return nil, fmt.Errorf(returnCodes.ErrProcessingData)
	}
	var resultSuc []interface{}
	if response.StatusCode == 200 {
		resp, err := HandleSuseManagerResponse(response.Body)
		if err != nil {
			p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrHandlingSuseManagerResponse, err), zap.Any("requestID", requestID))
			return nil, fmt.Errorf(returnCodes.ErrHandlingSuseManagerResponse)
		}
		byteArray, err := json.Marshal(resp)
		if err != nil {
			p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedMarshalling, err), zap.Any("requestID", requestID))
			return nil, fmt.Errorf(returnCodes.ErrFailedMarshalling)
		}
		err = json.Unmarshal(byteArray, &resultSuc)
		if err != nil {
			p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedUnMarshalling, err), zap.Any("requestID", requestID))
			return nil, fmt.Errorf(returnCodes.ErrFailedUnMarshalling)
		}
	}
	p.logger.Debug("Response from api", zap.Any("api", "schedule/listInProgressSystems"), zap.Any("requestID", requestID), zap.Any("result", resultSuc))
	return resultSuc, nil
}

// ListCompleteSystem - list system with completed actions started by SUSE Manager
//
// param: requestID
// param: auth
// param: actionID
// return:
func (p *Proxy) ListCompleteSystem(requestID string, auth AuthParams, actionID int) ([]interface{}, error) {
	body, err := json.Marshal(map[string]interface{}{"actionId": actionID})
	if err != nil {
		p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedMarshalling, err), zap.Any("requestID", requestID))
		return nil, fmt.Errorf(returnCodes.ErrFailedMarshalling)
	}
	path := "schedule/listCompletedSystems"
	response, err := p.suse.SuseManagerCall(body, http.MethodGet, auth.Host, path, auth.SessionKey)
	if err != nil {
		p.logger.Error("Error message recieved from suse-manger", zap.Any("requestID", requestID), zap.Any("error", err.Error()))
		return nil, fmt.Errorf(returnCodes.ErrProcessingData)
	}
	var resultSuc []interface{}
	if response.StatusCode == 200 {
		resp, err := HandleSuseManagerResponse(response.Body)
		if err != nil {
			p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrHandlingSuseManagerResponse, err), zap.Any("requestID", requestID))
			return nil, fmt.Errorf(returnCodes.ErrHandlingSuseManagerResponse)
		}
		byteArray, err := json.Marshal(resp)
		if err != nil {
			p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedMarshalling, err), zap.Any("requestID", requestID))
			return nil, fmt.Errorf(returnCodes.ErrFailedMarshalling)
		}
		err = json.Unmarshal(byteArray, &resultSuc)
		if err != nil {
			p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedUnMarshalling, err), zap.Any("requestID", requestID))
			return nil, fmt.Errorf(returnCodes.ErrFailedUnMarshalling)
		}
	}
	p.logger.Debug("response from api", zap.Any("requestID", requestID), zap.Any("api", "schedule/listCompletedSystems"), zap.Any("result", resultSuc))
	return resultSuc, nil
}

// CheckProgress - check the progress of the given action on the given system
//
// param: requestID
// param: auth
// param: actionID
// param: timeout
// param: action
// param: systemID
// return: int, error
func (p *Proxy) CheckProgress(requestID string, auth AuthParams, actionID int, timeout int, action string, systemID int) (int, error) {
	endTime := time.Now().Add(time.Second * time.Duration(timeout))
	waitTime := 15
	inProgress, err := p.ListInprogressSystem(requestID, auth, actionID)
	if err != nil {
		return 0, nil
	}
	for len(inProgress) > 0 {
		if time.Now().After(endTime) {
			p.logger.Error("action ran in timeout", zap.Any("requestID", requestID), zap.Any("action", action), zap.Any("systemID", systemID))
			return 0, fmt.Errorf("action: %s ran into timeout", action)
		}
		time.Sleep(time.Second * time.Duration(waitTime))
		inProgress, err = p.ListInprogressSystem(requestID, auth, actionID)
		if err != nil {
			return 0, nil
		}
	}
	completedSystems, err := p.ListCompleteSystem(requestID, auth, actionID)
	if err != nil {
		return 0, err
	}
	if len(completedSystems) > 0 {
		return 1, nil
	}
	return 0, fmt.Errorf("action %s is not completed", action)
}

// SystemListInstalledPackages - list installed packages on the given system
//
// param: requestID
// param: auth
// param: systemID
// return: []sumamodels.InstalledPackage, error
func (p *Proxy) SystemListInstalledPackages(requestID string, auth AuthParams, systemID int) ([]sumamodels.InstalledPackage, error) {
	body, err := json.Marshal(map[string]interface{}{"sid": systemID})
	if err != nil {
		p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedMarshalling, err), zap.Any("requestID", requestID))
		return nil, fmt.Errorf(returnCodes.ErrFailedMarshalling)
	}
	path := "system/listInstalledPackages"
	response, err := p.suse.SuseManagerCall(body, http.MethodGet, auth.Host, path, auth.SessionKey)
	if err != nil {
		return nil, fmt.Errorf(returnCodes.ErrProcessingData)
	}
	var pkgs []sumamodels.InstalledPackage
	if response.StatusCode == 200 {
		resp, err := HandleSuseManagerResponse(response.Body)
		if err != nil {
			p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrHandlingSuseManagerResponse, err), zap.Any("requestID", requestID))
			return nil, fmt.Errorf(returnCodes.ErrHandlingSuseManagerResponse)
		}
		byteArray, err := json.Marshal(resp)
		if err != nil {
			p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedMarshalling, err), zap.Any("requestID", requestID))
			return nil, fmt.Errorf(returnCodes.ErrFailedMarshalling)
		}
		err = json.Unmarshal(byteArray, &pkgs)
		if err != nil {
			p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedUnMarshalling, err), zap.Any("requestID", requestID))
			return nil, fmt.Errorf(returnCodes.ErrFailedUnMarshalling)
		}
	} else {
		p.logger.Error(fmt.Sprintf("fetching installed packages Failed. Http StatusCode: %v Http Response body: %v", response.StatusCode, string(response.Body)), zap.Any("requestID", requestID))
		return nil, fmt.Errorf(returnCodes.ErrProcessingData)
	}
	return pkgs, nil
}

// ListLatestInstallablePackages - list packages that are available for the given system
//
// param: requestID
// param: auth
// param: systemID
// return: []sumamodels.InstallablePackage, error
func (p *Proxy) ListLatestInstallablePackages(requestID string, auth AuthParams, systemID int) ([]sumamodels.InstallablePackage, error) {
	p.logger.Debug("Call list of installable packages of system api of suse manager", zap.Any("requestID", requestID))
	body, err := json.Marshal(map[string]interface{}{"sid": systemID})
	if err != nil {
		p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedMarshalling, err), zap.Any("requestID", requestID))
		return nil, fmt.Errorf(returnCodes.ErrFailedMarshalling)
	}
	path := "system/listLatestInstallablePackages"
	response, err := p.suse.SuseManagerCall(body, http.MethodGet, auth.Host, path, auth.SessionKey)
	if err != nil {
		return nil, fmt.Errorf("error while getting installable packages error: %s", err.Error())
	}
	var pacakges []sumamodels.InstallablePackage
	if response.StatusCode == 200 {
		resp, err := HandleSuseManagerResponse(response.Body)
		if err != nil {
			p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrHandlingSuseManagerResponse, err), zap.Any("requestID", requestID))
			return nil, fmt.Errorf(returnCodes.ErrHandlingSuseManagerResponse)
		}
		byteArray, err := json.Marshal(resp)
		if err != nil {
			p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedMarshalling, err), zap.Any("requestID", requestID))
			return nil, fmt.Errorf(returnCodes.ErrFailedMarshalling)
		}
		err = json.Unmarshal(byteArray, &pacakges)
		if err != nil {
			p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedUnMarshalling, err), zap.Any("requestID", requestID))
			return nil, fmt.Errorf(returnCodes.ErrFailedUnMarshalling)
		}
	} else {
		p.logger.Error(fmt.Sprintf("fetching installable packages Failed. Http StatusCode: %s Http Response body: %s", fmt.Sprint(response.StatusCode), fmt.Sprint(string(response.Body))), zap.Any("requestID", requestID))
		return nil, fmt.Errorf(returnCodes.ErrProcessingData)
	}
	return pacakges, nil
}

// SystemListActiveSystems - list system registered to SUSE Manager that are active
//
// param: requestID
// param: auth
// return: []sumamodels.ActiveSystem, error
func (p *Proxy) SystemListActiveSystems(requestID string, auth AuthParams) ([]sumamodels.ActiveSystem, error) {
	path := "system/listActiveSystems"
	response, err := p.suse.SuseManagerCall(nil, http.MethodGet, auth.Host, path, auth.SessionKey)
	if err != nil {
		return nil, fmt.Errorf("error while getting list of active systems. Error: %s", err.Error())
	}
	var systems []sumamodels.ActiveSystem
	if response.StatusCode == 200 {
		resp, err := HandleSuseManagerResponse(response.Body)
		if err != nil {
			p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrHandlingSuseManagerResponse, err), zap.Any("requestID", requestID))
			return nil, fmt.Errorf(returnCodes.ErrHandlingSuseManagerResponse)
		}
		byteArray, err := json.Marshal(resp)
		if err != nil {
			p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedMarshalling, err), zap.Any("requestID", requestID))
			return nil, fmt.Errorf(returnCodes.ErrFailedMarshalling)
		}
		err = json.Unmarshal(byteArray, &systems)
		if err != nil {
			p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedUnMarshalling, err), zap.Any("requestID", requestID))
			return nil, fmt.Errorf(returnCodes.ErrFailedUnMarshalling)
		}
	} else {
		p.logger.Error(fmt.Sprintf("calling active systems api Failed. Http StatusCode: %s Http Response body: %s", fmt.Sprint(response.StatusCode), fmt.Sprint(string(response.Body))), zap.Any("requestID", requestID))
		return nil, fmt.Errorf(returnCodes.ErrProcessingData)
	}
	return systems, nil
}

// SystemScheduleApplyHighstate - run a SALT highstate on the given system
//
// param: requestID
// param: auth
// param: systemID
// param: timeout
func (p *Proxy) SystemScheduleApplyHighstate(requestID string, auth AuthParams, systemID int, timeout int) error {
	body, err := json.Marshal(map[string]interface{}{"sid": systemID, "earliestOccurrence": time.Now(), "test": false})
	if err != nil {
		p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedMarshalling, err), zap.Any("requestID", requestID))
		return fmt.Errorf(returnCodes.ErrFailedMarshalling)
	}
	path := "system/scheduleApplyHighstate"
	response, err := p.suse.SuseManagerCall(body, http.MethodPost, auth.Host, path, auth.SessionKey)
	if err != nil {
		p.logger.Error("Error message recieved from suse-manger", zap.Any("requestID", requestID), zap.Any("error", err.Error()))
		return fmt.Errorf(returnCodes.ErrProcessingData)
	}
	return p.CheckResponseProgress(requestID, auth, response, timeout, systemID, "SystemScheduleApplyHighstate")
}

// SystemScheduleApplyStates - run a give state/states on the given system
//
// param: requestID
// param: auth
// param: systemID
// param: stateNames
// param: timeout
func (p *Proxy) SystemScheduleApplyStates(requestID string, auth AuthParams, systemID int, stateNames []string, timeout int) error {
	body, err := json.Marshal(map[string]interface{}{"sid": systemID, "stateNames": stateNames, "earliestOccurrence": time.Now(), "test": false})
	if err != nil {
		p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedMarshalling, err), zap.Any("requestID", requestID))
		return fmt.Errorf(returnCodes.ErrFailedMarshalling)
	}
	path := "system/scheduleApplyStates"
	response, err := p.suse.SuseManagerCall(body, http.MethodPost, auth.Host, path, auth.SessionKey)
	if err != nil {
		p.logger.Error("Error message recieved from suse-manger", zap.Any("requestID", requestID), zap.Any("error", err.Error()))
		return fmt.Errorf(returnCodes.ErrProcessingData)
	}
	return p.CheckResponseProgress(requestID, auth, response, timeout, systemID, "SystemScheduleApplyStates")
}

// SystemScheduleChangeChannels - change the software channels on the given system
//
// param: requestID
// param: auth
// param: systemID
// param: basechannel
// param: childChannels
func (p *Proxy) SystemScheduleChangeChannels(requestID string, auth AuthParams, systemID int, basechannel string, childChannels []sumamodels.ChannelSoftwareListChildren) error {
	p.logger.Debug("Schedule change channel api called", zap.Any("requestID", requestID))
	var childLabels []string
	for i := range childChannels {
		childLabels = append(childLabels, childChannels[i].Label)
	}
	body, err := json.Marshal(map[string]interface{}{
		"sid":                systemID,
		"baseChannelLabel":   basechannel,
		"childLabels":        childLabels,
		"earliestOccurrence": time.Now()})
	if err != nil {
		p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedMarshalling, err), zap.Any("requestID", requestID))
		return fmt.Errorf(returnCodes.ErrFailedMarshalling)
	}
	path := "system/scheduleChangeChannels"
	response, err := p.suse.SuseManagerCall(body, http.MethodPost, auth.Host, path, auth.SessionKey)
	if err != nil {
		p.logger.Error("Error message recieved from suse-manger", zap.Any("requestID", requestID), zap.Any("error", err.Error()))
		return fmt.Errorf(returnCodes.ErrProcessingData)
	}
	return p.CheckResponseProgress(requestID, auth, response, 12000, systemID, "SystemScheduleChangeChannels")
}

// SystemGetSubscribedBaseChannel - list the base channel for the give system
//
// param: requestID
// param: auth
// param: systemID
// return: sumamodels.SubscribedBaseChannel, error
func (p *Proxy) SystemGetSubscribedBaseChannel(requestID string, auth AuthParams, systemID int) (sumamodels.SubscribedBaseChannel, error) {
	p.logger.Debug("started getSubscribedBaseChannel", zap.Any("requestID", requestID))
	var result sumamodels.SubscribedBaseChannel
	body, err := json.Marshal(map[string]interface{}{"sid": systemID})
	if err != nil {
		p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedMarshalling, err), zap.Any("requestID", requestID))
		return result, fmt.Errorf(returnCodes.ErrFailedMarshalling)
	}
	path := "system/getSubscribedBaseChannel"
	response, err := p.suse.SuseManagerCall(body, http.MethodGet, auth.Host, path, auth.SessionKey)
	if err != nil {
		return result, fmt.Errorf(returnCodes.ErrProcessingData)
	}
	if response.StatusCode == 200 {
		resp, err := HandleSuseManagerResponse(response.Body)
		if err != nil {
			p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrHandlingSuseManagerResponse, err), zap.Any("requestID", requestID))
			return result, fmt.Errorf(returnCodes.ErrHandlingSuseManagerResponse)
		}
		byteArray, err := json.Marshal(resp)
		if err != nil {
			p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedMarshalling, err), zap.Any("requestID", requestID))
			return result, fmt.Errorf(returnCodes.ErrFailedMarshalling)
		}
		err = json.Unmarshal(byteArray, &result)
		if err != nil {
			p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedUnMarshalling, err), zap.Any("requestID", requestID))
			return result, fmt.Errorf(returnCodes.ErrFailedUnMarshalling)
		}
	} else {
		p.logger.Error(fmt.Sprintf("fetching basechannel Failed. Http StatusCode: %v Http Response body: %v", response.StatusCode, string(response.Body)), zap.Any("requestID", requestID))
		return result, fmt.Errorf(returnCodes.ErrProcessingData)
	}
	return result, nil
}
