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

func TestProxy_ChannelListSoftwareChannels(t *testing.T) {
	type fields struct {
		cfg               *SumanConfig
		contentTypeHeader map[string]string
		suse              ISuseManagerAPI
		logger            *zap.Logger
	}
	log := logger.Default("info")
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

	channelResp := []sumamodels.ChannelListSoftwareChannels{
		{
			Name:        "Test Channel",
			Label:       "New Label",
			ParentLabel: "Paremt Label",
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
		logger:            log,
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
		logger:            log,
	}
	suseManagerMockErr.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(helperErr, errors.New("failed to call suse manager"))

	suseManagerMockErr1 := new(MockISuseManagerAPI)
	negativeActivationKeyStatusCode := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMockErr1,
		logger:            log,
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
		logger:            log,
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
		logger:            log,
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
			name:    "Channel List Software Channels Positive",
			fields:  positiveKey,
			args:    arg,
			wantErr: false,
		},
		{
			name:    "Channel List Software Channels Negative",
			fields:  negativeKey,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "Channel List Software Channels Status Code Negative",
			fields:  negativeActivationKeyStatusCode,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "Channel List Software Channels UnMarshal Json Negative",
			fields:  negativeActivationKeyUnMarshalJSONErr,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "Channel List Software Channels UnMarshal Json Resp Negative",
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
			_, err := p.ChannelListSoftwareChannels(tt.args.requestID, tt.args.auth)
			if (err != nil) != tt.wantErr {
				t.Errorf("Proxy.ChannelListSoftwareChannels() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestProxy_ChannelSoftwareListChildren(t *testing.T) {
	type fields struct {
		cfg               *SumanConfig
		contentTypeHeader map[string]string
		suse              ISuseManagerAPI
		logger            *zap.Logger
	}
	log := logger.Default("info")
	type args struct {
		requestID string
		auth      AuthParams
		label     string
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

	resp := []byte(`{"success":true,"result":[{"id":0,"name":"Test Channel","label":"New Label","arch_label":"x64","summary":"test","description":"test",
	"last_modified":"Aug 26, 2022, 04:06:32 PM","yumrepo_last_sync":"Aug 26, 2022, 04:06:32 PM"}]}`)

	helper := &util.HTTPHelperStruct{
		Body:       resp,
		StatusCode: 200,
		Cookies:    nil,
	}

	positiveKey := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMock,
		logger:            log,
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
		logger:            log,
	}
	suseManagerMockErr.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(helperErr, errors.New("failed to call suse manager"))

	suseManagerMockErr1 := new(MockISuseManagerAPI)
	negativeKeyStatusCode := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMockErr1,
		logger:            log,
	}
	suseManagerMockErr1.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(helperErr, nil)

	suseManagerMock2 := new(MockISuseManagerAPI)
	errSuccesssHelper := &util.HTTPHelperStruct{
		Body:       []byte(`{"success":"abc","result":[{"name":"testName","label":"Test Label"}]}`),
		StatusCode: 200,
		Cookies:    nil,
	}

	negativeKeyUnMarshalJSONErr := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMock2,
		logger:            log,
	}
	suseManagerMock2.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errSuccesssHelper, nil)

	suseManagerMock3 := new(MockISuseManagerAPI)
	errUnMarshalJSONResp := &util.HTTPHelperStruct{
		Body:       []byte(`{"success":true,"result":[{"name":false,"label":"Test Label"}]}`),
		StatusCode: 200,
		Cookies:    nil,
	}

	negativeKeyUnMarshalJSONRespErr := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMock3,
		logger:            log,
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
			name:    "Channel Software List Children Positive",
			fields:  positiveKey,
			args:    arg,
			wantErr: false,
		},
		{
			name:    "Channel Software List Children Negative",
			fields:  negativeKey,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "Channel Software List Children Status Code Negative",
			fields:  negativeKeyStatusCode,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "CChannel Software List Children UnMarshal Json Negative",
			fields:  negativeKeyUnMarshalJSONErr,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "Channel Software List Children UnMarshal Json Resp Negative",
			fields:  negativeKeyUnMarshalJSONRespErr,
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
			_, err := p.ChannelSoftwareListChildren(tt.args.requestID, tt.args.auth, tt.args.label)
			if (err != nil) != tt.wantErr {
				t.Errorf("Proxy.ChannelSoftwareListChildren() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestProxy_ChannelSoftwareCreateRepo(t *testing.T) {
	type fields struct {
		cfg               *SumanConfig
		contentTypeHeader map[string]string
		suse              ISuseManagerAPI
		logger            *zap.Logger
		retrycount        int
	}
	type args struct {
		requestID string
		auth      AuthParams
		label     string
		typeRepo  string
		url       string
	}
	log := logger.Default("info")
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

	channelResp := sumamodels.ChannelSoftwareCreateRepo{
		SourceURL:         "file:///srv/repos/testmb",
		ID:                561,
		Label:             "testmb",
		Type:              "yum",
		HasSignedMetadata: false,
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
		logger:            log,
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
		logger:            log,
	}
	suseManagerMockErr.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(helperErr, errors.New("failed to call suse manager"))

	suseManagerMockErr1 := new(MockISuseManagerAPI)
	negativeKeyStatusCode := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMockErr1,
		logger:            log,
	}
	suseManagerMockErr1.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(helperErr, nil)

	suseManagerMock2 := new(MockISuseManagerAPI)
	errSuccesssHelper := &util.HTTPHelperStruct{
		Body:       []byte(`{"success":"abc","result":[{"name":"testName","label":"Test Label"}]}`),
		StatusCode: 200,
		Cookies:    nil,
	}

	negativeKeyUnMarshalJSONErr := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMock2,
		logger:            log,
	}
	suseManagerMock2.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errSuccesssHelper, nil)

	suseManagerMock3 := new(MockISuseManagerAPI)
	errUnMarshalJSONResp := &util.HTTPHelperStruct{
		Body:       []byte(`{"success":true,"result":[{"name":false,"label":"Test Label"}]}`),
		StatusCode: 200,
		Cookies:    nil,
	}

	negativeKeyUnMarshalJSONRespErr := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMock3,
		logger:            log,
	}
	suseManagerMock3.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errUnMarshalJSONResp, nil)

	arg := args{
		requestID: "test id",
		auth:      auth,
		label:     "testextra",
		typeRepo:  "yum",
		url:       "file:///srv/repos/testextra",
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Positive",
			fields:  positiveKey,
			args:    arg,
			wantErr: false,
		},
		{
			name:    " Negative",
			fields:  negativeKey,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "Status Code Negative",
			fields:  negativeKeyStatusCode,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "UnMarshal Json Negative",
			fields:  negativeKeyUnMarshalJSONErr,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "UnMarshal Json Resp Negative",
			fields:  negativeKeyUnMarshalJSONRespErr,
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
				retrycount:        tt.fields.retrycount,
			}
			_, err := p.ChannelSoftwareCreateRepo(tt.args.requestID, tt.args.auth, tt.args.label, tt.args.typeRepo, tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChannelSoftwareCreateRepo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestProxy_ChannelSoftwareCreate(t *testing.T) {
	type fields struct {
		cfg               *SumanConfig
		contentTypeHeader map[string]string
		suse              ISuseManagerAPI
		logger            *zap.Logger
		retrycount        int
	}
	type args struct {
		requestID   string
		auth        AuthParams
		label       string
		name        string
		summary     string
		archLabel   string
		parentLabel string
	}

	log := logger.Default("info")
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

	var channelResp int = 1

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
		logger:            log,
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
		logger:            log,
	}
	suseManagerMockErr.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(helperErr, errors.New("failed to call suse manager"))

	suseManagerMockErr1 := new(MockISuseManagerAPI)
	negativeKeyStatusCode := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMockErr1,
		logger:            log,
	}
	suseManagerMockErr1.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(helperErr, nil)

	suseManagerMock2 := new(MockISuseManagerAPI)
	errSuccesssHelper := &util.HTTPHelperStruct{
		Body:       []byte(`{"success":"abc","result":[{"name":"testName","label":"Test Label"}]}`),
		StatusCode: 200,
		Cookies:    nil,
	}

	negativeKeyUnMarshalJSONErr := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMock2,
		logger:            log,
	}
	suseManagerMock2.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errSuccesssHelper, nil)

	suseManagerMock3 := new(MockISuseManagerAPI)
	errUnMarshalJSONResp := &util.HTTPHelperStruct{
		Body:       []byte(`{"success":true,"result":[{"name":false,"label":"Test Label"}]}`),
		StatusCode: 200,
		Cookies:    nil,
	}

	negativeKeyUnMarshalJSONRespErr := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMock3,
		logger:            log,
	}
	suseManagerMock3.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errUnMarshalJSONResp, nil)

	arg := args{
		requestID:   "test id",
		auth:        auth,
		label:       "testextra",
		name:        "testextra",
		summary:     "testextra",
		archLabel:   "x86_64",
		parentLabel: "parent",
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Positive",
			fields:  positiveKey,
			args:    arg,
			wantErr: false,
		},
		{
			name:    " Negative",
			fields:  negativeKey,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "Status Code Negative",
			fields:  negativeKeyStatusCode,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "UnMarshal Json Negative",
			fields:  negativeKeyUnMarshalJSONErr,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "UnMarshal Json Resp Negative",
			fields:  negativeKeyUnMarshalJSONRespErr,
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
				retrycount:        tt.fields.retrycount,
			}
			_, err := p.ChannelSoftwareCreate(tt.args.requestID, tt.args.auth, tt.args.label, tt.args.name, tt.args.summary, tt.args.archLabel, tt.args.parentLabel)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChannelSoftwareCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestProxy_ChannelSoftwareSyncRepo(t *testing.T) {
	type fields struct {
		cfg               *SumanConfig
		contentTypeHeader map[string]string
		suse              ISuseManagerAPI
		logger            *zap.Logger
		retrycount        int
	}
	type args struct {
		requestID    string
		auth         AuthParams
		channelLabel string
	}
	log := logger.Default("info")
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

	var channelResp int = 1

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
		logger:            log,
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
		logger:            log,
	}
	suseManagerMockErr.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(helperErr, errors.New("failed to call suse manager"))

	suseManagerMockErr1 := new(MockISuseManagerAPI)
	negativeKeyStatusCode := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMockErr1,
		logger:            log,
	}
	suseManagerMockErr1.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(helperErr, nil)

	suseManagerMock2 := new(MockISuseManagerAPI)
	errSuccesssHelper := &util.HTTPHelperStruct{
		Body:       []byte(`{"success":"abc","result":[{"name":"testName","label":"Test Label"}]}`),
		StatusCode: 200,
		Cookies:    nil,
	}

	negativeKeyUnMarshalJSONErr := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMock2,
		logger:            log,
	}
	suseManagerMock2.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errSuccesssHelper, nil)

	suseManagerMock3 := new(MockISuseManagerAPI)
	errUnMarshalJSONResp := &util.HTTPHelperStruct{
		Body:       []byte(`{"success":true,"result":[{"name":false,"label":"Test Label"}]}`),
		StatusCode: 200,
		Cookies:    nil,
	}

	negativeKeyUnMarshalJSONRespErr := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMock3,
		logger:            log,
	}
	suseManagerMock3.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errUnMarshalJSONResp, nil)

	arg := args{
		requestID:    "test id",
		auth:         auth,
		channelLabel: "testextra",
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Positive",
			fields:  positiveKey,
			args:    arg,
			wantErr: false,
		},
		{
			name:    " Negative",
			fields:  negativeKey,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "Status Code Negative",
			fields:  negativeKeyStatusCode,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "UnMarshal Json Negative",
			fields:  negativeKeyUnMarshalJSONErr,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "UnMarshal Json Resp Negative",
			fields:  negativeKeyUnMarshalJSONRespErr,
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
				retrycount:        tt.fields.retrycount,
			}
			_, err := p.ChannelSoftwareSyncRepo(tt.args.requestID, tt.args.auth, tt.args.channelLabel)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChannelSoftwareSyncRepo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestProxy_ChannelSoftwareIsExisting(t *testing.T) {
	type fields struct {
		cfg               *SumanConfig
		contentTypeHeader map[string]string
		suse              ISuseManagerAPI
		logger            *zap.Logger
	}
	log := logger.Default("info")
	type args struct {
		requestID string
		auth      AuthParams
		label     string
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

	resp := []byte(`{"success":true,"result":true}`)

	helper := &util.HTTPHelperStruct{
		Body:       resp,
		StatusCode: 200,
		Cookies:    nil,
	}

	positiveKey := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMock,
		logger:            log,
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
		logger:            log,
	}
	suseManagerMockErr.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(helperErr, errors.New("failed to call suse manager"))

	suseManagerMockErr1 := new(MockISuseManagerAPI)
	negativeKeyStatusCode := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMockErr1,
		logger:            log,
	}
	suseManagerMockErr1.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(helperErr, nil)

	suseManagerMock2 := new(MockISuseManagerAPI)
	errSuccesssHelper := &util.HTTPHelperStruct{
		Body:       []byte(`{"success":"abc","result":false}`),
		StatusCode: 200,
		Cookies:    nil,
	}

	negativeKeyUnMarshalJSONErr := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMock2,
		logger:            log,
	}
	suseManagerMock2.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errSuccesssHelper, nil)

	suseManagerMock3 := new(MockISuseManagerAPI)
	errUnMarshalJSONResp := &util.HTTPHelperStruct{
		Body:       []byte(`{[{"name":false,"label":"Test Label"}]}`),
		StatusCode: 200,
		Cookies:    nil,
	}

	negativeKeyUnMarshalJSONRespErr := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMock3,
		logger:            log,
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
			name:    "Channel is present",
			fields:  positiveKey,
			args:    arg,
			wantErr: false,
		},
		{
			name:    "Channel not present",
			fields:  negativeKey,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "Status Code Negative",
			fields:  negativeKeyStatusCode,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "UnMarshal Json Negative",
			fields:  negativeKeyUnMarshalJSONErr,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "UnMarshal Json Resp Negative",
			fields:  negativeKeyUnMarshalJSONRespErr,
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
			_, err := p.ChannelSoftwareIsExisting(tt.args.requestID, tt.args.auth, tt.args.label)
			if (err != nil) != tt.wantErr {
				t.Errorf("Proxy.ChannelSoftwareIsExisting() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
