// Package susemanager - SUSE Manager api call and support functions
package susemanager

import (
	"encoding/json"
	"net/http"

	sumamodels "SUSE-Manager-Tools-V2/internal/models/susemanager"
	returnCodes "SUSE-Manager-Tools-V2/internal/util/returnCodes"
	"github.com/pkg/errors"

	"go.uber.org/zap"
)

// ContentManagementListProjects - list content projects
//
// param: auth
// return:
func (p *Proxy) ContentManagementListProjects(auth AuthParams) ([]sumamodels.ContentManagementListProjects, error) {
	p.logger.Debug("contentManagement.listProjects function called")
	path := "contentmanagement/listProjects"
	response, err := p.suse.SuseManagerCall(nil, http.MethodGet, auth.Host, path, auth.SessionKey)
	if err != nil {
		p.logger.Error(returnCodes.ErrHandlingSuseManagerResponse, zap.Any("error", err.Error()))
		return nil, errors.New(returnCodes.ErrHandlingSuseManagerResponse)
	}
	var result []sumamodels.ContentManagementListProjects
	if response.StatusCode == 200 {
		resp, err := HandleSuseManagerResponse(response.Body)
		if err != nil {
			p.logger.Error(returnCodes.ErrHandlingSuseManagerResponse, zap.Any("HTTP StatusCode", response.StatusCode), zap.Any("HTTP body", response.Body))
			return nil, errors.New(returnCodes.ErrHandlingSuseManagerResponse)
		}
		byteArray, _ := json.Marshal(resp)
		err = json.Unmarshal(byteArray, &result)
		if err != nil {
			p.logger.Error(returnCodes.ErrFailedUnMarshalling, zap.Any("error", err))
			return nil, errors.New(returnCodes.ErrFailedUnMarshalling)
		}
	} else {
		p.logger.Error(returnCodes.ErrHTTPSuseManagerResponse, zap.Any("HTTP Statuscode", response.StatusCode), zap.Any("HTTP body", response.Body))
		return result, errors.New(returnCodes.ErrHTTPSuseManagerResponse)
	}
	return result, nil
}

