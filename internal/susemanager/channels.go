// Package susemanager - SUSE Manager api call and support functions
package susemanager

import (
	"encoding/json"
	"fmt"

	sumamodels "SUSE-Manager-Tools-V2/internal/models/susemanager"
	returnCodes "SUSE-Manager-Tools-V2/internal/util/returnCodes"
	"github.com/pkg/errors"

	"go.uber.org/zap"
)

// ChannelListSoftwareChannels - list all software channels
//
// param: auth
// return:
func (p *Proxy) ChannelListSoftwareChannels(auth AuthParams) ([]sumamodels.ChannelListSoftwareChannels, error) {
	p.logger.Info("ChannelListSoftwareChannels function call started")
	path := "channel/listSoftwareChannels"
	response, err := p.suse.SuseManagerCall(nil, "GET", auth.Host, path, auth.SessionKey)
	if err != nil {
		p.logger.Error("Error message recieved from suse-manger", zap.Any("error", err.Error()))
		return nil, fmt.Errorf("error while calling list software channels manager err: %s", err.Error())
	}
	var resultSuc []sumamodels.ChannelListSoftwareChannels
	if response.StatusCode == 200 {
		resp, err := HandleSuseManagerResponse(response.Body)
		if err != nil {
			return nil, fmt.Errorf("error while handling suse manager response err: %s", err.Error())
		}
		byteArray, _ := json.Marshal(resp)
		err = json.Unmarshal(byteArray, &resultSuc)
		if err != nil {
			p.logger.Error("unmarshling error", zap.Any("error", err.Error()))
			return nil, fmt.Errorf("error while calling list software channels manager err: %s", err.Error())
		}
	} else {
		p.logger.Error("list software channels call failed", zap.Any("status code", response.StatusCode))
		return nil, fmt.Errorf("calling list software channels manager Failed. Http StatusCode: %s Http Body: %s", fmt.Sprint(response.StatusCode), fmt.Sprint(response.Body))
	}
	p.logger.Info("Completed ChannelListSoftwareChannels function")
	return resultSuc, nil
}

// ChannelSoftwareListChildren - list software channels from given parent
//
// param: auth
// param: label
// return:
func (p *Proxy) ChannelSoftwareListChildren(auth AuthParams, label string) ([]sumamodels.ChannelSoftwareListChildren, error) {
	p.logger.Info("Inside ChannelSoftwareListChildren function")
	body, _ := json.Marshal(map[string]interface{}{"channelLabel": label})
	path := "channel/software/listChildren"
	response, err := p.suse.SuseManagerCall(body, "GET", auth.Host, path, auth.SessionKey)
	if err != nil {
		p.logger.Error("Error message recieved from suse-manger", zap.Any("error", err.Error()))
		return nil, fmt.Errorf("error while calling list software channels err: %s", err.Error())
	}
	var resultSuc []sumamodels.ChannelSoftwareListChildren
	if response.StatusCode == 200 {
		resp, err := HandleSuseManagerResponse(response.Body)
		if err != nil {
			return nil, fmt.Errorf("error while handling suse manager response err: %s", err.Error())
		}
		byteArray, _ := json.Marshal(resp)
		err = json.Unmarshal(byteArray, &resultSuc)
		if err != nil {
			p.logger.Error("unmarshling error", zap.Any("error", err.Error()))
			return nil, fmt.Errorf("error while calling list child software channels , err: %s", err.Error())
		}
	} else {
		return nil, fmt.Errorf("calling list software channels Failed. Http StatusCode: %s Http Body: %s", fmt.Sprint(response.StatusCode), fmt.Sprint(response.Body))
	}
	p.logger.Info("Completed ChannelSoftwareListChildren function")
	return resultSuc, nil
}

// ChannelSoftwareCreateRepo
//
// param: auth
// param: label
// param: typeRepo
// param: url
// return:
func (p *Proxy) ChannelSoftwareCreateRepo(auth AuthParams, label string, typeRepo string, url string) (sumamodels.ChannelSoftwareCreateRepo, error) {
	p.logger.Info("Inside ChannelSoftwareCreateRepo function")
	body, err := json.Marshal(map[string]interface{}{"label": label, "type": typeRepo, "url": url})
	var resultSuc sumamodels.ChannelSoftwareCreateRepo
	if err != nil {
		p.logger.Error(returnCodes.ErrFailedMarshalling, zap.Any("error", err))
		return resultSuc, errors.New(returnCodes.ErrFailedMarshalling)
	}
	path := "channel/software/createRepo"
	response, err := p.suse.SuseManagerCall(body, "POST", auth.Host, path, auth.SessionKey)
	if err != nil {
		p.logger.Error(returnCodes.ErrHandlingSuseManagerResponse, zap.Any("error", err.Error()))
		return resultSuc, fmt.Errorf(returnCodes.ErrHandlingSuseManagerResponse)
	}
	if response.StatusCode == 200 {
		resp, err := HandleSuseManagerResponse(response.Body)
		if err != nil {
			p.logger.Error(returnCodes.ErrHandlingSuseManagerResponse, zap.Any("response", resp), zap.Any("error", err.Error()))
			return resultSuc, errors.New(returnCodes.ErrHandlingSuseManagerResponse)
		}
		byteArray, _ := json.Marshal(resp)
		err = json.Unmarshal(byteArray, &resultSuc)
		if err != nil {
			p.logger.Error(returnCodes.ErrFailedUnMarshalling, zap.Any("error", err.Error()))
			return resultSuc, errors.New(returnCodes.ErrFailedUnMarshalling)
		}
	} else {
		p.logger.Error(returnCodes.ErrHTTPSuseManagerResponse, zap.Any("HTTP Statuscode", response.StatusCode), zap.Any("HTTP body", response.Body))
		return resultSuc, errors.New(returnCodes.ErrHandlingSuseManagerResponse)
	}
	p.logger.Info("Completed ChannelSoftwareCreateRepo function")
	return resultSuc, nil
}

