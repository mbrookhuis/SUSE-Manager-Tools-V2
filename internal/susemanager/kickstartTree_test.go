// Package SUSE Manager - SUSE Manager api call and support functions
package susemanager

import (
	"encoding/json"
	"testing"

	sumamodels "ecp-golang-cm/pkg/models/susemanager"
	"ecp-golang-cm/pkg/util/logger"
	util "ecp-golang-cm/pkg/util/rest"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/mock"
	"go.uber.org/zap"
)

func TestProxy_KickstartTreeGetDetails(t *testing.T) {
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
		distributionName string
	}
	log := logger.Default("info")
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
		Body:       []byte(`{"success":"abc","result":{"id":123,"name":"Test Name"}}`),
		StatusCode: 200,
		Cookies:    nil,
	}

	negativeUnMarshalJSONErr := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMock2,
		logger:            log,
	}
	suseManagerMock2.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errSuccesssHelper, nil)

	suseManagerMock3 := new(MockISuseManagerAPI)
	errUnMarshalJSONResp := &util.HTTPHelperStruct{
		Body:       []byte(`{"success":true,"result":{"id":false,"name":"Test Name"}}`),
		StatusCode: 200,
		Cookies:    nil,
	}

	negativeUnMarshalJSONRespErr := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMock3,
		logger:            log,
	}
	suseManagerMock3.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errUnMarshalJSONResp, nil)

	auth := AuthParams{
		SessionKey: "test key",
		Host:       "test Hostname",
	}

	arg := args{
		requestID:        "test id",
		auth:             auth,
		distributionName: "test group",
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
			name:    "UnMarshal Json Negative",
			fields:  negativeUnMarshalJSONErr,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "UnMarshal Json Resp Negative",
			fields:  negativeUnMarshalJSONRespErr,
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
			_, err := p.KickstartTreeGetDetails(tt.args.requestID, tt.args.auth, tt.args.distributionName)
			if (err != nil) != tt.wantErr {
				t.Errorf("KickstartTreeGetDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestProxy_KickstartTreeCreate(t *testing.T) {
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
		treeLabel    string
		basePath     string
		channelLabel string
		installType  string
	}

	log := logger.Default("info")
	header := make(map[string]string)
	header["Content-Type"] = "application/json"

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
		Body:       []byte(`{"success":"abc","result": 1}`),
		StatusCode: 200,
		Cookies:    nil,
	}

	negativeUnMarshalJSONErr := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMock2,
		logger:            log,
	}
	suseManagerMock2.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errSuccesssHelper, nil)

	suseManagerMock3 := new(MockISuseManagerAPI)
	errUnMarshalJSONResp := &util.HTTPHelperStruct{
		Body:       []byte(`{"success":true,"result": true}`),
		StatusCode: 200,
		Cookies:    nil,
	}

	negativeUnMarshalJSONRespErr := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMock3,
		logger:            log,
	}
	suseManagerMock3.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errUnMarshalJSONResp, nil)

	suseManagerMock4 := new(MockISuseManagerAPI)
	errMarshalJSONResp := &util.HTTPHelperStruct{
		Body:       []byte(`{"success":"abc","result": 1}`),
		StatusCode: 200,
		Cookies:    nil,
	}

	negativeMarshalJSONErr := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMock4,
		logger:            log,
	}
	suseManagerMock4.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errMarshalJSONResp, nil)

	auth := AuthParams{
		SessionKey: "test key",
		Host:       "test Hostname",
	}

	arg := args{
		requestID:    "test id",
		auth:         auth,
		treeLabel:    "label",
		basePath:     "/path/to/kernel",
		channelLabel: "channel Label",
		installType:  "sles15",
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
			name:    "UnMarshal Json Negative",
			fields:  negativeUnMarshalJSONErr,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "Marshal Json Resp Negative",
			fields:  negativeMarshalJSONErr,
			args:    arg,
			wantErr: true,
		}, {
			name:    "UnMarshal Json Resp Negative",
			fields:  negativeUnMarshalJSONRespErr,
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
			_, err := p.KickstartTreeCreate(tt.args.requestID, tt.args.auth, tt.args.treeLabel, tt.args.basePath, tt.args.channelLabel, tt.args.installType)
			if (err != nil) != tt.wantErr {
				t.Errorf("KickstartTreeCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestProxy_KickstartTreeCreateKernelOptions(t *testing.T) {
	type fields struct {
		cfg               *SumanConfig
		contentTypeHeader map[string]string
		suse              ISuseManagerAPI
		logger            *zap.Logger
		retrycount        int
	}
	type args struct {
		requestID         string
		auth              AuthParams
		treeLabel         string
		basePath          string
		channelLabel      string
		installType       string
		kernelOptions     string
		postKernelOptions string
	}

	log := logger.Default("info")
	header := make(map[string]string)
	header["Content-Type"] = "application/json"

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
		Body:       []byte(`{"success":"abc","result": 1}`),
		StatusCode: 200,
		Cookies:    nil,
	}

	negativeUnMarshalJSONErr := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMock2,
		logger:            log,
	}
	suseManagerMock2.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errSuccesssHelper, nil)

	suseManagerMock3 := new(MockISuseManagerAPI)
	errUnMarshalJSONResp := &util.HTTPHelperStruct{
		Body:       []byte(`{"success":true,"result": true}`),
		StatusCode: 200,
		Cookies:    nil,
	}

	negativeUnMarshalJSONRespErr := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMock3,
		logger:            log,
	}
	suseManagerMock3.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errUnMarshalJSONResp, nil)

	suseManagerMock4 := new(MockISuseManagerAPI)
	errMarshalJSONResp := &util.HTTPHelperStruct{
		Body:       []byte(`{"success":"abc","result": 1}`),
		StatusCode: 200,
		Cookies:    nil,
	}

	negativeMarshalJSONErr := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMock4,
		logger:            log,
	}
	suseManagerMock4.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errMarshalJSONResp, nil)

	auth := AuthParams{
		SessionKey: "test key",
		Host:       "test Hostname",
	}

	arg := args{
		requestID:         "test id",
		auth:              auth,
		treeLabel:         "label",
		basePath:          "/path/to/kernel",
		channelLabel:      "channel Label",
		installType:       "sles15",
		kernelOptions:     "useonlinerepo=1 insecure=1",
		postKernelOptions: "",
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
			name:    "UnMarshal Json Negative",
			fields:  negativeUnMarshalJSONErr,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "Marshal Json Resp Negative",
			fields:  negativeMarshalJSONErr,
			args:    arg,
			wantErr: true,
		}, {
			name:    "UnMarshal Json Resp Negative",
			fields:  negativeUnMarshalJSONRespErr,
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
			_, err := p.KickstartTreeCreateKernelOptions(tt.args.requestID, tt.args.auth, tt.args.treeLabel, tt.args.basePath, tt.args.channelLabel, tt.args.installType, tt.args.kernelOptions, tt.args.postKernelOptions)
			if (err != nil) != tt.wantErr {
				t.Errorf("KickstartTreeCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestProxy_KickstartImportRawFile(t *testing.T) {
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
		profileLabel string
		virtType     string
		channelLabel string
		dataXML      string
	}

	log := logger.Default("info")
	header := make(map[string]string)
	header["Content-Type"] = "application/json"

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
		Body:       []byte(`{"success":"abc","result": 1}`),
		StatusCode: 200,
		Cookies:    nil,
	}

	negativeUnMarshalJSONErr := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMock2,
		logger:            log,
	}
	suseManagerMock2.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errSuccesssHelper, nil)

	suseManagerMock3 := new(MockISuseManagerAPI)
	errUnMarshalJSONResp := &util.HTTPHelperStruct{
		Body:       []byte(`{"success":true,"result": true}`),
		StatusCode: 200,
		Cookies:    nil,
	}

	negativeUnMarshalJSONRespErr := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMock3,
		logger:            log,
	}
	suseManagerMock3.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errUnMarshalJSONResp, nil)

	suseManagerMock4 := new(MockISuseManagerAPI)
	errMarshalJSONResp := &util.HTTPHelperStruct{
		Body:       []byte(`{"success":"abc","result": 1}`),
		StatusCode: 200,
		Cookies:    nil,
	}

	negativeMarshalJSONErr := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMock4,
		logger:            log,
	}
	suseManagerMock4.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errMarshalJSONResp, nil)

	auth := AuthParams{
		SessionKey: "test key",
		Host:       "test Hostname",
	}

	arg := args{
		requestID:    "test id",
		auth:         auth,
		profileLabel: "label",
		virtType:     "none",
		channelLabel: "channel Label",
		dataXML:      "<xml>data</xml>\n<line>1</line>\n",
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
			name:    "UnMarshal Json Negative",
			fields:  negativeUnMarshalJSONErr,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "Marshal Json Resp Negative",
			fields:  negativeMarshalJSONErr,
			args:    arg,
			wantErr: true,
		}, {
			name:    "UnMarshal Json Resp Negative",
			fields:  negativeUnMarshalJSONRespErr,
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
			_, err := p.KickstartImportRawFile(tt.args.requestID, tt.args.auth, tt.args.profileLabel, tt.args.virtType, tt.args.channelLabel, tt.args.dataXML)
			if (err != nil) != tt.wantErr {
				t.Errorf("KickstartTreeCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestProxy_KickstartListKickstarts(t *testing.T) {
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
	}

	log := logger.Default("info")
	header := make(map[string]string)
	header["Content-Type"] = "application/json"

	config := &SumanConfig{
		Host:     "test host",
		Password: "test",
		Insecure: true,
		Login:    "test",
	}

	suseManagerMock := new(MockISuseManagerAPI)

	kickstartResp := []sumamodels.KickstartListProfiles{
		{
			Name:  "present",
			Label: "New Label",
		},
	}

	success := sumamodels.RespAPISuccess{
		Success: true,
		Result:  kickstartResp,
	}

	resp, err := json.Marshal(success)
	if err != nil {
		panic(err)
	}
	//kickstartRespError := []sumamodels.KickstartListProfiles{
	//	{},
	//}

	failure := sumamodels.RespAPISuccess{
		Success: true,
		Result:  nil,
	}

	respErr, err := json.Marshal(failure)
	if err != nil {
		panic(err)
	}
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
		Body:       respErr,
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
		Body:       respErr,
		StatusCode: 200,
		Cookies:    nil,
	}

	suseManagerMock2.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errSuccesssHelper, nil)

	suseManagerMock3 := new(MockISuseManagerAPI)
	errUnMarshalJSONResp := &util.HTTPHelperStruct{
		Body:       respErr,
		StatusCode: 200,
		Cookies:    nil,
	}

	suseManagerMock3.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errUnMarshalJSONResp, nil)

	suseManagerMock4 := new(MockISuseManagerAPI)
	errMarshalJSONResp := &util.HTTPHelperStruct{
		Body:       respErr,
		StatusCode: 200,
		Cookies:    nil,
	}

	suseManagerMock4.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errMarshalJSONResp, nil)

	auth := AuthParams{
		SessionKey: "test key",
		Host:       "test Hostname",
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
			_, err := p.KickstartListKickstarts(tt.args.requestID, tt.args.auth)
			if (err != nil) != tt.wantErr {
				t.Errorf("KickstartTreeCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestProxy_KickstartDeleteProfile(t *testing.T) {
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
		profileLabel string
	}

	log := logger.Default("info")
	header := make(map[string]string)
	header["Content-Type"] = "application/json"

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
		Body:       []byte(`{"success":"abc","result": 1}`),
		StatusCode: 200,
		Cookies:    nil,
	}

	negativeUnMarshalJSONErr := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMock2,
		logger:            log,
	}
	suseManagerMock2.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errSuccesssHelper, nil)

	suseManagerMock3 := new(MockISuseManagerAPI)
	errUnMarshalJSONResp := &util.HTTPHelperStruct{
		Body:       []byte(`{"success":true,"result": true}`),
		StatusCode: 200,
		Cookies:    nil,
	}

	negativeUnMarshalJSONRespErr := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMock3,
		logger:            log,
	}
	suseManagerMock3.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errUnMarshalJSONResp, nil)

	suseManagerMock4 := new(MockISuseManagerAPI)
	errMarshalJSONResp := &util.HTTPHelperStruct{
		Body:       []byte(`{"success":"abc","result": 1}`),
		StatusCode: 200,
		Cookies:    nil,
	}

	negativeMarshalJSONErr := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMock4,
		logger:            log,
	}
	suseManagerMock4.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errMarshalJSONResp, nil)

	auth := AuthParams{
		SessionKey: "test key",
		Host:       "test Hostname",
	}

	arg := args{
		requestID:    "test id",
		auth:         auth,
		profileLabel: "label",
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
			name:    "UnMarshal Json Negative",
			fields:  negativeUnMarshalJSONErr,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "Marshal Json Resp Negative",
			fields:  negativeMarshalJSONErr,
			args:    arg,
			wantErr: true,
		}, {
			name:    "UnMarshal Json Resp Negative",
			fields:  negativeUnMarshalJSONRespErr,
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
			_, err := p.KickstartDeleteProfile(tt.args.requestID, tt.args.auth, tt.args.profileLabel)
			if (err != nil) != tt.wantErr {
				t.Errorf("KickstartTreeCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestProxy_KickstartProfileSetVariables(t *testing.T) {
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
		profileLabel     string
		profileVariables map[string]interface{}
	}

	log := logger.Default("info")
	header := make(map[string]string)
	header["Content-Type"] = "application/json"

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
		Body:       []byte(`{"success":"abc","result": 1}`),
		StatusCode: 200,
		Cookies:    nil,
	}

	negativeUnMarshalJSONErr := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMock2,
		logger:            log,
	}
	suseManagerMock2.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errSuccesssHelper, nil)

	suseManagerMock3 := new(MockISuseManagerAPI)
	errUnMarshalJSONResp := &util.HTTPHelperStruct{
		Body:       []byte(`{"success":true,"result": true}`),
		StatusCode: 200,
		Cookies:    nil,
	}

	negativeUnMarshalJSONRespErr := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMock3,
		logger:            log,
	}
	suseManagerMock3.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errUnMarshalJSONResp, nil)

	suseManagerMock4 := new(MockISuseManagerAPI)
	errMarshalJSONResp := &util.HTTPHelperStruct{
		Body:       []byte(`{"success":"abc","result": 1}`),
		StatusCode: 200,
		Cookies:    nil,
	}

	negativeMarshalJSONErr := fields{
		cfg:               config,
		contentTypeHeader: header,
		suse:              suseManagerMock4,
		logger:            log,
	}
	suseManagerMock4.On("SuseManagerCall", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errMarshalJSONResp, nil)

	auth := AuthParams{
		SessionKey: "test key",
		Host:       "test Hostname",
	}

	arg := args{
		requestID:        "test id",
		auth:             auth,
		profileLabel:     "label",
		profileVariables: nil,
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
			name:    "UnMarshal Json Negative",
			fields:  negativeUnMarshalJSONErr,
			args:    arg,
			wantErr: true,
		},
		{
			name:    "Marshal Json Resp Negative",
			fields:  negativeMarshalJSONErr,
			args:    arg,
			wantErr: true,
		}, {
			name:    "UnMarshal Json Resp Negative",
			fields:  negativeUnMarshalJSONRespErr,
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
			_, err := p.KickstartProfileSetVariables(tt.args.requestID, tt.args.auth, tt.args.profileLabel, tt.args.profileVariables)
			if (err != nil) != tt.wantErr {
				t.Errorf("KickstartProfileSetVariables() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
