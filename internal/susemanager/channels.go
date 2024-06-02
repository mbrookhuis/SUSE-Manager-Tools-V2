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
// param: requestID
// param: auth
// return:
func (p *Proxy) ChannelListSoftwareChannels(requestID string, auth AuthParams) ([]sumamodels.ChannelListSoftwareChannels, error) {
	p.logger.Info("ChannelListSoftwareChannels function call started", zap.Any("requestID", requestID))
	path := "channel/listSoftwareChannels"
	response, err := p.suse.SuseManagerCall(nil, "GET", auth.Host, path, auth.SessionKey)
	if err != nil {
		p.logger.Error("Error message recieved from suse-manger", zap.Any("requestID", requestID), zap.Any("error", err.Error()))
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
			p.logger.Error("unmarshling error", zap.Any("requestID", requestID), zap.Any("error", err.Error()))
			return nil, fmt.Errorf("error while calling list software channels manager err: %s", err.Error())
		}
	} else {
		p.logger.Error("list software channels call failed", zap.Any("resquestId", requestID), zap.Any("status code", response.StatusCode))
		return nil, fmt.Errorf("calling list software channels manager Failed. Http StatusCode: %s Http Body: %s", fmt.Sprint(response.StatusCode), fmt.Sprint(response.Body))
	}
	p.logger.Info("Completed ChannelListSoftwareChannels function", zap.Any("requestID", requestID))
	return resultSuc, nil
}

// ChannelSoftwareListChildren - list software channels from given parent
//
// param: requestID
// param: auth
// param: label
// return:
func (p *Proxy) ChannelSoftwareListChildren(requestID string, auth AuthParams, label string) ([]sumamodels.ChannelSoftwareListChildren, error) {
	p.logger.Info("Inside ChannelSoftwareListChildren function", zap.Any("requestID", requestID))
	body, _ := json.Marshal(map[string]interface{}{"channelLabel": label})
	path := "channel/software/listChildren"
	response, err := p.suse.SuseManagerCall(body, "GET", auth.Host, path, auth.SessionKey)
	if err != nil {
		p.logger.Error("Error message recieved from suse-manger", zap.Any("requestID", requestID), zap.Any("error", err.Error()))
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
			p.logger.Error("unmarshling error", zap.Any("requestID", requestID), zap.Any("error", err.Error()))
			return nil, fmt.Errorf("error while calling list child software channels , err: %s", err.Error())
		}
	} else {
		return nil, fmt.Errorf("calling list software channels Failed. Http StatusCode: %s Http Body: %s", fmt.Sprint(response.StatusCode), fmt.Sprint(response.Body))
	}
	p.logger.Info("Completed ChannelSoftwareListChildren function", zap.Any("requestID", requestID))
	return resultSuc, nil
}

func (p *Proxy) ChannelSoftwareCreateRepo(requestID string, auth AuthParams, label string, typeRepo string, url string) (sumamodels.ChannelSoftwareCreateRepo, error) {
	p.logger.Info("Inside ChannelSoftwareCreateRepo function", zap.Any("requestID", requestID))
	body, err := json.Marshal(map[string]interface{}{"label": label, "type": typeRepo, "url": url})
	var resultSuc sumamodels.ChannelSoftwareCreateRepo
	if err != nil {
		p.logger.Error(returnCodes.ErrFailedMarshalling, zap.Any("error", err))
		return resultSuc, errors.New(returnCodes.ErrFailedMarshalling)
	}
	path := "channel/software/createRepo"
	response, err := p.suse.SuseManagerCall(body, "POST", auth.Host, path, auth.SessionKey)
	if err != nil {
		p.logger.Error(returnCodes.ErrHandlingSuseManagerResponse, zap.Any("requestID", requestID), zap.Any("error", err.Error()))
		return resultSuc, fmt.Errorf(returnCodes.ErrHandlingSuseManagerResponse)
	}
	if response.StatusCode == 200 {
		resp, err := HandleSuseManagerResponse(response.Body)
		if err != nil {
			p.logger.Error(returnCodes.ErrHandlingSuseManagerResponse, zap.Any("requestID", requestID), zap.Any("response", resp), zap.Any("error", err.Error()))
			return resultSuc, errors.New(returnCodes.ErrHandlingSuseManagerResponse)
		}
		byteArray, _ := json.Marshal(resp)
		err = json.Unmarshal(byteArray, &resultSuc)
		if err != nil {
			p.logger.Error(returnCodes.ErrFailedUnMarshalling, zap.Any("requestID", requestID), zap.Any("error", err.Error()))
			return resultSuc, errors.New(returnCodes.ErrFailedUnMarshalling)
		}
	} else {
		p.logger.Error(returnCodes.ErrHTTPSuseManagerResponse, zap.Any("HTTP Statuscode", response.StatusCode), zap.Any("HTTP body", response.Body))
		return resultSuc, errors.New(returnCodes.ErrHandlingSuseManagerResponse)
	}
	p.logger.Info("Completed ChannelSoftwareCreateRepo function", zap.Any("requestID", requestID))
	return resultSuc, nil
}