// ChannelSoftwareCreate
//
// param: auth
// param: label
// param: name
// param: summary
// param: archLabel
// param: parentLabel
// return:
func (p *Proxy) ChannelSoftwareCreate(auth AuthParams, label string, name string, summary string, archLabel string, parentLabel string) (int, error) {
	p.logger.Info("Inside ChannelSoftwareCreate function")
	var resultSuc int
	body, err := json.Marshal(map[string]interface{}{"label": label, "summary": summary, "archLabel": archLabel, "parentLabel": parentLabel, "name": name})
	if err != nil {
		p.logger.Error(returnCodes.ErrFailedMarshalling, zap.Any("error", err))
		return 0, errors.New(returnCodes.ErrFailedMarshalling)
	}
	path := "channel/software/create"
	response, err := p.suse.SuseManagerCall(body, "POST", auth.Host, path, auth.SessionKey)
	if err != nil {
		p.logger.Error(returnCodes.ErrHandlingSuseManagerResponse, zap.Any("error", err.Error()))
		return 0, fmt.Errorf(returnCodes.ErrHandlingSuseManagerResponse)
	}
	if response.StatusCode == 200 {
		resp, err := HandleSuseManagerResponse(response.Body)
		if err != nil {
			p.logger.Error(returnCodes.ErrHandlingSuseManagerResponse, zap.Any("response", resp), zap.Any("error", err.Error()))
			return 0, errors.New(returnCodes.ErrHandlingSuseManagerResponse)
		}
		byteArray, _ := json.Marshal(resp)
		err = json.Unmarshal(byteArray, &resultSuc)
		if err != nil {
			p.logger.Error(returnCodes.ErrFailedUnMarshalling, zap.Any("error", err.Error()))
			return 0, errors.New(returnCodes.ErrFailedUnMarshalling)
		}
	} else {
		p.logger.Error(returnCodes.ErrHTTPSuseManagerResponse, zap.Any("HTTP Statuscode", response.StatusCode), zap.Any("HTTP body", response.Body))
		return 0, errors.New(returnCodes.ErrHandlingSuseManagerResponse)
	}
	p.logger.Info("Completed ChannelSoftwareCreate function")
	return resultSuc, nil
}

// ChannelSoftwareAssociateRepo
//
// param: auth
// param: channelLabel
// param: repoLabel
// return:
func (p *Proxy) ChannelSoftwareAssociateRepo(auth AuthParams, channelLabel string, repoLabel string) (sumamodels.ChannelSoftwareListChildren, error) {
	p.logger.Info("Inside ChannelSoftwareAssociateRepo function")
	var resultSuc sumamodels.ChannelSoftwareListChildren
	body, err := json.Marshal(map[string]interface{}{"channelLabel": channelLabel, "repoLabel": repoLabel})
	if err != nil {
		p.logger.Error(returnCodes.ErrFailedMarshalling, zap.Any("error", err))
		return resultSuc, errors.New(returnCodes.ErrFailedMarshalling)
	}
	path := "channel/software/associateRepo"
	response, err := p.suse.SuseManagerCall(body, "POST", auth.Host, path, auth.SessionKey)
	if err != nil {
		p.logger.Error(returnCodes.ErrHandlingSuseManagerResponse, zap.Any("error", err.Error()))
		return resultSuc, fmt.Errorf(returnCodes.ErrHandlingSuseManagerResponse)
	}
	if response.StatusCode == 200 {
		resp, err := HandleSuseManagerResponse(response.Body)
		if err != nil {
			p.logger.Error(returnCodes.ErrHandlingSuseManagerResponse, zap.Any("response", resp), zap.Any("error", err.Error()))
			return resultSuc, errors.New(returnCodes.ErrHandlingSuseManagerResponse)
		}
		byteArray, _ := json.Marshal(resp)
		err = json.Unmarshal(byteArray, &resultSuc)
		if err != nil {
			p.logger.Error(returnCodes.ErrFailedUnMarshalling, zap.Any("error", err.Error()))
			return resultSuc, errors.New(returnCodes.ErrFailedUnMarshalling)
		}
	} else {
		p.logger.Error(returnCodes.ErrHTTPSuseManagerResponse, zap.Any("HTTP Statuscode", response.StatusCode), zap.Any("HTTP body", response.Body))
		return resultSuc, errors.New(returnCodes.ErrHandlingSuseManagerResponse)
	}
	p.logger.Info("Completed ChannelSoftwareAssociateRepo function")
	return resultSuc, nil
}

