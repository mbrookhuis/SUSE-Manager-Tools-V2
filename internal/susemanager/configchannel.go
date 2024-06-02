// Package susemanager - SUSE Manager api call and support functions
package susemanager

import (
	"encoding/json"
	"fmt"

	sumamodels "SUSE-Manager-Tools-V2/internal/models/susemanager"

	"go.uber.org/zap"
)

// ConfigChannelListGlobals - list configchannels
//
// param: requestID
// param: auth
// return:
func (p *Proxy) ConfigChannelListGlobals(requestID string, auth AuthParams) ([]sumamodels.ConfigChannelListGlobals, error) {
	p.logger.Info("ConfigChannelListGlobals function call started", zap.Any("resquestId", requestID))
	path := "configchannel/listGlobals"
	response, err := p.suse.SuseManagerCall(nil, "GET", auth.Host, path, auth.SessionKey)
	if err != nil {
		p.logger.Error("Error message recieved from suse-manger", zap.Any("requestID", requestID), zap.Any("error", err.Error()))
		return nil, fmt.Errorf("error while getting list of configuration channels. Error: %s", err.Error())
	}
	var result []sumamodels.ConfigChannelListGlobals
	if response.StatusCode == 200 {
		resp, err := HandleSuseManagerResponse(response.Body)
		if err != nil {
			return nil, fmt.Errorf("error while handling suse manager response err: %s", err.Error())
		}
		byteArray, _ := json.Marshal(resp)
		err = json.Unmarshal(byteArray, &result)
		if err != nil {
			p.logger.Error("unmarshling error", zap.Any("requestID", requestID), zap.Any("error", err.Error()))
			return nil, fmt.Errorf("unable to process the received data. err: %s", err.Error())
		}
	} else {
		return nil, fmt.Errorf("fetching configuration channel list failed. Http StatusCode: %s Http Response body: %s", fmt.Sprint(response.StatusCode), fmt.Sprint(string(response.Body)))
	}
	return result, nil
}
