// Package susemanager suma api calls for kickstart.tree
package susemanager

import (
	"encoding/json"
	"net/http"

	"errors"

	sumamodels "SUSE-Manager-Tools-V2/internal/models/susemanager"
	returnCodes "SUSE-Manager-Tools-V2/internal/util/returnCodes"
	"go.uber.org/zap"
)

// KickstartTreeGetDetails - get autoinstall details
//
// param: auth
// param: distributionName
// return:
func (p *Proxy) KickstartTreeGetDetails(auth AuthParams, distributionName string) (sumamodels.KickstartTreeGetDetails, error) {
	p.logger.Debug("Kickstart.tree.getDetails called")
	var result sumamodels.KickstartTreeGetDetails
	body, err := json.Marshal(map[string]any{"treeLabel": distributionName})
	if err != nil {
		p.logger.Error(returnCodes.ErrFailedMarshalling, zap.Any("error", err.Error()))
		return result, errors.New(returnCodes.ErrFailedMarshalling)
	}
	path := "kickstart/tree/getDetails"
	response, err := p.suse.SuseManagerCall(body, http.MethodGet, auth.Host, path, auth.SessionKey)
	if err != nil {
		return result, err
	}
	if response.StatusCode == 200 {
		resp, err := HandleSuseManagerResponse(response.Body)
		if err != nil {
			p.logger.Error(returnCodes.ErrHTTPSuseManagerResponse, zap.Any("error", err.Error()))
			return result, errors.New(returnCodes.ErrHandlingSuseManagerResponse)
		}
		byteArray, _ := json.Marshal(resp)
		err = json.Unmarshal(byteArray, &result)
		if err != nil {
			p.logger.Error(returnCodes.ErrFailedUnMarshalling, zap.Any("error", err))
			return result, errors.New(returnCodes.ErrFailedUnMarshalling)
		}
	} else {
		p.logger.Error(returnCodes.ErrHTTPSuseManagerResponse, zap.Any("HTTP Statuscode", response.StatusCode), zap.Any("HTTP body", response.Body))
		return result, errors.New(returnCodes.ErrHTTPSuseManagerResponse)
	}
	return result, nil
}

// KickstartTreeCreate - create autoinstall
//
// param: auth
// param: treeLabel
// param: basePath
// param: channelLabel
// param: installType
// return:
func (p *Proxy) KickstartTreeCreate(auth AuthParams, treeLabel string, basePath string, channelLabel string, installType string) (int, error) {
	p.logger.Debug("Kickstart.tree.create called")
	var result int
	body, err := json.Marshal(map[string]any{"treeLabel": treeLabel, "basePath": basePath, "channelLabel": channelLabel, "installType": installType})
	if err != nil {
		p.logger.Error(returnCodes.ErrFailedMarshalling, zap.Any("error", err.Error()))
		return result, errors.New(returnCodes.ErrFailedMarshalling)
	}
	path := "kickstart/tree/create"
	response, err := p.suse.SuseManagerCall(body, http.MethodPost, auth.Host, path, auth.SessionKey)
	if err != nil {
		return result, err
	}
	if response.StatusCode == 200 {
		resp, err := HandleSuseManagerResponse(response.Body)
		if err != nil {
			p.logger.Error(returnCodes.ErrHTTPSuseManagerResponse, zap.Any("error", err.Error()))
			return result, errors.New(returnCodes.ErrHandlingSuseManagerResponse)
		}
		byteArray, _ := json.Marshal(resp)
		err = json.Unmarshal(byteArray, &result)
		if err != nil {
			p.logger.Error(returnCodes.ErrFailedUnMarshalling, zap.Any("error", err))
			return result, errors.New(returnCodes.ErrFailedUnMarshalling)
		}
	} else {
		p.logger.Error(returnCodes.ErrHTTPSuseManagerResponse, zap.Any("HTTP Statuscode", response.StatusCode), zap.Any("HTTP body", response.Body))
		return result, errors.New(returnCodes.ErrHTTPSuseManagerResponse)
	}
	return result, nil
}

