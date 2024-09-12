// Package susemanager api call for SUSE Manager related to system
package susemanager

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	sumamodels "SUSE-Manager-Tools-V2/internal/models/susemanager"
	returnCodes "SUSE-Manager-Tools-V2/internal/util/returnCodes"

	"go.uber.org/zap"
)

// SystemGetID - get systemID from the given server
//
// param: auth
// param: systemName
// return: []sumamodels.System, error
func (p *Proxy) SystemGetID(auth AuthParams, systemName string) ([]sumamodels.System, error) {
	var systeminfo []sumamodels.System
	body, err := json.Marshal(map[string]string{"name": systemName})
	if err != nil {
		p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedMarshalling, err))
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
			p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrHandlingSuseManagerResponse, err))
			return nil, fmt.Errorf(returnCodes.ErrHandlingSuseManagerResponse)
		}
		byteArray, err := json.Marshal(resp)
		if err != nil {
			p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedMarshalling, err))
			return nil, fmt.Errorf(returnCodes.ErrFailedMarshalling)
		}
		err = json.Unmarshal(byteArray, &systeminfo)
		if err != nil {
			p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedUnMarshalling, err))
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
// param: auth
// param: systemID
func (p *Proxy) SchedulePackageRefresh(auth AuthParams, systemID int) error {
	p.logger.Debug("Scheduling package refresh called")
	body, err := json.Marshal(map[string]interface{}{"sid": systemID, "earliestOccurrence": time.Now()})
	if err != nil {
		p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedMarshalling, err))
		return fmt.Errorf(returnCodes.ErrFailedMarshalling)
	}
	path := "system/schedulePackageRefresh"
	response, err := p.suse.SuseManagerCall(body, http.MethodPost, auth.Host, path, auth.SessionKey)
	if err != nil {
		p.logger.Error(fmt.Sprintf("error while scheduling package refresh err: %s", response.Body))
		return errors.New(returnCodes.ErrProcessingData)
	}
	return p.CheckResponseProgress(auth, response, 12000, systemID, "SchedulePackageRefresh")
}

// ScheduleScriptRun - schedule running a script on the given server
//
// param: auth
// param: systemID
// param: timeout
// param: script
func (p *Proxy) ScheduleScriptRun(auth AuthParams, systemID int, timeout int, script string) error {
	p.logger.Info("script run scheduled", zap.Any("script", script), zap.Any("systemID", systemID))
	body, err := json.Marshal(map[string]interface{}{"sid": systemID,
		"username":           "root",
		"groupname":          "root",
		"timeout":            timeout,
		"script":             script,
		"earliestOccurrence": time.Now()})
	if err != nil {
		p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedMarshalling, err))
		return fmt.Errorf(returnCodes.ErrFailedMarshalling)
	}
	path := "system/scheduleScriptRun"
	response, err := p.suse.SuseManagerCall(body, http.MethodPost, auth.Host, path, auth.SessionKey)
	if err != nil {
		p.logger.Error("Error message recieved from suse-manger", zap.Any("suma-host", auth.Host), zap.Any("error", err.Error()))
		return fmt.Errorf(returnCodes.ErrProcessingData)
	}
	return p.CheckResponseProgress(auth, response, timeout, systemID, "ScheduleScriptRun")
}

