// Package susemanager - this is a collection of api call for SUSE Manager
package susemanager

import (
	"encoding/json"
	"net/http"

	returnCodes "ecp-golang-cm/pkg/util/returnCodes"
	"github.com/pkg/errors"

	"go.uber.org/zap"
)

// GetSystemFormulaData - get formula data from system
//
// param: auth
// param: sid
// param: formulaname
// return:
func (p *Proxy) GetSystemFormulaData(auth AuthParams, sid int, formulaName string) (interface{}, error) {
	var resp interface{}
	body, err := json.Marshal(map[string]any{"systemId": sid, "formulaName": formulaName})
	if err != nil {
		p.logger.Error(returnCodes.ErrFailedMarshalling, zap.Any("error", err))
		return nil, errors.New(returnCodes.ErrFailedMarshalling)
	}
	path := "formula/getSystemFormulaData"
	response, err := p.suse.SuseManagerCall(body, http.MethodGet, auth.Host, path, auth.SessionKey)
	if err != nil {
		return nil, err
	}
	if response.StatusCode == 200 {
		resp, err = HandleSuseManagerResponse(response.Body)
		if err != nil {
			p.logger.Error(returnCodes.ErrHandlingSuseManagerResponse, zap.Any("response", resp), zap.Any("error", err.Error()))
			return nil, errors.New(returnCodes.ErrHandlingSuseManagerResponse)
		}
	} else {
		p.logger.Error(returnCodes.ErrHTTPSuseManagerResponse, zap.Any("HTTP Statuscode", response.StatusCode), zap.Any("HTTP body", response.Body))
		return nil, errors.New(returnCodes.ErrHTTPSuseManagerResponse)
	}
	p.logger.Debug("response from api", zap.Any("api", "GetSystemFormulaData"), zap.Any("response", resp))
	return resp, nil

}

// GetGroupFormulaData - get formula data from group
//
// param: auth
// param: groupID
// param: formulaName
// return:
func (p *Proxy) GetGroupFormulaData(auth AuthParams, groupID int, formulaName string) (interface{}, error) {
	body, err := json.Marshal(map[string]any{"groupId": groupID, "formulaName": formulaName})
	if err != nil {
		p.logger.Error(returnCodes.ErrFailedMarshalling, zap.Any("error", err))
		return nil, errors.New(returnCodes.ErrFailedMarshalling)
	}
	path := "formula/getGroupFormulaData"
	response, err := p.suse.SuseManagerCall(body, http.MethodGet, auth.Host, path, auth.SessionKey)
	if err != nil {
		return nil, err
	}
	var resp interface{}
	if response.StatusCode == 200 {
		resp, err = HandleSuseManagerResponse(response.Body)
		if err != nil {
			p.logger.Error(returnCodes.ErrHandlingSuseManagerResponse, zap.Any("response", resp), zap.Any("error", err.Error()))
			return nil, errors.New(returnCodes.ErrHandlingSuseManagerResponse)
		}
	} else {
		p.logger.Error(returnCodes.ErrHTTPSuseManagerResponse, zap.Any("HTTP Statuscode", response.StatusCode), zap.Any("HTTP body", response.Body))
		return nil, errors.New(returnCodes.ErrHTTPSuseManagerResponse)
	}
	return resp, nil

}

// SetSystemFormulaData - save formula data to system
//
// param: requestID
// param: auth
// param: systemID
// param: formulaName
// param: formulaData
// return:
func (p *Proxy) SetSystemFormulaData(requestID string, auth AuthParams, systemID int, formulaName string, formulaData interface{}) (int, error) {
	body, err := json.Marshal(map[string]any{"systemId": systemID, "formulaName": formulaName, "content": formulaData})
	if err != nil {
		p.logger.Error(returnCodes.ErrFailedMarshalling, zap.Any("error", err))
		return 0, errors.New(returnCodes.ErrFailedMarshalling)
	}
	path := "formula/setSystemFormulaData"
	response, err := p.suse.SuseManagerCall(body, http.MethodPost, auth.Host, path, auth.SessionKey)
	if err != nil {
		p.logger.Error("Error message recieved from suse-manger", zap.Any("requestID", requestID), zap.Any("error", err.Error()))
		return 0, err
	}
	var result int
	if response.StatusCode == 200 {
		resp, err := HandleSuseManagerResponse(response.Body)
		if err != nil {
			p.logger.Error(returnCodes.ErrHandlingSuseManagerResponse, zap.Any("requestID", requestID), zap.Any("response", resp), zap.Any("error", err.Error()))
			return 0, errors.New(returnCodes.ErrHandlingSuseManagerResponse)
		}
		byteArray, _ := json.Marshal(resp)
		err = json.Unmarshal(byteArray, &result)
		if err != nil {
			p.logger.Error(returnCodes.ErrFailedUnMarshalling, zap.Any("requestID", requestID), zap.Any("error", err.Error()))
			return 0, errors.New(returnCodes.ErrFailedUnMarshalling)
		}
	} else {
		p.logger.Error(returnCodes.ErrHTTPSuseManagerResponse, zap.Any("HTTP Statuscode", response.StatusCode), zap.Any("HTTP body", response.Body))
		return 0, errors.New(returnCodes.ErrHandlingSuseManagerResponse)
	}
	return result, nil
}

