// Package susemanager - SUSE Manager api call and support functions
package susemanager

import (
	"fmt"

	"SUSE-Manager-Tools-V2/internal/util/rest"

	"go.uber.org/zap"
)

// SuMaAPI description for API call
type SuMaAPI struct {
	basepath   string
	insecure   bool
	logger     *zap.Logger
	susestub   bool
	retrycount int
}

// NewSuseManagerAPI - open new API call to SUSE Manager
//
// param: basepath
// param: insecure
// param: logger
// param: retrycount
// param: susestub
// return:
func NewSuseManagerAPI(basepath string, insecure bool, logger *zap.Logger, retrycount int, susestub ...bool) ISuseManagerAPI {
	var s = &SuMaAPI{
		basepath:   basepath,
		insecure:   insecure,
		logger:     logger,
		retrycount: retrycount,
	}
	// default suse stub to false
	s.susestub = false
	if len(susestub) > 0 {
		s.susestub = susestub[0]
	}
	return s
}

// SuseManagerCall - Call api at SUSE Manager
//
// param: body
// param: method
// param: hostname
// param: path
// param: sessionKey
// return:
func (s *SuMaAPI) SuseManagerCall(body []byte, method string, hostname string, path string, sessionKey string) (output *rest.HTTPHelperStruct, er error) {
	header := make(map[string]string)
	header["Cookie"] = sessionKey
	var httproto = "https"
	if s.susestub {
		httproto = "http"
	}
	url := fmt.Sprintf("%s://%s/%s/%s", httproto, hostname, s.basepath, path)
	response, err := rest.HTTPHelper(s.logger, s.retrycount, body, method, url, s.insecure, header)
	if err != nil {
		s.logger.Error("Error message recieved", zap.Any("error", err.Error()))
		return nil, err
	}
	return response, nil
}