// SystemGetScriptResult - get results from script run
//
// param: auth
// param: actionID
// param: resultCompleted
// return: string, error
func (p *Proxy) SystemGetScriptResult(auth AuthParams, actionID int, resultCompleted int) (string, error) {
	body, err := json.Marshal(map[string]interface{}{"actionId": actionID})
	if err != nil {
		p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedMarshalling, err))
		return "", fmt.Errorf(returnCodes.ErrFailedMarshalling)
	}
	path := "system/getScriptResults"
	response, err := p.suse.SuseManagerCall(body, http.MethodPost, auth.Host, path, auth.SessionKey)
	if err != nil {
		p.logger.Error("Error message recieved from suse-manger", zap.Any("error", err.Error()))
		return "", fmt.Errorf(returnCodes.ErrProcessingData)
	}
	var scriptResults []sumamodels.ScriptResult
	var output string
	if response.StatusCode == 200 {
		resp, err := HandleSuseManagerResponse(response.Body)
		if err != nil {
			p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrHandlingSuseManagerResponse, err))
			return "", fmt.Errorf(returnCodes.ErrHandlingSuseManagerResponse)
		}
		byteArray, err := json.Marshal(resp)
		if err != nil {
			p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedMarshalling, err))
			return "", fmt.Errorf(returnCodes.ErrFailedMarshalling)
		}
		err = json.Unmarshal(byteArray, &scriptResults)
		if err != nil {
			p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedUnMarshalling, err))
			return "", fmt.Errorf(returnCodes.ErrFailedUnMarshalling)
		}
		for i := range scriptResults {
			output = scriptResults[i].Output
		}
		if resultCompleted != 1 {
			p.logger.Error("Error from script run", zap.Any("Script message", output))
			return "", fmt.Errorf(returnCodes.ErrProcessingData)
		}
	}
	return output, nil
}

// SystemScheduleReboot - reboot the given system
//
// param: auth
// param: systemID
// param: timeout
func (p *Proxy) SystemScheduleReboot(auth AuthParams, systemID int, timeout int) error {
	body, err := json.Marshal(map[string]interface{}{"sid": systemID, "earliestOccurrence": time.Now()})
	if err != nil {
		p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedMarshalling, err))
		return fmt.Errorf(returnCodes.ErrFailedMarshalling)
	}
	path := "system/scheduleReboot"
	response, err := p.suse.SuseManagerCall(body, http.MethodPost, auth.Host, path, auth.SessionKey)
	if err != nil {
		p.logger.Error("Error message recieved from suse-manger", zap.Any("error", err.Error()))
		return fmt.Errorf(returnCodes.ErrProcessingData)
	}
	return p.CheckResponseProgress(auth, response, timeout, systemID, "SystemScheduleReboot")
}

// ListInprogressSystem - list systems with running actions started from SUSE Manager
//
// param: auth
// param: actionID
// return:
func (p *Proxy) ListInprogressSystem(auth AuthParams, actionID int) ([]interface{}, error) {
	body, err := json.Marshal(map[string]interface{}{"actionId": actionID})
	if err != nil {
		p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedMarshalling, err))
		return nil, fmt.Errorf(returnCodes.ErrFailedMarshalling)
	}
	path := "schedule/listInProgressSystems"
	response, err := p.suse.SuseManagerCall(body, http.MethodGet, auth.Host, path, auth.SessionKey)
	if err != nil {
		p.logger.Error("Error message recieved from suse-manger", zap.Any("error", err.Error()))
		return nil, fmt.Errorf(returnCodes.ErrProcessingData)
	}
	var resultSuc []interface{}
	if response.StatusCode == 200 {
		resp, err := HandleSuseManagerResponse(response.Body)
		if err != nil {
			p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrHandlingSuseManagerResponse, err))
			return nil, fmt.Errorf(returnCodes.ErrHandlingSuseManagerResponse)
		}
		byteArray, err := json.Marshal(resp)
		if err != nil {
			p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedMarshalling, err))
			return nil, fmt.Errorf(returnCodes.ErrFailedMarshalling)
		}
		err = json.Unmarshal(byteArray, &resultSuc)
		if err != nil {
			p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedUnMarshalling, err))
			return nil, fmt.Errorf(returnCodes.ErrFailedUnMarshalling)
		}
	}
	p.logger.Debug("Response from api", zap.Any("api", "schedule/listInProgressSystems"), zap.Any("result", resultSuc))
	return resultSuc, nil
}

