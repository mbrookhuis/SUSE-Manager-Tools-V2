// Package SUSE Manager - SUSE Manager api call and support functions
package susemanager

import (
	"encoding/json"
	"errors"
	"testing"

	sumamodels "ecp-golang-cm/pkg/models/susemanager"
	"ecp-golang-cm/pkg/util/logger"
	util "ecp-golang-cm/pkg/util/rest"

	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

func TestProxy_ConfigChannelListGlobals(t *testing.T) {
	type fields struct {
		cfg               *SumanConfig
		contentTypeHeader map[string]string
		suse              ISuseManagerAPI
		loger             *zap.Logger
	}
	logger := logger.Default("info")
	type args struct {
		requestID string
		auth      AuthParams
	}

	header := make(map[string]string)
	header["Content-Type"] = "application/json"

	auth := AuthParams{
		SessionKey: "test key",
		Host:       "test Hostname",
	}
	config := &SumanConfig{
		Host:     "test host",
		Password: "test",
		Insecure: true,
		Login:    "test",
	}

	suseManagerMock := new(MockISuseManagerAPI)

	channelResp := []sumamodels.ConfigChannelListGlobals{
		{
			Name:  "Test Channel",
			Label: "New Label",
		},
	}

	success := sumamodels.RespAPISuccess{
		Success: true,
		Result:  channelResp,
	}

	channelByteArr, err := json.Marshal(success)
	if err != nil {
		panic(err)
	}

	helper := &util.HTTPHelperStruct{
		Body:       channelByteArr,
		StatusCode: 200,
		Cookies:    nil,
	}

	positiveKey := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMock,
		loger:             logger,
	}

	suseManagerMock.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(helper, nil)

	suseManagerMockErr := new(MockISuseManagerAPI)

	helperErr := &util.HTTPHelperStruct{
		Body:       nil,
		StatusCode: 400,
		Cookies:    nil,
	}

	negativeKey := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMockErr,
		loger:             logger,
	}
	suseManagerMockErr.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(helperErr, errors.New("failed to call suse manager"))

	suseManagerMockErr1 := new(MockISuseManagerAPI)
	negativeActivationKeyStatusCode := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMockErr1,
		loger:             logger,
	}
	suseManagerMockErr1.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(helperErr, nil)

	suseManagerMock2 := new(MockISuseManagerAPI)
	errSuccesssHelper := &util.HTTPHelperStruct{
		Body:       []byte(`{"success":"abc","result":[{"name":"testName","label":"Test Label"}]}`),
		StatusCode: 200,
		Cookies:    nil,
	}

	negativeActivationKeyUnMarshalJSONErr := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMock2,
		loger:             logger,
	}
	suseManagerMock2.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errSuccesssHelper, nil)

	suseManagerMock3 := new(MockISuseManagerAPI)
	errUnMarshalJSONResp := &util.HTTPHelperStruct{
		Body:       []byte(`{"success":true,"result":{"name":false,"label":"Test Label"}}`),
		StatusCode: 200,
		Cookies:    nil,
	}

	negativeActivationKeyUnMarshalJSONRespErr := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMock3,
		loger:             logger,
	}
	suseManagerMock3.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errUnMarshalJSONResp, nil)

	arg := args{
		requestID: "test id",
		auth:      auth,
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Config Channel List Globals Positive",
			fields:  positiveKey,
			args:    arg,
			wantErr: false,
		},
		{
			name:    "Config Channel List Globals Negative",
			fields:  negativeKey,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "Config Channel List Globals Status Code Negative",
			fields:  negativeActivationKeyStatusCode,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "Config Channel List Globals UnMarshal Json Negative",
			fields:  negativeActivationKeyUnMarshalJSONErr,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "Config Channel List Globals UnMarshal Json Resp Negative",
			fields:  negativeActivationKeyUnMarshalJSONRespErr,
			args:    arg,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Proxy{
				cfg:               tt.fields.cfg,
				contentTypeHeader: tt.fields.contentTypeHeader,
				suse:              tt.fields.suse,
				logger:            tt.fields.loger,
			}
			_, err := p.ConfigChannelListGlobals(tt.args.requestID, tt.args.auth)
			if (err != nil) != tt.wantErr {
				t.Errorf("Proxy.ConfigChannelListGlobals() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
