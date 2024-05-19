// Package susemanager - SUSE Manager api call and support functions
package susemanager

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"go.uber.org/zap"
)

// implements all suse manager API calls here which act as a proxy

// SumanLogin - login to SUSE Manager
//
// param: requestID
// return:
func (p *Proxy) SumanLogin(requestID string) (string, error) {
	reqBody, _ := json.Marshal(map[string]interface{}{
		"login":    strings.TrimSpace(p.cfg.Login),
		"password": strings.TrimSpace(p.cfg.Password)})
	sumaCookie, err := p.GetSessionKey(reqBody, p.cfg.Host, requestID)
	if err != nil {
		return "", err
	}
	p.logger.Info("Successfull login to suse manager", zap.Any("requestID", requestID), zap.Any("host", p.cfg.Host))
	return sumaCookie, nil
}

// GetSessionKey - get session key
//
// param: body
// param: host
// param: requestID
// return:
func (p *Proxy) GetSessionKey(body []byte, host string, requestID string) (string, error) {
	response, err := p.suse.SuseManagerCall(body, "POST", host, "auth/login", "")
	if err != nil {
		p.logger.Error("error while login to suse manager", zap.Any("requestID", requestID), zap.Any("error", err.Error()))
		return "", errors.New("error while login to suse manager")
	}
	_, err = HandleSuseManagerResponse(response.Body)
	if err != nil {
		p.logger.Fatal("Unable to retrieve Cookie. Login problem", zap.Any("requestID", requestID), zap.Any("host", p.cfg.Host), zap.Any("Error", err.Error()))
	} else {
		p.logger.Info("succesfully retrieved Cookie.", zap.Any("requestID", requestID), zap.Any("host", p.cfg.Host))
	}
	Cookie := response.Cookies
	sumaCookie := fmt.Sprint(Cookie[2])
	return sumaCookie, nil
}

// SumanLogout - logout from SUSE Manager
//
// param: requestID
// param: auth
func (p *Proxy) SumanLogout(requestID string, auth AuthParams) error {
	path := "auth/logout"
	_, err := p.suse.SuseManagerCall(nil, "GET", auth.Host, path, auth.SessionKey)
	if err != nil {
		p.logger.Error("Unable to logout the request from suse manager", zap.Any("requestID", requestID), zap.Any("error", err.Error()))
		return errors.New("error while logout suse manager")
	}
	p.logger.Info("Successfully logout from SUSE Manager Server", zap.Any("requestID", requestID), zap.Any("host", p.cfg.Host))
	return nil
}