// ListCompleteSystem - list system with completed actions started by SUSE Manager
//
// param: auth
// param: actionID
// return:
func (p *Proxy) ListCompleteSystem(auth AuthParams, actionID int) ([]interface{}, error) {
	body, err := json.Marshal(map[string]interface{}{"actionId": actionID})
	if err != nil {
		p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedMarshalling, err))
		return nil, fmt.Errorf(returnCodes.ErrFailedMarshalling)
	}
	path := "schedule/listCompletedSystems"
	response, err := p.suse.SuseManagerCall(body, http.MethodGet, auth.Host, path, auth.SessionKey)
	if err != nil {
		p.logger.Error("Error message recieved from suse-manger", zap.Any("error", err.Error()))
		return nil, fmt.Errorf(returnCodes.ErrProcessingData)
	}
	var resultSuc []interface{}
	if response.StatusCode == 200 {
		resp, err := HandleSuseManagerResponse(response.Body)
		if err != nil {
			p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrHandlingSuseManagerResponse, err))
			return nil, fmt.Errorf(returnCodes.ErrHandlingSuseManagerResponse)
		}
		byteArray, err := json.Marshal(resp)
		if err != nil {
			p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedMarshalling, err))
			return nil, fmt.Errorf(returnCodes.ErrFailedMarshalling)
		}
		err = json.Unmarshal(byteArray, &resultSuc)
		if err != nil {
			p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedUnMarshalling, err))
			return nil, fmt.Errorf(returnCodes.ErrFailedUnMarshalling)
		}
	}
	p.logger.Debug("response from api", zap.Any("api", "schedule/listCompletedSystems"), zap.Any("result", resultSuc))
	return resultSuc, nil
}

// CheckProgress - check the progress of the given action on the given system
//
// param: auth
// param: actionID
// param: timeout
// param: action
// param: systemID
// return: int, error
func (p *Proxy) CheckProgress(auth AuthParams, actionID int, timeout int, action string, systemID int) (int, error) {
	endTime := time.Now().Add(time.Second * time.Duration(timeout))
	waitTime := 15
	inProgress, err := p.ListInprogressSystem(auth, actionID)
	if err != nil {
		return 0, nil
	}
	for len(inProgress) > 0 {
		if time.Now().After(endTime) {
			p.logger.Error("action ran in timeout", zap.Any("action", action), zap.Any("systemID", systemID))
			return 0, fmt.Errorf("action: %s ran into timeout", action)
		}
		time.Sleep(time.Second * time.Duration(waitTime))
		inProgress, err = p.ListInprogressSystem(auth, actionID)
		if err != nil {
			return 0, nil
		}
	}
	completedSystems, err := p.ListCompleteSystem(auth, actionID)
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
// param: auth
// param: systemID
// return: []sumamodels.InstalledPackage, error
func (p *Proxy) SystemListInstalledPackages(auth AuthParams, systemID int) ([]sumamodels.InstalledPackage, error) {
	body, err := json.Marshal(map[string]interface{}{"sid": systemID})
	if err != nil {
		p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedMarshalling, err))
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
			p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrHandlingSuseManagerResponse, err))
			return nil, fmt.Errorf(returnCodes.ErrHandlingSuseManagerResponse)
		}
		byteArray, err := json.Marshal(resp)
		if err != nil {
			p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedMarshalling, err))
			return nil, fmt.Errorf(returnCodes.ErrFailedMarshalling)
		}
		err = json.Unmarshal(byteArray, &pkgs)
		if err != nil {
			p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedUnMarshalling, err))
			return nil, fmt.Errorf(returnCodes.ErrFailedUnMarshalling)
		}
	} else {
		p.logger.Error(fmt.Sprintf("fetching installed packages Failed. Http StatusCode: %v Http Response body: %v", response.StatusCode, string(response.Body)))
		return nil, fmt.Errorf(returnCodes.ErrProcessingData)
	}
	return pkgs, nil
}

// ListLatestInstallablePackages - list packages that are available for the given system
//
// param: auth
// param: systemID
// return: []sumamodels.InstallablePackage, error
func (p *Proxy) ListLatestInstallablePackages(auth AuthParams, systemID int) ([]sumamodels.InstallablePackage, error) {
	p.logger.Debug("Call list of installable packages of system api of suse manager")
	body, err := json.Marshal(map[string]interface{}{"sid": systemID})
	if err != nil {
		p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedMarshalling, err))
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
			p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrHandlingSuseManagerResponse, err))
			return nil, fmt.Errorf(returnCodes.ErrHandlingSuseManagerResponse)
		}
		byteArray, err := json.Marshal(resp)
		if err != nil {
			p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedMarshalling, err))
			return nil, fmt.Errorf(returnCodes.ErrFailedMarshalling)
		}
		err = json.Unmarshal(byteArray, &pacakges)
		if err != nil {
			p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedUnMarshalling, err))
			return nil, fmt.Errorf(returnCodes.ErrFailedUnMarshalling)
		}
	} else {
		p.logger.Error(fmt.Sprintf("fetching installable packages Failed. Http StatusCode: %s Http Response body: %s", fmt.Sprint(response.StatusCode), fmt.Sprint(string(response.Body))))
		return nil, fmt.Errorf(returnCodes.ErrProcessingData)
	}
	return pacakges, nil
}

