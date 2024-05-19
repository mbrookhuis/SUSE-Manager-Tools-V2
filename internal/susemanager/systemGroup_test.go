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

func TestProxy_SystemGroupGetDetails(t *testing.T) {
	type fields struct {
		cfg               *SumanConfig
		contentTypeHeader map[string]string
		suse              ISuseManagerAPI
		logger            *zap.Logger
	}
	logger := logger.Default("info")
	type args struct {
		requestID string
		auth      AuthParams
		groupName string
	}

	header := make(map[string]string)
	header["Content-Type"] = "application/json"

	config := &SumanConfig{
		Host:     "test host",
		Password: "test",
		Insecure: true,
		Login:    "test",
	}

	suseManagerMock := new(MockISuseManagerAPI)

	channelResp := sumamodels.SystemGroupGetDetails{
		ID:          123,
		Name:        "New Label",
		Description: "Test Description",
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
		logger:            logger,
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
		logger:            logger,
	}
	suseManagerMockErr.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(helperErr, errors.New("failed to call suse manager"))

	suseManagerMockErr1 := new(MockISuseManagerAPI)
	negativeKeyStatusCode := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMockErr1,
		logger:            logger,
	}
	suseManagerMockErr1.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(helperErr, nil)

	suseManagerMock2 := new(MockISuseManagerAPI)
	errSuccesssHelper := &util.HTTPHelperStruct{
		Body:       []byte(`{"success":"abc","result":{"id":123,"name":"Test Name"}}`),
		StatusCode: 200,
		Cookies:    nil,
	}

	negativeActivationKeyUnMarshalJSONErr := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMock2,
		logger:            logger,
	}
	suseManagerMock2.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errSuccesssHelper, nil)

	suseManagerMock3 := new(MockISuseManagerAPI)
	errUnMarshalJSONResp := &util.HTTPHelperStruct{
		Body:       []byte(`{"success":true,"result":{"id":false,"name":"Test Name"}}`),
		StatusCode: 200,
		Cookies:    nil,
	}

	negativeActivationKeyUnMarshalJSONRespErr := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMock3,
		logger:            logger,
	}
	suseManagerMock3.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errUnMarshalJSONResp, nil)

	auth := AuthParams{
		SessionKey: "test key",
		Host:       "test Hostname",
	}

	arg := args{
		requestID: "test id",
		auth:      auth,
		groupName: "test group",
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "System Group Get Details Positive",
			fields:  positiveKey,
			args:    arg,
			wantErr: false,
		},
		{
			name:    "System Group Get Details Negative",
			fields:  negativeKey,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "System Group Get Details Status Code Negative",
			fields:  negativeKeyStatusCode,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "System Group Get Details UnMarshal Json Negative",
			fields:  negativeActivationKeyUnMarshalJSONErr,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "System Group Get Details UnMarshal Json Resp Negative",
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
				logger:            tt.fields.logger,
			}
			_, err := p.SystemGroupGetDetails(tt.args.requestID, tt.args.auth, tt.args.groupName)
			if (err != nil) != tt.wantErr {
				t.Errorf("Proxy.SystemGroupGetDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestProxy_SystemGroupListSystemsMinimal(t *testing.T) {
	type fields struct {
		cfg               *SumanConfig
		contentTypeHeader map[string]string
		suse              ISuseManagerAPI
		logger            *zap.Logger
	}
	logger := logger.Default("info")
	type args struct {
		requestID string
		auth      AuthParams
		groupName string
	}

	header := make(map[string]string)
	header["Content-Type"] = "application/json"

	config := &SumanConfig{
		Host:     "test host",
		Password: "test",
		Insecure: true,
		Login:    "test",
	}

	suseManagerMock := new(MockISuseManagerAPI)

	resp := []byte(`{"success":true,"result":[{"id":123,"last_checkin":"Aug 26, 2022, 04:06:32 PM","created":"Aug 26, 2022, 04:06:32 PM","last_boot":"Aug 26, 2022, 04:06:32 PM","name":"Test Name","outdated_pkg_count":0,"extra_pkg_count":0}]}`)

	helper := &util.HTTPHelperStruct{
		Body:       resp,
		StatusCode: 200,
		Cookies:    nil,
	}

	positiveKey := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMock,
		logger:            logger,
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
		logger:            logger,
	}
	suseManagerMockErr.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(helperErr, errors.New("failed to call suse manager"))

	suseManagerMock2 := new(MockISuseManagerAPI)
	errSuccesssHelper := &util.HTTPHelperStruct{
		Body:       []byte(`{"success":"abc","result":[{"id":123,"name":"Test name"}]}`),
		StatusCode: 200,
		Cookies:    nil,
	}

	negativeActivationKeyUnMarshalJSONErr := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMock2,
		logger:            logger,
	}
	suseManagerMock2.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errSuccesssHelper, nil)

	suseManagerMock3 := new(MockISuseManagerAPI)
	errUnMarshalJSONResp := &util.HTTPHelperStruct{
		Body:       []byte(`{"success":true,"result":[{"id":false,"name":"Test name"}]}`),
		StatusCode: 200,
		Cookies:    nil,
	}

	negativeActivationKeyUnMarshalJSONRespErr := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMock3,
		logger:            logger,
	}
	suseManagerMock3.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errUnMarshalJSONResp, nil)

	suseManagerMockErr1 := new(MockISuseManagerAPI)
	negativeKeyStatusCode := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMockErr1,
		logger:            logger,
	}
	suseManagerMockErr1.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(helperErr, nil)

	auth := AuthParams{
		SessionKey: "test key",
		Host:       "test Hostname",
	}

	arg := args{
		requestID: "test id",
		auth:      auth,
		groupName: "test group Name",
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "System Group List Systems Minimal Positive",
			fields:  positiveKey,
			args:    arg,
			wantErr: false,
		},
		{
			name:    "System Group List Systems Minimal Negative",
			fields:  negativeKey,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "Sync Master Set CaCert Status Code Negative",
			fields:  negativeKeyStatusCode,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "System Group List Systems Minimal UnMarshal Json Negative",
			fields:  negativeActivationKeyUnMarshalJSONErr,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "System Group List Systems Minimal UnMarshal Json Resp Negative",
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
				logger:            tt.fields.logger,
			}
			_, err := p.SystemGroupListSystemsMinimal(tt.args.requestID, tt.args.auth, tt.args.groupName)
			if (err != nil) != tt.wantErr {
				t.Errorf("Proxy.SystemGroupListSystemsMinimal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
