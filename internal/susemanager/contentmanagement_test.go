// Package SUSE Manager - SUSE Manager api call and support functions
package susemanager

import (
	"errors"
	"testing"

	sumamodels "ecp-golang-cm/pkg/models/susemanager"
	"ecp-golang-cm/pkg/util/logger"
	util "ecp-golang-cm/pkg/util/rest"

	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

func TestProxy_ContentManagementListProjects(t *testing.T) {
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

	resp := []byte(`{"success":true,"result":[{"id":0,"label":"New Label","name":"Test Channel","description":"","lastBuildDate":"Aug 26, 2022, 04:06:32 PM","firstEnvironment":"","orgId":0}]}`)

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
	negativeResultStatusCode := fields{
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

	negativeResultUnMarshalJSONErr := fields{
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

	negativeResultUnMarshalJSONRespErr := fields{
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
			name:    "Positive",
			fields:  positiveKey,
			args:    arg,
			wantErr: false,
		},
		{
			name:    "Negative",
			fields:  negativeKey,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "Status Code Negative",
			fields:  negativeResultStatusCode,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "UnMarshal Json Negative",
			fields:  negativeResultUnMarshalJSONErr,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "UnMarshal Json Resp Negative",
			fields:  negativeResultUnMarshalJSONRespErr,
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
			_, err := p.ContentManagementListProjects(tt.args.requestID, tt.args.auth)
			if (err != nil) != tt.wantErr {
				t.Errorf("Proxy.ContentManagementListProjects() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestProxy_ContentManagementCreate(t *testing.T) {
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
		projectLabel string
		name         string
		description  string
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

	resp := []byte(`{"success":true,"result":{"id":123,"lastBuildDate":"Aug 26, 2022, 04:06:32 PM","label":"new label","name":"test name","description":"Test Description","firstEnvironmnet": "first","orgId":0}}`)

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
		Body:       []byte(`{"success":"abc","result":1}`),
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
		Body:       []byte(`{"success":true,"result":true}`),
		StatusCode: 200,
		Cookies:    nil,
	}

	negativeResultUnMarshalJSONRespErr := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMock3,
		logger:            log,
	}
	suseManagerMock3.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errUnMarshalJSONResp, nil)

	arg := args{
		requestID:    "test id",
		auth:         auth,
		projectLabel: "Test label",
		description:  "blabla",
		name:         "Test label",
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Content Management Create",
			fields:  positiveKey,
			args:    arg,
			wantErr: false,
		},
		{
			name:    "Negative",
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
			name:    "Json Negative",
			fields:  negativeKeyUnMarshalJSONErr,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "Json Resp Negative",
			fields:  negativeResultUnMarshalJSONRespErr,
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
			_, err := p.ContentManagementCreate(tt.args.requestID, tt.args.auth, tt.args.projectLabel, tt.args.name, tt.args.description)
			if (err != nil) != tt.wantErr {
				t.Errorf("ContentManagementCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestProxy_ContentManagementListFilters(t *testing.T) {
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

	resp := []byte(`{"success":true,"result":[{"id":0,"label":"New Label","name":"Test Channel","description":"","lastBuildDate":"Aug 26, 2022, 04:06:32 PM","firstEnvironment":"","orgId":0}]}`)

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
	negativeResultStatusCode := fields{
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

	negativeResultUnMarshalJSONErr := fields{
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

	negativeResultUnMarshalJSONRespErr := fields{
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
			name:    "Positive",
			fields:  positiveKey,
			args:    arg,
			wantErr: false,
		},
		{
			name:    "Negative",
			fields:  negativeKey,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "Status Code Negative",
			fields:  negativeResultStatusCode,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "UnMarshal Json Negative",
			fields:  negativeResultUnMarshalJSONErr,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "UnMarshal Json Resp Negative",
			fields:  negativeResultUnMarshalJSONRespErr,
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
			_, err := p.ContentManagementListFilters(tt.args.requestID, tt.args.auth)
			if (err != nil) != tt.wantErr {
				t.Errorf("Proxy.ContentManagementListFilters() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func TestProxy_ContentManagementAttachSource(t *testing.T) {
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
		projectLabel string
		sourceType   string
		sourceLabel  string
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

	resp := []byte(`{"success":true,"result":{"id":123,"lastBuildDate":"Aug 26, 2022, 04:06:32 PM","label":"new label","name":"test name","description":"Test Description","firstEnvironmnet": "first","orgId":0}}`)

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
		Body:       []byte(`{"success":"abc","result":1}`),
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
		Body:       []byte(`{"success":true,"result":true}`),
		StatusCode: 200,
		Cookies:    nil,
	}

	negativeResultUnMarshalJSONRespErr := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMock3,
		logger:            log,
	}
	suseManagerMock3.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errUnMarshalJSONResp, nil)

	arg := args{
		requestID:    "test id",
		auth:         auth,
		projectLabel: "Test label",
		sourceType:   "blabla",
		sourceLabel:  "Test label",
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
			name:    "Negative",
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
			name:    "Json Negative",
			fields:  negativeKeyUnMarshalJSONErr,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "Json Resp Negative",
			fields:  negativeResultUnMarshalJSONRespErr,
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
			_, err := p.ContentManagementAttachSource(tt.args.requestID, tt.args.auth, tt.args.projectLabel, tt.args.sourceType, tt.args.sourceLabel)
			if (err != nil) != tt.wantErr {
				t.Errorf("ContentManagementAttachSource() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestProxy_ContentManagementCreateFilter(t *testing.T) {
	type fields struct {
		cfg               *SumanConfig
		contentTypeHeader map[string]string
		suse              ISuseManagerAPI
		logger            *zap.Logger
		retrycount        int
	}
	type args struct {
		requestID  string
		auth       AuthParams
		name       string
		rule       string
		entityType string
		criteria   sumamodels.FilterCriteria
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

	resp := []byte(`{"success":true,"result":{"entityType":"erratum","criteria":{"field":"issue_date","matcher":"greatereq","value":"2023-01-04T00:00:00Z"},"name":"release-230104","rule":"deny","id":7,"orgId":1}}`)

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
		Body:       []byte(`{"success":"abc","result":1}`),
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
		Body:       []byte(`{"success":true,"result":true}`),
		StatusCode: 200,
		Cookies:    nil,
	}

	negativeResultUnMarshalJSONRespErr := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMock3,
		logger:            log,
	}
	suseManagerMock3.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errUnMarshalJSONResp, nil)

	var criteria sumamodels.FilterCriteria
	criteria.Field = "Field"
	criteria.Value = "2023-01-04T00:00:00Z"
	criteria.Matcher = "matcher"

	arg := args{
		requestID:  "test id",
		auth:       auth,
		name:       "Test label",
		entityType: "blabla",
		criteria:   criteria,
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
			name:    "Negative",
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
			name:    "Json Negative",
			fields:  negativeKeyUnMarshalJSONErr,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "Json Resp Negative",
			fields:  negativeResultUnMarshalJSONRespErr,
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
			_, err := p.ContentManagementCreateFilter(tt.args.requestID, tt.args.auth, tt.args.name, tt.args.rule, tt.args.entityType, tt.args.criteria)
			if (err != nil) != tt.wantErr {
				t.Errorf("ContentManagementCreateFilter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestProxy_ContentManagementCreateEnvironment(t *testing.T) {
	type fields struct {
		cfg               *SumanConfig
		contentTypeHeader map[string]string
		suse              ISuseManagerAPI
		logger            *zap.Logger
		retrycount        int
	}
	type args struct {
		requestID        string
		auth             AuthParams
		projectLabel     string
		predecessorLabel string
		envlabel         string
		name             string
		description      string
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

	resp := []byte(`{"success":true,"result":{"entityType":"erratum","criteria":{"field":"issue_date","matcher":"greatereq","value":"2023-01-04T00:00:00Z"},"name":"release-230104","rule":"deny","id":7,"orgId":1}}`)

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
		Body:       []byte(`{"success":"abc","result":1}`),
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
		Body:       []byte(`{"success":true,"result":true}`),
		StatusCode: 200,
		Cookies:    nil,
	}

	negativeResultUnMarshalJSONRespErr := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMock3,
		logger:            log,
	}
	suseManagerMock3.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errUnMarshalJSONResp, nil)

	var criteria sumamodels.FilterCriteria
	criteria.Field = "Field"
	criteria.Value = "2023-01-04T00:00:00Z"
	criteria.Matcher = "matcher"

	arg := args{
		requestID:        "test id",
		auth:             auth,
		projectLabel:     "Test label",
		predecessorLabel: "f",
		envlabel:         "r001",
		name:             "r001",
		description:      "r001",
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
			name:    "Negative",
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
			name:    "Json Negative",
			fields:  negativeKeyUnMarshalJSONErr,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "Json Resp Negative",
			fields:  negativeResultUnMarshalJSONRespErr,
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
			_, err := p.ContentManagementCreateEnvironment(tt.args.requestID, tt.args.auth, tt.args.projectLabel, tt.args.predecessorLabel, tt.args.envlabel, tt.args.name, tt.args.description)
			if (err != nil) != tt.wantErr {
				t.Errorf("ContentManagementCreateEnvironment() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestProxy_ContentManagementBuildProject(t *testing.T) {
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
		projectLabel string
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

	resp := []byte(`{"success":true,"result": 1}`)

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
		Body:       []byte(`{"success":"abc","result":1}`),
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
		Body:       []byte(`{"success":true,"result":true}`),
		StatusCode: 200,
		Cookies:    nil,
	}

	negativeResultUnMarshalJSONRespErr := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMock3,
		logger:            log,
	}
	suseManagerMock3.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errUnMarshalJSONResp, nil)

	var criteria sumamodels.FilterCriteria
	criteria.Field = "Field"
	criteria.Value = "2023-01-04T00:00:00Z"
	criteria.Matcher = "matcher"

	arg := args{
		requestID:    "test id",
		auth:         auth,
		projectLabel: "Test label",
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
			name:    "Negative",
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
			name:    "Json Negative",
			fields:  negativeKeyUnMarshalJSONErr,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "Json Resp Negative",
			fields:  negativeResultUnMarshalJSONRespErr,
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
			_, err := p.ContentManagementBuildProject(tt.args.requestID, tt.args.auth, tt.args.projectLabel)
			if (err != nil) != tt.wantErr {
				t.Errorf("ContentManagementBuildProject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