func (p *Proxy) ChannelSoftwareCreate(requestID string, auth AuthParams, label string, name string, summary string, archLabel string, parentLabel string) (int, error) {
	p.logger.Info("Inside ChannelSoftwareCreate function", zap.Any("requestID", requestID))
	var resultSuc int
	body, err := json.Marshal(map[string]interface{}{"label": label, "summary": summary, "archLabel": archLabel, "parentLabel": parentLabel, "name": name})
	if err != nil {
		p.logger.Error(returnCodes.ErrFailedMarshalling, zap.Any("error", err))
		return 0, errors.New(returnCodes.ErrFailedMarshalling)
	}
	path := "channel/software/create"
	response, err := p.suse.SuseManagerCall(body, "POST", auth.Host, path, auth.SessionKey)
	if err != nil {
		p.logger.Error(returnCodes.ErrHandlingSuseManagerResponse, zap.Any("requestID", requestID), zap.Any("error", err.Error()))
		return 0, fmt.Errorf(returnCodes.ErrHandlingSuseManagerResponse)
	}
	if response.StatusCode == 200 {
		resp, err := HandleSuseManagerResponse(response.Body)
		if err != nil {
			p.logger.Error(returnCodes.ErrHandlingSuseManagerResponse, zap.Any("requestID", requestID), zap.Any("response", resp), zap.Any("error", err.Error()))
			return 0, errors.New(returnCodes.ErrHandlingSuseManagerResponse)
		}
		byteArray, _ := json.Marshal(resp)
		err = json.Unmarshal(byteArray, &resultSuc)
		if err != nil {
			p.logger.Error(returnCodes.ErrFailedUnMarshalling, zap.Any("requestID", requestID), zap.Any("error", err.Error()))
			return 0, errors.New(returnCodes.ErrFailedUnMarshalling)
		}
	} else {
		p.logger.Error(returnCodes.ErrHTTPSuseManagerResponse, zap.Any("HTTP Statuscode", response.StatusCode), zap.Any("HTTP body", response.Body))
		return 0, errors.New(returnCodes.ErrHandlingSuseManagerResponse)
	}
	p.logger.Info("Completed ChannelSoftwareCreate function", zap.Any("requestID", requestID))
	return resultSuc, nil
}

func (p *Proxy) ChannelSoftwareAssociateRepo(requestID string, auth AuthParams, channelLabel string, repoLabel string) (sumamodels.ChannelSoftwareListChildren, error) {
	p.logger.Info("Inside ChannelSoftwareAssociateRepo function", zap.Any("requestID", requestID))
	var resultSuc sumamodels.ChannelSoftwareListChildren
	body, err := json.Marshal(map[string]interface{}{"channelLabel": channelLabel, "repoLabel": repoLabel})
	if err != nil {
		p.logger.Error(returnCodes.ErrFailedMarshalling, zap.Any("error", err))
		return resultSuc, errors.New(returnCodes.ErrFailedMarshalling)
	}
	path := "channel/software/associateRepo"
	response, err := p.suse.SuseManagerCall(body, "POST", auth.Host, path, auth.SessionKey)
	if err != nil {
		p.logger.Error(returnCodes.ErrHandlingSuseManagerResponse, zap.Any("requestID", requestID), zap.Any("error", err.Error()))
		return resultSuc, fmt.Errorf(returnCodes.ErrHandlingSuseManagerResponse)
	}
	if response.StatusCode == 200 {
		resp, err := HandleSuseManagerResponse(response.Body)
		if err != nil {
			p.logger.Error(returnCodes.ErrHandlingSuseManagerResponse, zap.Any("requestID", requestID), zap.Any("response", resp), zap.Any("error", err.Error()))
			return resultSuc, errors.New(returnCodes.ErrHandlingSuseManagerResponse)
		}
		byteArray, _ := json.Marshal(resp)
		err = json.Unmarshal(byteArray, &resultSuc)
		if err != nil {
			p.logger.Error(returnCodes.ErrFailedUnMarshalling, zap.Any("requestID", requestID), zap.Any("error", err.Error()))
			return resultSuc, errors.New(returnCodes.ErrFailedUnMarshalling)
		}
	} else {
		p.logger.Error(returnCodes.ErrHTTPSuseManagerResponse, zap.Any("HTTP Statuscode", response.StatusCode), zap.Any("HTTP body", response.Body))
		return resultSuc, errors.New(returnCodes.ErrHandlingSuseManagerResponse)
	}
	p.logger.Info("Completed ChannelSoftwareAssociateRepo function", zap.Any("requestID", requestID))
	return resultSuc, nil
}