// SetGroupFormulaData - set formula data for group
//
// param: requestID
// param: auth
// param: groupID
// param: formulaName
// param: formulaData
// return:
func (p *Proxy) SetGroupFormulaData(requestID string, auth AuthParams, groupID int, formulaName string, formulaData interface{}) (int, error) {
	body, err := json.Marshal(map[string]interface{}{"groupId": groupID, "formulaName": formulaName, "content": formulaData})
	if err != nil {
		p.logger.Error(returnCodes.ErrFailedMarshalling, zap.Any("error", err))
		return 0, errors.New(returnCodes.ErrFailedMarshalling)
	}
	path := "formula/setGroupFormulaData"
	response, err := p.suse.SuseManagerCall(body, http.MethodPost, auth.Host, path, auth.SessionKey)
	if err != nil {
		return 0, err
	}
	var result int
	if response.StatusCode == 200 {
		resp, err := HandleSuseManagerResponse(response.Body)
		if err != nil {
			p.logger.Error(returnCodes.ErrHandlingSuseManagerResponse, zap.Any("requestID", requestID), zap.Any("response", resp), zap.Any("error", err.Error()))
			return 0, errors.New(returnCodes.ErrHandlingSuseManagerResponse)
		}
		byteArray, _ := json.Marshal(resp)
		err = json.Unmarshal(byteArray, &result)
		if err != nil {
			p.logger.Error(returnCodes.ErrFailedUnMarshalling, zap.Any("requestID", requestID), zap.Any("error", err.Error()))
			return 0, errors.New(returnCodes.ErrFailedUnMarshalling)
		}
	} else {
		p.logger.Error(returnCodes.ErrHTTPSuseManagerResponse, zap.Any("HTTP Statuscode", response.StatusCode), zap.Any("HTTP body", response.Body))
		return 0, errors.New(returnCodes.ErrHandlingSuseManagerResponse)
	}
	return result, nil
}

// GetFormulasByServerID -  get list of formulas for system
//
// param: requestID
// param: auth
// param: systemID
// return:
func (p *Proxy) GetFormulasByServerID(requestID string, auth AuthParams, systemID int) ([]string, error) {
	body, err := json.Marshal(map[string]interface{}{"sid": systemID})
	if err != nil {
		p.logger.Error(returnCodes.ErrFailedMarshalling, zap.Any("error", err))
		return nil, errors.New(returnCodes.ErrFailedMarshalling)
	}
	path := "formula/getFormulasByServerId"
	response, err := p.suse.SuseManagerCall(body, http.MethodGet, auth.Host, path, auth.SessionKey)
	if err != nil {
		return nil, err
	}
	var result []string
	if response.StatusCode == 200 {
		resp, err := HandleSuseManagerResponse(response.Body)
		if err != nil {
			p.logger.Error(returnCodes.ErrHandlingSuseManagerResponse, zap.Any("requestID", requestID), zap.Any("response", resp), zap.Any("error", err.Error()))
			return nil, errors.New(returnCodes.ErrHandlingSuseManagerResponse)
		}
		byteArray, _ := json.Marshal(resp)
		err = json.Unmarshal(byteArray, &result)
		if err != nil {
			p.logger.Error(returnCodes.ErrFailedUnMarshalling, zap.Any("requestID", requestID), zap.Any("error", err.Error()))
			return nil, errors.New(returnCodes.ErrFailedUnMarshalling)
		}
	} else {
		p.logger.Error(returnCodes.ErrHTTPSuseManagerResponse, zap.Any("HTTP Statuscode", response.StatusCode), zap.Any("HTTP body", response.Body))
		return nil, errors.New(returnCodes.ErrHandlingSuseManagerResponse)
	}
	return result, nil
}