// ContentManagementCreate - create
//
// param: auth
// param: projectLabel
// param: name
// param: description
// return:
func (p *Proxy) ContentManagementCreate(auth AuthParams, projectLabel string, name string, description string) (sumamodels.ContentManagementListProjects, error) {
	p.logger.Debug("contentManagement.create")
	var result sumamodels.ContentManagementListProjects
	path := "contentmanagement/createProject"
	body, err := json.Marshal(map[string]any{"name": name, "description": description, "projectLabel": projectLabel})
	if err != nil {
		p.logger.Error(returnCodes.ErrFailedMarshalling, zap.Any("error", err.Error()))
		return result, errors.New(returnCodes.ErrFailedMarshalling)
	}
	response, err := p.suse.SuseManagerCall(body, http.MethodPost, auth.Host, path, auth.SessionKey)
	if err != nil {
		p.logger.Error(returnCodes.ErrHTTPSuseManagerResponse, zap.Any("error", err.Error()))
		return result, errors.New(returnCodes.ErrHandlingSuseManagerResponse)
	}
	if response.StatusCode == 200 {
		resp, err := HandleSuseManagerResponse(response.Body)
		if err != nil {
			p.logger.Error(returnCodes.ErrFailedMarshalling, zap.Any("error", err.Error()))
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

// ContentManagementAttachSource - list attached channels
//
// param: auth
// param: projectLabel
// param: sourceType
// param: sourceLabel
// return:
func (p *Proxy) ContentManagementAttachSource(auth AuthParams, projectLabel string, sourceType string, sourceLabel string) (sumamodels.ContentManagementSource, error) {
	p.logger.Debug("contentManagement.source")
	var result sumamodels.ContentManagementSource
	path := "contentmanagement/attachSource"
	body, err := json.Marshal(map[string]any{"projectLabel": projectLabel, "sourceType": sourceType, "sourceLabel": sourceLabel})
	if err != nil {
		p.logger.Error(returnCodes.ErrFailedMarshalling, zap.Any("error", err.Error()))
		return result, errors.New(returnCodes.ErrFailedMarshalling)
	}
	response, err := p.suse.SuseManagerCall(body, http.MethodPost, auth.Host, path, auth.SessionKey)
	if err != nil {
		p.logger.Error(returnCodes.ErrHTTPSuseManagerResponse, zap.Any("error", err.Error()))
		return result, errors.New(returnCodes.ErrHandlingSuseManagerResponse)
	}
	if response.StatusCode == 200 {
		resp, err := HandleSuseManagerResponse(response.Body)
		if err != nil {
			p.logger.Error(returnCodes.ErrFailedMarshalling, zap.Any("error", err.Error()))
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

// ContentManagementListFilters - list available filters for content management
//
// param: auth
// return:
func (p *Proxy) ContentManagementListFilters(auth AuthParams) ([]sumamodels.ContentManagementFilter, error) {
	p.logger.Debug("contentManagement.listFilters function called")
	path := "contentmanagement/listFilters"
	response, err := p.suse.SuseManagerCall(nil, http.MethodGet, auth.Host, path, auth.SessionKey)
	if err != nil {
		p.logger.Error(returnCodes.ErrHTTPSuseManagerResponse, zap.Any("error", err.Error()))
		return nil, errors.New(returnCodes.ErrHandlingSuseManagerResponse)
	}
	var result []sumamodels.ContentManagementFilter
	if response.StatusCode == 200 {
		resp, err := HandleSuseManagerResponse(response.Body)
		if err != nil {
			p.logger.Error(returnCodes.ErrFailedMarshalling, zap.Any("error", err.Error()))
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

// ContentManagementCreateFilter - create filter for content management
//
// param: auth
// param: name
// param: rule
// param: entityType
// param: criteria
// return:
func (p *Proxy) ContentManagementCreateFilter(auth AuthParams, name string, rule string, entityType string, criteria sumamodels.FilterCriteria) (sumamodels.ContentManagementFilter, error) {
	p.logger.Debug("contentManagement.createFilter function called")
	var result sumamodels.ContentManagementFilter
	path := "contentmanagement/createFilter"
	body, err := json.Marshal(map[string]any{"name": name, "rule": rule, "entityType": entityType, "criteria": criteria})
	if err != nil {
		p.logger.Error(returnCodes.ErrFailedMarshalling, zap.Any("error", err.Error()))
		return result, errors.New(returnCodes.ErrFailedMarshalling)
	}
	response, err := p.suse.SuseManagerCall(body, http.MethodPost, auth.Host, path, auth.SessionKey)
	if err != nil {
		p.logger.Error(returnCodes.ErrHTTPSuseManagerResponse, zap.Any("error", err.Error()))
		return result, errors.New(returnCodes.ErrHandlingSuseManagerResponse)
	}
	if response.StatusCode == 200 {
		resp, err := HandleSuseManagerResponse(response.Body)
		if err != nil {
			p.logger.Error(returnCodes.ErrFailedMarshalling, zap.Any("error", err.Error()))
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

// ContentManagementAttachFilter - attach a filter to a specific project
//
// param: auth
// param: projectLabel
// param: filterID
// return:
func (p *Proxy) ContentManagementAttachFilter(auth AuthParams, projectLabel string, filterID int) (sumamodels.ContentManagementFilter, error) {
	p.logger.Debug("contentManagement.attachFilter function called")
	var result sumamodels.ContentManagementFilter
	path := "contentmanagement/attachFilter"
	body, err := json.Marshal(map[string]any{"projectLabel": projectLabel, "filterId": filterID})
	if err != nil {
		p.logger.Error(returnCodes.ErrFailedMarshalling, zap.Any("error", err.Error()))
		return result, errors.New(returnCodes.ErrFailedMarshalling)
	}
	response, err := p.suse.SuseManagerCall(body, http.MethodPost, auth.Host, path, auth.SessionKey)
	if err != nil {
		p.logger.Error(returnCodes.ErrHTTPSuseManagerResponse, zap.Any("error", err.Error()))
		return result, errors.New(returnCodes.ErrHandlingSuseManagerResponse)
	}
	if response.StatusCode == 200 {
		resp, err := HandleSuseManagerResponse(response.Body)
		if err != nil {
			p.logger.Error(returnCodes.ErrFailedMarshalling, zap.Any("error", err.Error()))
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

// ContentManagementCreateEnvironment - create environment for a project
//
// param: auth
// param: projectLabel
// param: predecessorLabel
// param: envlabel
// param: name
// param: description
// return:
func (p *Proxy) ContentManagementCreateEnvironment(auth AuthParams, projectLabel string, predecessorLabel string, envlabel string, name string, description string) (sumamodels.ContentManagementEnvironment, error) {
	p.logger.Debug("contentManagement.createEnvironment function called")
	var result sumamodels.ContentManagementEnvironment
	path := "contentmanagement/createEnvironment"
	body, err := json.Marshal(map[string]any{"projectLabel": projectLabel, "predecessorLabel": predecessorLabel, "envLabel": envlabel, "name": name, "description": description})
	if err != nil {
		p.logger.Error(returnCodes.ErrFailedMarshalling, zap.Any("error", err.Error()))
		return result, errors.New(returnCodes.ErrFailedMarshalling)
	}
	response, err := p.suse.SuseManagerCall(body, http.MethodPost, auth.Host, path, auth.SessionKey)
	if err != nil {
		p.logger.Error(returnCodes.ErrHTTPSuseManagerResponse, zap.Any("error", err.Error()))
		return result, errors.New(returnCodes.ErrHandlingSuseManagerResponse)
	}
	if response.StatusCode == 200 {
		resp, err := HandleSuseManagerResponse(response.Body)
		if err != nil {
			p.logger.Error(returnCodes.ErrFailedMarshalling, zap.Any("error", err.Error()))
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

// ContentManagementBuildProject - build a project
//
// param: auth
// param: projectLabel
// return:
func (p *Proxy) ContentManagementBuildProject(auth AuthParams, projectLabel string) (int, error) {
	p.logger.Debug("contentManagement.buildProject function called")
	var result int
	path := "contentmanagement/buildProject"
	body, err := json.Marshal(map[string]any{"projectLabel": projectLabel})
	if err != nil {
		p.logger.Error(returnCodes.ErrFailedMarshalling, zap.Any("error", err.Error()))
		return result, errors.New(returnCodes.ErrFailedMarshalling)
	}
	response, err := p.suse.SuseManagerCall(body, http.MethodPost, auth.Host, path, auth.SessionKey)
	if err != nil {
		p.logger.Error(returnCodes.ErrFailedMarshalling, zap.Any("error", err.Error()))
		return result, errors.New(returnCodes.ErrHandlingSuseManagerResponse)
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

// ContentManagementListProjectSources
//
// param: auth
// param: project
// return:
func (p *Proxy) ContentManagementListProjectSources(auth AuthParams, project string) ([]sumamodels.ContentManagementSource, error) {
	p.logger.Debug("contentManagement.listProjectSources function called")
	var result []sumamodels.ContentManagementSource
	path := "contentmanagement/listProjectSources"
	body, err := json.Marshal(map[string]any{"projectLabel": project})
	if err != nil {
		p.logger.Error(returnCodes.ErrFailedMarshalling, zap.Any("error", err.Error()))
		return result, errors.New(returnCodes.ErrFailedMarshalling)
	}
	response, err := p.suse.SuseManagerCall(body, http.MethodGet, auth.Host, path, auth.SessionKey)
	if err != nil {
		p.logger.Error(returnCodes.ErrHandlingSuseManagerResponse, zap.Any("error", err.Error()))
		return nil, errors.New(returnCodes.ErrHandlingSuseManagerResponse)
	}
	if response.StatusCode == 200 {
		resp, err := HandleSuseManagerResponse(response.Body)
		if err != nil {
			p.logger.Error(returnCodes.ErrHandlingSuseManagerResponse, zap.Any("HTTP Statuscode", response.StatusCode), zap.Any("HTTP body", response.Body))
			return nil, errors.New(returnCodes.ErrHandlingSuseManagerResponse)
		}
		byteArray, _ := json.Marshal(resp)
		err = json.Unmarshal(byteArray, &result)
		if err != nil {
			p.logger.Error(returnCodes.ErrFailedUnMarshalling, zap.Any("error", err))
			return nil, errors.New(returnCodes.ErrFailedUnMarshalling)
		}
	} else {
		p.logger.Error(returnCodes.ErrHTTPSuseManagerResponse, zap.Any("HTTP Statuscode", response.StatusCode), zap.Any("HTTP body", response.Body))
		return result, errors.New(returnCodes.ErrHTTPSuseManagerResponse)
	}
	return result, nil
}