// KickstartTreeCreateKernelOptions
//
// param: auth
// param: treeLabel
// param: basePath
// param: channelLabel
// param: installType
// param: kernelOptions
// param: postKernelOptions
// return:
func (p *Proxy) KickstartTreeCreateKernelOptions(auth AuthParams, treeLabel string, basePath string, channelLabel string, installType string, kernelOptions string, postKernelOptions string) (int, error) {
	p.logger.Debug("Kickstart.tree.create called")
	var result int
	body, err := json.Marshal(map[string]any{"treeLabel": treeLabel, "basePath": basePath, "channelLabel": channelLabel, "installType": installType, "kernelOptions": kernelOptions, "postKernelOptions": postKernelOptions})
	if err != nil {
		p.logger.Error(returnCodes.ErrFailedMarshalling, zap.Any("error", err.Error()))
		return result, errors.New(returnCodes.ErrFailedMarshalling)
	}
	path := "kickstart/tree/create"
	response, err := p.suse.SuseManagerCall(body, http.MethodPost, auth.Host, path, auth.SessionKey)
	if err != nil {
		return result, err
	}
	if response.StatusCode == 200 {
		resp, err := HandleSuseManagerResponse(response.Body)
		if err != nil {
			p.logger.Error(returnCodes.ErrHTTPSuseManagerResponse, zap.Any("error", err.Error()))
			return result, errors.New(returnCodes.ErrHandlingSuseManagerResponse)
		}
		byteArray, _ := json.Marshal(resp)
		err = json.Unmarshal(byteArray, &result)
		if err != nil {
			p.logger.Error(returnCodes.ErrFailedUnMarshalling, zap.Any("error", err))
			return result, errors.New(returnCodes.ErrFailedUnMarshalling)
		}
	} else {
		p.logger.Error(returnCodes.ErrHTTPSuseManagerResponse, zap.Any("HTTP Statuscode", response.StatusCode), zap.Any("HTTP body", response.Body))
		return result, errors.New(returnCodes.ErrHTTPSuseManagerResponse)
	}
	return result, nil
}

// KickstartImportRawFile
//
// param: auth
// param: profileLabel
// param: virtType
// param: channelLabel
// param: dataXML
// return:
func (p *Proxy) KickstartImportRawFile(auth AuthParams, profileLabel string, virtType string, channelLabel string, dataXML string) (int, error) {
	p.logger.Debug("Kickstart.importRawFile called")
	var result int
	body, err := json.Marshal(map[string]any{"profileLabel": profileLabel, "virtualizationType": virtType, "kickstartableTreeLabel": channelLabel, "kickstartFileContents": dataXML})
	if err != nil {
		p.logger.Error(returnCodes.ErrFailedMarshalling, zap.Any("error", err.Error()))
		return result, errors.New(returnCodes.ErrFailedMarshalling)
	}
	path := "kickstart/importRawFile"
	response, err := p.suse.SuseManagerCall(body, http.MethodPost, auth.Host, path, auth.SessionKey)
	if err != nil {
		return result, err
	}
	if response.StatusCode == 200 {
		resp, err := HandleSuseManagerResponse(response.Body)
		if err != nil {
			p.logger.Error(returnCodes.ErrHTTPSuseManagerResponse, zap.Any("error", err.Error()))
			return result, errors.New(returnCodes.ErrHandlingSuseManagerResponse)
		}
		byteArray, _ := json.Marshal(resp)
		err = json.Unmarshal(byteArray, &result)
		if err != nil {
			p.logger.Error(returnCodes.ErrFailedUnMarshalling, zap.Any("error", err))
			return result, errors.New(returnCodes.ErrFailedUnMarshalling)
		}
	} else {
		p.logger.Error(returnCodes.ErrHTTPSuseManagerResponse, zap.Any("HTTP Statuscode", response.StatusCode), zap.Any("HTTP body", response.Body))
		return result, errors.New(returnCodes.ErrHTTPSuseManagerResponse)
	}
	return result, nil
}

// KickstartListKickstarts
//
// param: auth
// return:
func (p *Proxy) KickstartListKickstarts(auth AuthParams) ([]sumamodels.KickstartListProfiles, error) {
	p.logger.Debug("Kickstart.listKickstarts called")
	var result []sumamodels.KickstartListProfiles
	path := "kickstart/listKickstarts"
	response, err := p.suse.SuseManagerCall(nil, http.MethodGet, auth.Host, path, auth.SessionKey)
	if err != nil {
		return result, err
	}
	if response.StatusCode == 200 {
		resp, err := HandleSuseManagerResponse(response.Body)
		if err != nil {
			p.logger.Error(returnCodes.ErrHTTPSuseManagerResponse, zap.Any("error", err.Error()))
			return result, errors.New(returnCodes.ErrHandlingSuseManagerResponse)
		}
		byteArray, _ := json.Marshal(resp)
		err = json.Unmarshal(byteArray, &result)
		if err != nil {
			p.logger.Error(returnCodes.ErrFailedUnMarshalling, zap.Any("error", err))
			return result, errors.New(returnCodes.ErrFailedUnMarshalling)
		}
	} else {
		p.logger.Error(returnCodes.ErrHTTPSuseManagerResponse, zap.Any("HTTP Statuscode", response.StatusCode), zap.Any("HTTP body", response.Body))
		return result, errors.New(returnCodes.ErrHTTPSuseManagerResponse)
	}
	return result, nil
}