// GetFormulasByGroupID
//
// param: requestID
// param: auth
// param: groupID
// return: string list of formulars, error
func (p *Proxy) GetFormulasByGroupID(requestID string, auth AuthParams, groupID int) ([]string, error) {
	body, err := json.Marshal(map[string]interface{}{"systemGroupId": groupID})
	if err != nil {
		p.logger.Error(returnCodes.ErrFailedMarshalling, zap.Any("error", err))
		return nil, errors.New(returnCodes.ErrFailedMarshalling)
	}
	path := "formula/getFormulasByGroupId"
	response, err := p.suse.SuseManagerCall(body, http.MethodGet, auth.Host, path, auth.SessionKey)
	if err != nil {
		return nil, err
	}
	var result []string
	if response.StatusCode == 200 {
		resp, err := HandleSuseManagerResponse(response.Body)
		if err != nil {
			p.logger.Error(returnCodes.ErrHandlingSuseManagerResponse, zap.Any("requestID", requestID), zap.Any("response", resp), zap.Any("error", err.Error()))
			return nil, errors.New(returnCodes.ErrHandlingSuseManagerResponse)
		}
		byteArray, _ := json.Marshal(resp)
		err = json.Unmarshal(byteArray, &result)
		if err != nil {
			p.logger.Error(returnCodes.ErrFailedUnMarshalling, zap.Any("requestID", requestID), zap.Any("error", err.Error()))
			return nil, errors.New(returnCodes.ErrFailedUnMarshalling)
		}
	} else {
		p.logger.Error(returnCodes.ErrHTTPSuseManagerResponse, zap.Any("HTTP Statuscode", response.StatusCode), zap.Any("HTTP body", response.Body))
		return nil, errors.New(returnCodes.ErrHandlingSuseManagerResponse)
	}
	return result, nil
}

// FormulaSetFormulasOfGroup - set formulas to group
//
// param: requestID
// param: auth
// param: systemID
// param: formulaNames
// return:
func (p *Proxy) FormulaSetFormulasOfGroup(requestID string, auth AuthParams, systemID int, formulaNames []string) (int, error) {
	body, err := json.Marshal(map[string]any{"systemGroupId": systemID, "formulas": formulaNames})
	if err != nil {
		p.logger.Error(returnCodes.ErrFailedMarshalling, zap.Any("error", err))
		return 0, errors.New(returnCodes.ErrFailedMarshalling)
	}
	path := "formula/setFormulasOfGroup"
	response, err := p.suse.SuseManagerCall(body, http.MethodPost, auth.Host, path, auth.SessionKey)
	if err != nil {
		p.logger.Error("Error message received from suse-manger", zap.Any("requestID", requestID), zap.Any("error", err.Error()))
		return 0, err
	}
	var result int
	if response.StatusCode == 200 {
		resp, err := HandleSuseManagerResponse(response.Body)
		if err != nil {
			p.logger.Error(returnCodes.ErrHandlingSuseManagerResponse, zap.Any("requestID", requestID), zap.Any("response", resp), zap.Any("error", err.Error()))
			return 0, errors.New(returnCodes.ErrHandlingSuseManagerResponse)
		}
		byteArray, _ := json.Marshal(resp)
		err = json.Unmarshal(byteArray, &result)
		if err != nil {
			p.logger.Error(returnCodes.ErrFailedUnMarshalling, zap.Any("requestID", requestID), zap.Any("error", err.Error()))
			return 0, errors.New(returnCodes.ErrFailedUnMarshalling)
		}
	} else {
		p.logger.Error(returnCodes.ErrHTTPSuseManagerResponse, zap.Any("HTTP Statuscode", response.StatusCode), zap.Any("HTTP body", response.Body))
		return 0, errors.New(returnCodes.ErrHandlingSuseManagerResponse)
	}
	return result, nil
}

// FormulaSetFormulasOfSystem -  set formulas for system
//
// param: requestID
// param: auth
// param: systemID
// param: formulaNames
// return:
func (p *Proxy) FormulaSetFormulasOfSystem(requestID string, auth AuthParams, systemID int, formulaNames []string) (int, error) {
	body, err := json.Marshal(map[string]any{"sid": systemID, "formulas": formulaNames})
	if err != nil {
		p.logger.Error(returnCodes.ErrFailedMarshalling, zap.Any("error", err))
		return 0, errors.New(returnCodes.ErrFailedMarshalling)
	}
	path := "formula/setFormulasOfServer"
	response, err := p.suse.SuseManagerCall(body, http.MethodPost, auth.Host, path, auth.SessionKey)
	if err != nil {
		p.logger.Error("Error message received from suse-manger", zap.Any("requestID", requestID), zap.Any("error", err.Error()))
		return 0, err
	}
	var result int
	if response.StatusCode == 200 {
		resp, err := HandleSuseManagerResponse(response.Body)
		if err != nil {
			p.logger.Error(returnCodes.ErrHandlingSuseManagerResponse, zap.Any("requestID", requestID), zap.Any("response", resp), zap.Any("error", err.Error()))
			return 0, errors.New(returnCodes.ErrHandlingSuseManagerResponse)
		}
		byteArray, _ := json.Marshal(resp)
		err = json.Unmarshal(byteArray, &result)
		if err != nil {
			p.logger.Error(returnCodes.ErrFailedUnMarshalling, zap.Any("requestID", requestID), zap.Any("error", err.Error()))
			return 0, errors.New(returnCodes.ErrFailedUnMarshalling)
		}
	} else {
		p.logger.Error(returnCodes.ErrHTTPSuseManagerResponse, zap.Any("HTTP Statuscode", response.StatusCode), zap.Any("HTTP body", response.Body))
		return 0, errors.New(returnCodes.ErrHandlingSuseManagerResponse)
	}
	return result, nil
}