// ChannelSoftwareSyncRepo
//
// param: auth
// param: channelLabel
// return:
func (p *Proxy) ChannelSoftwareSyncRepo(auth AuthParams, channelLabel string) (int, error) {
	p.logger.Info("Inside ChannelSoftwareSyncRepo function")
	var resultSuc int
	body, err := json.Marshal(map[string]interface{}{"channelLabel": channelLabel})
	if err != nil {
		p.logger.Error(returnCodes.ErrFailedMarshalling, zap.Any("error", err))
		return resultSuc, errors.New(returnCodes.ErrFailedMarshalling)
	}
	path := "channel/software/syncRepo"
	response, err := p.suse.SuseManagerCall(body, "POST", auth.Host, path, auth.SessionKey)
	if err != nil {
		p.logger.Error(returnCodes.ErrHandlingSuseManagerResponse, zap.Any("error", err.Error()))
		return resultSuc, fmt.Errorf(returnCodes.ErrHandlingSuseManagerResponse)
	}
	if response.StatusCode == 200 {
		resp, err := HandleSuseManagerResponse(response.Body)
		if err != nil {
			p.logger.Error(returnCodes.ErrHandlingSuseManagerResponse, zap.Any("response", resp), zap.Any("error", err.Error()))
			return resultSuc, errors.New(returnCodes.ErrHandlingSuseManagerResponse)
		}
		byteArray, _ := json.Marshal(resp)
		err = json.Unmarshal(byteArray, &resultSuc)
		if err != nil {
			p.logger.Error(returnCodes.ErrFailedUnMarshalling, zap.Any("error", err.Error()))
			return resultSuc, errors.New(returnCodes.ErrFailedUnMarshalling)
		}
	} else {
		p.logger.Error(returnCodes.ErrHTTPSuseManagerResponse, zap.Any("HTTP Statuscode", response.StatusCode), zap.Any("HTTP body", response.Body))
		return resultSuc, errors.New(returnCodes.ErrHandlingSuseManagerResponse)
	}
	p.logger.Info("Completed ChannelSoftwareSyncRepo function")
	return resultSuc, nil
}

// ChannelSoftwareIsExisting
//
// param: auth
// param: label
// return:
func (p *Proxy) ChannelSoftwareIsExisting(auth AuthParams, label string) (bool, error) {
	p.logger.Info("Inside ChannelSoftwareIsExisting function", zap.Any("Label", label))
	var resultSuc bool
	body, err := json.Marshal(map[string]interface{}{"channelLabel": label})
	if err != nil {
		p.logger.Error(returnCodes.ErrFailedMarshalling, zap.Any("error", err))
		return false, errors.New(returnCodes.ErrFailedMarshalling)
	}
	path := "channel/software/isExisting"
	response, err := p.suse.SuseManagerCall(body, "GET", auth.Host, path, auth.SessionKey)
	if err != nil {
		p.logger.Error(returnCodes.ErrHandlingSuseManagerResponse, zap.Any("error", err.Error()))
		return false, fmt.Errorf(returnCodes.ErrHandlingSuseManagerResponse)
	}
	if response.StatusCode == 200 {
		resp, err := HandleSuseManagerResponse(response.Body)
		if err != nil {
			p.logger.Error(returnCodes.ErrHandlingSuseManagerResponse, zap.Any("response", resp), zap.Any("error", err.Error()))
			return false, errors.New(returnCodes.ErrFailedUnMarshalling)
		}

		byteArray, _ := json.Marshal(resp)
		// p.logger.Info(byteArray)
		err = json.Unmarshal(byteArray, &resultSuc)
		if err != nil {
			p.logger.Error(returnCodes.ErrFailedUnMarshalling, zap.Any("error", err.Error()))
			return false, errors.New(returnCodes.ErrFailedUnMarshalling)
		}
	} else {
		p.logger.Error(returnCodes.ErrHTTPSuseManagerResponse, zap.Any("HTTP Statuscode", response.StatusCode), zap.Any("HTTP body", response.Body))
		return false, errors.New(returnCodes.ErrHandlingSuseManagerResponse)
	}
	p.logger.Info("Completed ChannelSoftwareIsExisting function")
	return resultSuc, nil
}