// SystemListActiveSystems - list system registered to SUSE Manager that are active
//
// param: auth
// return: []sumamodels.ActiveSystem, error
func (p *Proxy) SystemListActiveSystems(auth AuthParams) ([]sumamodels.ActiveSystem, error) {
	path := "system/listActiveSystems"
	response, err := p.suse.SuseManagerCall(nil, http.MethodGet, auth.Host, path, auth.SessionKey)
	if err != nil {
		return nil, fmt.Errorf("error while getting list of active systems. Error: %s", err.Error())
	}
	var systems []sumamodels.ActiveSystem
	if response.StatusCode == 200 {
		resp, err := HandleSuseManagerResponse(response.Body)
		if err != nil {
			p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrHandlingSuseManagerResponse, err))
			return nil, fmt.Errorf(returnCodes.ErrHandlingSuseManagerResponse)
		}
		byteArray, err := json.Marshal(resp)
		if err != nil {
			p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedMarshalling, err))
			return nil, fmt.Errorf(returnCodes.ErrFailedMarshalling)
		}
		err = json.Unmarshal(byteArray, &systems)
		if err != nil {
			p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedUnMarshalling, err))
			return nil, fmt.Errorf(returnCodes.ErrFailedUnMarshalling)
		}
	} else {
		p.logger.Error(fmt.Sprintf("calling active systems api Failed. Http StatusCode: %s Http Response body: %s", fmt.Sprint(response.StatusCode), fmt.Sprint(string(response.Body))))
		return nil, fmt.Errorf(returnCodes.ErrProcessingData)
	}
	return systems, nil
}

// SystemListInActiveSystems
//
// param: auth
// return:
func (p *Proxy) SystemListInActiveSystems(auth AuthParams) ([]sumamodels.ActiveSystem, error) {
	path := "system/listInactiveSystems"
	response, err := p.suse.SuseManagerCall(nil, http.MethodGet, auth.Host, path, auth.SessionKey)
	if err != nil {
		return nil, fmt.Errorf("error while getting list of active systems. Error: %s", err.Error())
	}
	var systems []sumamodels.ActiveSystem
	if response.StatusCode == 200 {
		resp, err := HandleSuseManagerResponse(response.Body)
		if err != nil {
			p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrHandlingSuseManagerResponse, err))
			return nil, fmt.Errorf(returnCodes.ErrHandlingSuseManagerResponse)
		}
		byteArray, err := json.Marshal(resp)
		if err != nil {
			p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedMarshalling, err))
			return nil, fmt.Errorf(returnCodes.ErrFailedMarshalling)
		}
		err = json.Unmarshal(byteArray, &systems)
		if err != nil {
			p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedUnMarshalling, err))
			return nil, fmt.Errorf(returnCodes.ErrFailedUnMarshalling)
		}
	} else {
		p.logger.Error(fmt.Sprintf("calling active systems api Failed. Http StatusCode: %s Http Response body: %s", fmt.Sprint(response.StatusCode), fmt.Sprint(string(response.Body))))
		return nil, fmt.Errorf(returnCodes.ErrProcessingData)
	}
	return systems, nil
}