func (p *Proxy) ChannelSoftwareSyncRepo(requestID string, auth AuthParams, channelLabel string) (int, error) {
	p.logger.Info("Inside ChannelSoftwareSyncRepo function", zap.Any("requestID", requestID))
	var resultSuc int
	body, err := json.Marshal(map[string]interface{}{"channelLabel": channelLabel})
	if err != nil {
		p.logger.Error(returnCodes.ErrFailedMarshalling, zap.Any("error", err))
		return resultSuc, errors.New(returnCodes.ErrFailedMarshalling)
	}
	path := "channel/software/syncRepo"
	response, err := p.suse.SuseManagerCall(body, "POST", auth.Host, path, auth.SessionKey)
	if err != nil {
		p.logger.Error(returnCodes.ErrHandlingSuseManagerResponse, zap.Any("requestID", requestID), zap.Any("error", err.Error()))
		return resultSuc, fmt.Errorf(returnCodes.ErrHandlingSuseManagerResponse)
	}
	if response.StatusCode == 200 {
		resp, err := HandleSuseManagerResponse(response.Body)
		if err != nil {
			p.logger.Error(returnCodes.ErrHandlingSuseManagerResponse, zap.Any("requestID", requestID), zap.Any("response", resp), zap.Any("error", err.Error()))
			return resultSuc, errors.New(returnCodes.ErrHandlingSuseManagerResponse)
		}
		byteArray, _ := json.Marshal(resp)
		err = json.Unmarshal(byteArray, &resultSuc)
		if err != nil {
			p.logger.Error(returnCodes.ErrFailedUnMarshalling, zap.Any("requestID", requestID), zap.Any("error", err.Error()))
			return resultSuc, errors.New(returnCodes.ErrFailedUnMarshalling)
		}
	} else {
		p.logger.Error(returnCodes.ErrHTTPSuseManagerResponse, zap.Any("HTTP Statuscode", response.StatusCode), zap.Any("HTTP body", response.Body))
		return resultSuc, errors.New(returnCodes.ErrHandlingSuseManagerResponse)
	}
	p.logger.Info("Completed ChannelSoftwareSyncRepo function", zap.Any("requestID", requestID))
	return resultSuc, nil
}

func (p *Proxy) ChannelSoftwareIsExisting(requestID string, auth AuthParams, label string) (bool, error) {
	p.logger.Info("Inside ChannelSoftwareIsExisting function", zap.Any("requestID", requestID), zap.Any("Label", label))
	var resultSuc bool
	body, err := json.Marshal(map[string]interface{}{"channelLabel": label})
	if err != nil {
		p.logger.Error(returnCodes.ErrFailedMarshalling, zap.Any("error", err))
		return false, errors.New(returnCodes.ErrFailedMarshalling)
	}
	path := "channel/software/isExisting"
	response, err := p.suse.SuseManagerCall(body, "GET", auth.Host, path, auth.SessionKey)
	if err != nil {
		p.logger.Error(returnCodes.ErrHandlingSuseManagerResponse, zap.Any("requestID", requestID), zap.Any("error", err.Error()))
		return false, fmt.Errorf(returnCodes.ErrHandlingSuseManagerResponse)
	}
	if response.StatusCode == 200 {
		resp, err := HandleSuseManagerResponse(response.Body)
		if err != nil {
			p.logger.Error(returnCodes.ErrHandlingSuseManagerResponse, zap.Any("requestID", requestID), zap.Any("response", resp), zap.Any("error", err.Error()))
			return false, errors.New(returnCodes.ErrFailedUnMarshalling)
		}

		byteArray, _ := json.Marshal(resp)
		// p.logger.Info(byteArray)
		err = json.Unmarshal(byteArray, &resultSuc)
		if err != nil {
			p.logger.Error(returnCodes.ErrFailedUnMarshalling, zap.Any("requestID", requestID), zap.Any("error", err.Error()))
			return false, errors.New(returnCodes.ErrFailedUnMarshalling)
		}
	} else {
		p.logger.Error(returnCodes.ErrHTTPSuseManagerResponse, zap.Any("HTTP Statuscode", response.StatusCode), zap.Any("HTTP body", response.Body))
		return false, errors.New(returnCodes.ErrHandlingSuseManagerResponse)
	}
	p.logger.Info("Completed ChannelSoftwareIsExisting function", zap.Any("requestID", requestID))
	return resultSuc, nil
}