// KickstartDeleteProfile
//
// param: auth
// param: profileLabel
// param: virtType
// param: channelLabel
// param: dataXML
// return:
func (p *Proxy) KickstartDeleteProfile(auth AuthParams, profileLabel string) (int, error) {
	p.logger.Debug("Kickstart.deleteProfile called")
	var result int
	body, err := json.Marshal(map[string]any{"ksLabel": profileLabel})
	if err != nil {
		p.logger.Error(returnCodes.ErrFailedMarshalling, zap.Any("error", err.Error()))
		return result, errors.New(returnCodes.ErrFailedMarshalling)
	}
	path := "kickstart/deleteProfile"
	response, err := p.suse.SuseManagerCall(body, http.MethodPost, auth.Host, path, auth.SessionKey)
	if err != nil {
		return result, err
	}
	if response.StatusCode == 200 {
		resp, err := HandleSuseManagerResponse(response.Body)
		if err != nil {
			p.logger.Error(returnCodes.ErrHTTPSuseManagerResponse, zap.Any("error", err.Error()))
			return result, errors.New(returnCodes.ErrHandlingSuseManagerResponse)
		}
		byteArray, _ := json.Marshal(resp)
		err = json.Unmarshal(byteArray, &result)
		if err != nil {
			p.logger.Error(returnCodes.ErrFailedUnMarshalling, zap.Any("error", err))
			return result, errors.New(returnCodes.ErrFailedUnMarshalling)
		}
	} else {
		p.logger.Error(returnCodes.ErrHTTPSuseManagerResponse, zap.Any("HTTP Statuscode", response.StatusCode), zap.Any("HTTP body", response.Body))
		return result, errors.New(returnCodes.ErrHTTPSuseManagerResponse)
	}
	return result, nil
}

// KickstartProfileSetVariables
//
// param: auth
// param: profileLabel
// param: profileVariables
// return:
func (p *Proxy) KickstartProfileSetVariables(auth AuthParams, profileLabel string, profileVariables interface{}) (int, error) {
	p.logger.Debug("Kickstart.profile.setVariables called", zap.Any("profileLabel", profileLabel), zap.Any("profileVariables", profileVariables))
	var result int
	body, err := json.Marshal(map[string]any{"ksLabel": profileLabel, "variables": profileVariables})
	if err != nil {
		p.logger.Error(returnCodes.ErrFailedMarshalling, zap.Any("error", err.Error()))
		return result, errors.New(returnCodes.ErrFailedMarshalling)
	}
	path := "kickstart/profile/setVariables"
	response, err := p.suse.SuseManagerCall(body, http.MethodPost, auth.Host, path, auth.SessionKey)
	if err != nil {
		return result, err
	}
	if response.StatusCode == 200 {
		resp, err := HandleSuseManagerResponse(response.Body)
		if err != nil {
			p.logger.Error(returnCodes.ErrHTTPSuseManagerResponse, zap.Any("error", err.Error()))
			return result, errors.New(returnCodes.ErrHandlingSuseManagerResponse)
		}
		byteArray, _ := json.Marshal(resp)
		err = json.Unmarshal(byteArray, &result)
		if err != nil {
			p.logger.Error(returnCodes.ErrFailedUnMarshalling, zap.Any("error", err))
			return result, errors.New(returnCodes.ErrFailedUnMarshalling)
		}
	} else {
		p.logger.Error(returnCodes.ErrHTTPSuseManagerResponse, zap.Any("HTTP Statuscode", response.StatusCode), zap.Any("HTTP body", response.Body))
		return result, errors.New(returnCodes.ErrHTTPSuseManagerResponse)
	}
	return result, nil
}