// SystemScheduleApplyHighstate - run a SALT highstate on the given system
//
// param: auth
// param: systemID
// param: timeout
func (p *Proxy) SystemScheduleApplyHighstate(auth AuthParams, systemID int, timeout int) error {
	body, err := json.Marshal(map[string]interface{}{"sid": systemID, "earliestOccurrence": time.Now(), "test": false})
	if err != nil {
		p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedMarshalling, err))
		return fmt.Errorf(returnCodes.ErrFailedMarshalling)
	}
	path := "system/scheduleApplyHighstate"
	response, err := p.suse.SuseManagerCall(body, http.MethodPost, auth.Host, path, auth.SessionKey)
	if err != nil {
		p.logger.Error("Error message recieved from suse-manger", zap.Any("error", err.Error()))
		return fmt.Errorf(returnCodes.ErrProcessingData)
	}
	return p.CheckResponseProgress(auth, response, timeout, systemID, "SystemScheduleApplyHighstate")
}

// SystemScheduleApplyStates - run a give state/states on the given system
//
// param: auth
// param: systemID
// param: stateNames
// param: timeout
func (p *Proxy) SystemScheduleApplyStates(auth AuthParams, systemID int, stateNames []string, timeout int) error {
	body, err := json.Marshal(map[string]interface{}{"sid": systemID, "stateNames": stateNames, "earliestOccurrence": time.Now(), "test": false})
	if err != nil {
		p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedMarshalling, err))
		return fmt.Errorf(returnCodes.ErrFailedMarshalling)
	}
	path := "system/scheduleApplyStates"
	response, err := p.suse.SuseManagerCall(body, http.MethodPost, auth.Host, path, auth.SessionKey)
	if err != nil {
		p.logger.Error("Error message recieved from suse-manger", zap.Any("error", err.Error()))
		return fmt.Errorf(returnCodes.ErrProcessingData)
	}
	return p.CheckResponseProgress(auth, response, timeout, systemID, "SystemScheduleApplyStates")
}

// SystemScheduleChangeChannels - change the software channels on the given system
//
// param: auth
// param: systemID
// param: basechannel
// param: childChannels
func (p *Proxy) SystemScheduleChangeChannels(auth AuthParams, systemID int, basechannel string, childChannels []sumamodels.ChannelSoftwareListChildren) error {
	p.logger.Debug("Schedule change channel api called")
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
		p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedMarshalling, err))
		return fmt.Errorf(returnCodes.ErrFailedMarshalling)
	}
	path := "system/scheduleChangeChannels"
	response, err := p.suse.SuseManagerCall(body, http.MethodPost, auth.Host, path, auth.SessionKey)
	if err != nil {
		p.logger.Error("Error message recieved from suse-manger", zap.Any("error", err.Error()))
		return fmt.Errorf(returnCodes.ErrProcessingData)
	}
	return p.CheckResponseProgress(auth, response, 12000, systemID, "SystemScheduleChangeChannels")
}

// SystemGetSubscribedBaseChannel - list the base channel for the give system
//
// param: auth
// param: systemID
// return: sumamodels.SubscribedBaseChannel, error
func (p *Proxy) SystemGetSubscribedBaseChannel(auth AuthParams, systemID int) (sumamodels.SubscribedBaseChannel, error) {
	p.logger.Debug("started getSubscribedBaseChannel")
	var result sumamodels.SubscribedBaseChannel
	body, err := json.Marshal(map[string]interface{}{"sid": systemID})
	if err != nil {
		p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedMarshalling, err))
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
			p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrHandlingSuseManagerResponse, err))
			return result, fmt.Errorf(returnCodes.ErrHandlingSuseManagerResponse)
		}
		byteArray, err := json.Marshal(resp)
		if err != nil {
			p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedMarshalling, err))
			return result, fmt.Errorf(returnCodes.ErrFailedMarshalling)
		}
		err = json.Unmarshal(byteArray, &result)
		if err != nil {
			p.logger.Error(fmt.Sprintf("%v error %v", returnCodes.ErrFailedUnMarshalling, err))
			return result, fmt.Errorf(returnCodes.ErrFailedUnMarshalling)
		}
	} else {
		p.logger.Error(fmt.Sprintf("fetching basechannel Failed. Http StatusCode: %v Http Response body: %v", response.StatusCode, string(response.Body)))
		return result, fmt.Errorf(returnCodes.ErrProcessingData)
	}
	return result, nil
}
