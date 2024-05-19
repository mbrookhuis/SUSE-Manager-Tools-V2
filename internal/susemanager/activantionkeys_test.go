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

func TestProxy_ActivationKeyListActivationKeys(t *testing.T) {
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
	}
	//-------------------------------------------------------------------
	activationResp := []sumamodels.ActivationkeyGetDetails{
		{
			Key:         "Test Key",
			Description: "Test Description",
		},
	}

	success := sumamodels.RespAPISuccess{
		Success: true,
		Result:  activationResp,
	}

	activationByteArr, err := json.Marshal(success)
	if err != nil {
		panic(err)
	}

	helper := &util.HTTPHelperStruct{
		Body:       activationByteArr,
		StatusCode: 200,
		Cookies:    nil,
	}
	//---------------------------------------------------------------------
	falseResp := sumamodels.ActivationkeyPackages{
		PackageName: "test",
	}

	falseSuccess := sumamodels.RespAPISuccess{
		Success: true,
		Result:  falseResp,
	}

	falseRespByteArr, err := json.Marshal(falseSuccess)
	if err != nil {
		panic(err)
	}

	falseHelper := &util.HTTPHelperStruct{
		Body:       falseRespByteArr,
		StatusCode: 200,
		Cookies:    nil,
	}
	//---------------------------------------------------------------------
	var MessBody []string
	MessBody = append(MessBody, "Failed to get response from suse manager")

	failureSuccess := sumamodels.RespAPI{
		Success:  false,
		Result:   nil,
		Messages: MessBody,
	}

	failureByteArr, err := json.Marshal(failureSuccess)
	if err != nil {
		panic(err)
	}

	failureHelper := &util.HTTPHelperStruct{
		Body:       failureByteArr,
		StatusCode: 200,
		Cookies:    nil,
	}
	//---------------------------------------------------------------------
	errSuccesssHelper := &util.HTTPHelperStruct{
		Body:       []byte(`{"success":"abc","messages":["Error Description"]}`),
		StatusCode: 200,
		Cookies:    nil,
	}

	helperErr := &util.HTTPHelperStruct{
		Body:       nil,
		StatusCode: 400,
		Cookies:    nil,
	}

	suseManagerMock := new(MockISuseManagerAPI)
	suseManagerMock1 := new(MockISuseManagerAPI)
	suseManagerMock2 := new(MockISuseManagerAPI)
	suseManagerMock3 := new(MockISuseManagerAPI)
	suseManagerMockErr := new(MockISuseManagerAPI)
	suseManagerMockErr1 := new(MockISuseManagerAPI)

	suseManagerMock.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(helper, nil)
	suseManagerMock1.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(falseHelper, nil)
	suseManagerMock2.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(failureHelper, nil)
	suseManagerMock3.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errSuccesssHelper, nil)
	suseManagerMockErr.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(helperErr, errors.New("failed to call suse manager"))
	suseManagerMockErr1.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(helperErr, nil)

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

	positiveActivationKeyFields := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMock,
		logger:            logger,
	}

	negativeActivationKeyFields := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMockErr,
		logger:            logger,
	}

	negativeStatusCodeActivationKeyFields := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMockErr1,
		logger:            logger,
	}

	negativeFalseRespActivationKeyFields := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMock1,
		logger:            logger,
	}

	negativeFailureRespActivationKeyFields := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMock2,
		logger:            logger,
	}

	negativeErrSuccessBodyActivationKeyFields := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMock3,
		logger:            logger,
	}

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
			name:    "Activation Key List Activation Keys Positive",
			fields:  positiveActivationKeyFields,
			args:    arg,
			wantErr: false,
		},
		{
			name:    "Activation Key List Activation Keys Negative",
			fields:  negativeActivationKeyFields,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "Activation Key List Activation Keys Negative",
			fields:  negativeStatusCodeActivationKeyFields,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "Activation Key List Activation Keys False Res[] Negative",
			fields:  negativeFalseRespActivationKeyFields,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "Activation Key List Activation Keys Failure Resp Negative",
			fields:  negativeFailureRespActivationKeyFields,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "Activation Key List Activation Keys Err Success Body Negative",
			fields:  negativeErrSuccessBodyActivationKeyFields,
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
			_, err := p.ActivationKeyListActivationKeys(tt.args.requestID, tt.args.auth)
			if (err != nil) != tt.wantErr {
				t.Errorf("Proxy.ActivationKeyListActivationKeys() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestProxy_ActivationKeyGetDetails(t *testing.T) {
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
		keyName   string
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

	activationResp := sumamodels.ActivationkeyGetDetails{
		Key:         "Test Key",
		Description: "Test Description",
	}

	success := sumamodels.RespAPISuccess{
		Success: true,
		Result:  activationResp,
	}

	activationByteArr, err := json.Marshal(success)
	if err != nil {
		panic(err)
	}

	helper := &util.HTTPHelperStruct{
		Body:       activationByteArr,
		StatusCode: 200,
		Cookies:    nil,
	}

	positiveActivationKey := fields{
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

	negativeActivationKey := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMockErr,
		logger:            logger,
	}
	suseManagerMockErr.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(helperErr, errors.New("failed to call suse manager"))

	suseManagerMockErr1 := new(MockISuseManagerAPI)
	negativeActivationKeyStatusCode := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMockErr1,
		logger:            logger,
	}
	suseManagerMockErr1.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(helperErr, nil)

	suseManagerMock2 := new(MockISuseManagerAPI)
	errSuccesssHelper := &util.HTTPHelperStruct{
		Body:       []byte(`{"success":"abc","result":[{"key":"Test Key","name":"Test Description"}]}`),
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
		Body:       []byte(`{"success":true,"result":{"key":false,"name":"Test Description"}}`),
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

	arg := args{
		requestID: "test id",
		auth:      auth,
		keyName:   "Test Key",
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Activation Key Details Positive",
			fields:  positiveActivationKey,
			args:    arg,
			wantErr: false,
		},
		{
			name:    "Activation Key Details Negative",
			fields:  negativeActivationKey,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "Activation Key Details Status Code Negative",
			fields:  negativeActivationKeyStatusCode,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "Activation Key Details UnMarshal Json Negative",
			fields:  negativeActivationKeyUnMarshalJSONErr,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "Activation Key Details UnMarshal Json Resp Negative",
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
			_, err := p.ActivationKeyGetDetails(tt.args.requestID, tt.args.auth, tt.args.keyName)
			if (err != nil) != tt.wantErr {
				t.Errorf("Proxy.ActivationKeyGetDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func TestProxy_ActivationKeyRemovePackages(t *testing.T) {
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
		keyName   string
		pckgs     []sumamodels.ActivationkeyPackages
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

	var resp int = 1

	success := sumamodels.RespAPISuccess{
		Success: true,
		Result:  resp,
	}

	activationByteArr, err := json.Marshal(success)
	if err != nil {
		panic(err)
	}

	helper := &util.HTTPHelperStruct{
		Body:       activationByteArr,
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
		Body:       []byte(`{"success":"abc","result":"1"}`),
		StatusCode: 200,
		Cookies:    nil,
	}

	negativeKeyUnMarshalJSONErr := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMock2,
		logger:            logger,
	}
	suseManagerMock2.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errSuccesssHelper, nil)

	suseManagerMock3 := new(MockISuseManagerAPI)
	errUnMarshalJSONResp := &util.HTTPHelperStruct{
		Body:       []byte(`{"success":true,"result":"1"}`),
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

	arg := args{
		requestID: "test id",
		auth:      auth,
		keyName:   "Test Key",
		pckgs: []sumamodels.ActivationkeyPackages{
			{
				PackageName: "Test Package",
				ArchLabel:   "1.2.3.4",
			},
		},
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Activation Key Remove Package Positive",
			fields:  positiveKey,
			args:    arg,
			wantErr: false,
		},
		{
			name:    "Activation Key Remove Package Negative",
			fields:  negativeKey,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "Activation Key Remove Package Status Code Negative",
			fields:  negativeKeyStatusCode,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "Activation Key Remove Package Json Negative",
			fields:  negativeKeyUnMarshalJSONErr,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "Activation Key Remove Package Json Resp Negative",
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
			_, err := p.ActivationKeyRemovePackages(tt.args.requestID, tt.args.auth, tt.args.keyName, tt.args.pckgs)
			if (err != nil) != tt.wantErr {
				t.Errorf("Proxy.ActivationKeyRemovePackages() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestProxy_ActivationKeyCreate(t *testing.T) {
	type fields struct {
		cfg               *SumanConfig
		contentTypeHeader map[string]string
		suse              ISuseManagerAPI
		logger            *zap.Logger
	}
	logger := logger.Default("info")
	type args struct {
		requestID   string
		auth        AuthParams
		keyName     string
		baseChannel string
		entitlement []string
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

	var resp string = "key"

	success := sumamodels.RespAPISuccess{
		Success: true,
		Result:  resp,
	}

	activationByteArr, err := json.Marshal(success)
	if err != nil {
		panic(err)
	}

	helper := &util.HTTPHelperStruct{
		Body:       activationByteArr,
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
		Body:       []byte(`{"success":"abc","result":1}`),
		StatusCode: 200,
		Cookies:    nil,
	}

	negativeKeyUnMarshalJSONErr := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMock2,
		logger:            logger,
	}
	suseManagerMock2.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errSuccesssHelper, nil)

	suseManagerMock3 := new(MockISuseManagerAPI)
	errUnMarshalJSONResp := &util.HTTPHelperStruct{
		Body:       []byte(`{"success":true,"result":true}`),
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
	var entitlement []string
	_ = append(entitlement, "monitoring_entitled")
	arg := args{
		requestID:   "test id",
		auth:        auth,
		keyName:     "Test Key",
		baseChannel: "Base Channel",
		entitlement: entitlement,
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Activation Key Create",
			fields:  positiveKey,
			args:    arg,
			wantErr: false,
		},
		{
			name:    "Activation Key Create Negative",
			fields:  negativeKey,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "Activation Key Create Status Code Negative",
			fields:  negativeKeyStatusCode,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "Activation Key Create Json Negative",
			fields:  negativeKeyUnMarshalJSONErr,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "Activation Key Create Json Resp Negative",
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
			_, err := p.ActivationKeyCreate(tt.args.requestID, tt.args.auth, tt.args.keyName, tt.args.baseChannel, tt.args.entitlement)
			if (err != nil) != tt.wantErr {
				t.Errorf("Proxy.ActivationKeyCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestProxy_ActivationKeyAddChildChannels(t *testing.T) {
	type fields struct {
		cfg               *SumanConfig
		contentTypeHeader map[string]string
		suse              ISuseManagerAPI
		logger            *zap.Logger
	}
	logger := logger.Default("info")
	type args struct {
		requestID     string
		auth          AuthParams
		keyName       string
		childChannels []string
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

	var resp int = 1

	success := sumamodels.RespAPISuccess{
		Success: true,
		Result:  resp,
	}

	activationByteArr, err := json.Marshal(success)
	if err != nil {
		panic(err)
	}

	helper := &util.HTTPHelperStruct{
		Body:       activationByteArr,
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
		Body:       []byte(`{"success":"abc","result":"1"}`),
		StatusCode: 200,
		Cookies:    nil,
	}

	negativeKeyUnMarshalJSONErr := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMock2,
		logger:            logger,
	}
	suseManagerMock2.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errSuccesssHelper, nil)

	suseManagerMock3 := new(MockISuseManagerAPI)
	errUnMarshalJSONResp := &util.HTTPHelperStruct{
		Body:       []byte(`{"success":true,"result":"1"}`),
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

	arg := args{
		requestID:     "test id",
		auth:          auth,
		keyName:       "Test Key",
		childChannels: []string{"Child Channel1", "Child Channel2"},
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Activation Key Add Child Channels Positive",
			fields:  positiveKey,
			args:    arg,
			wantErr: false,
		},
		{
			name:    "Activation Key Add Child Channels Negative",
			fields:  negativeKey,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "Activation Key Add Child Channels Status Code Negative",
			fields:  negativeKeyStatusCode,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "Activation Key Add Child Channels Json Negative",
			fields:  negativeKeyUnMarshalJSONErr,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "Activation Key Add Child Channels Json Resp Negative",
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
			_, err := p.ActivationKeyAddChildChannels(tt.args.requestID, tt.args.auth, tt.args.keyName, tt.args.childChannels)
			if (err != nil) != tt.wantErr {
				t.Errorf("Proxy.ActivationKeyAddChildChannels() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func TestProxy_ActivationKeyAddServerGroups(t *testing.T) {
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
		keyName   string
		groups    []int
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

	var resp int = 1

	success := sumamodels.RespAPISuccess{
		Success: true,
		Result:  resp,
	}

	activationByteArr, err := json.Marshal(success)
	if err != nil {
		panic(err)
	}

	helper := &util.HTTPHelperStruct{
		Body:       activationByteArr,
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
		Body:       []byte(`{"success":"abc","result":"1"}`),
		StatusCode: 200,
		Cookies:    nil,
	}

	negativeKeyUnMarshalJSONErr := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMock2,
		logger:            logger,
	}
	suseManagerMock2.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errSuccesssHelper, nil)

	suseManagerMock3 := new(MockISuseManagerAPI)
	errUnMarshalJSONResp := &util.HTTPHelperStruct{
		Body:       []byte(`{"success":true,"result":"1"}`),
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

	arg := args{
		requestID: "test id",
		auth:      auth,
		keyName:   "Test Key",
		groups:    []int{1, 2},
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Activation Key Add Server Groups Positive",
			fields:  positiveKey,
			args:    arg,
			wantErr: false,
		},
		{
			name:    "Activation Key Add Server Groups Negative",
			fields:  negativeKey,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "Activation Key Add Server Groups Status Code Negative",
			fields:  negativeKeyStatusCode,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "Activation Key Add Server Groups Json Negative",
			fields:  negativeKeyUnMarshalJSONErr,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "Activation Key Add Server Groups Json Resp Negative",
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
			_, err := p.ActivationKeyAddServerGroups(tt.args.requestID, tt.args.auth, tt.args.keyName, tt.args.groups)
			if (err != nil) != tt.wantErr {
				t.Errorf("Proxy.ActivationKeyAddServerGroups() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestProxy_ActivationKeyDelete(t *testing.T) {
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
		keyName   string
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

	var resp int = 1

	success := sumamodels.RespAPISuccess{
		Success: true,
		Result:  resp,
	}

	activationByteArr, err := json.Marshal(success)
	if err != nil {
		panic(err)
	}

	helper := &util.HTTPHelperStruct{
		Body:       activationByteArr,
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
		Body:       []byte(`{"success":"abc","result":"1"}`),
		StatusCode: 200,
		Cookies:    nil,
	}

	negativeKeyUnMarshalJSONErr := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMock2,
		logger:            logger,
	}
	suseManagerMock2.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errSuccesssHelper, nil)

	suseManagerMock3 := new(MockISuseManagerAPI)
	errUnMarshalJSONResp := &util.HTTPHelperStruct{
		Body:       []byte(`{"success":true,"result":"1"}`),
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

	arg := args{
		requestID: "test id",
		auth:      auth,
		keyName:   "Test Key",
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Activation Key Delete Positive",
			fields:  positiveKey,
			args:    arg,
			wantErr: false,
		},
		{
			name:    "Activation Key Delete Negative",
			fields:  negativeKey,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "Activation Key Delete Status Code Negative",
			fields:  negativeKeyStatusCode,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "Activation Key Delete Json Negative",
			fields:  negativeKeyUnMarshalJSONErr,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "Activation Key Delete Json Resp Negative",
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
			_, err := p.ActivationKeyDelete(tt.args.requestID, tt.args.auth, tt.args.keyName)
			if (err != nil) != tt.wantErr {
				t.Errorf("Proxy.ActivationKeyDelete() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
