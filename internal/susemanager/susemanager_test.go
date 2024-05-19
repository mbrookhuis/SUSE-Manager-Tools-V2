// Package SUSE Manager - SUSE Manager api call and support functions
package susemanager

import (
	"encoding/json"
	"reflect"
	"testing"

	sumamodels "ecp-golang-cm/pkg/models/susemanager"
)

func TestHandleSuseManagerResponse(t *testing.T) {
	type args struct {
		body []byte
	}

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
	arg := args{
		body: activationByteArr,
	}
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

	arg2 := args{
		body: failureByteArr,
	}

	failureSuccessFail := sumamodels.RespAPI{
		Success:  false,
		Result:   nil,
		Messages: MessBody,
	}

	failureByteArrFail, err := json.Marshal(failureSuccessFail)
	if err != nil {
		panic(err)
	}

	arg3 := args{
		body: failureByteArrFail,
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Handle Suse Manager Resp Positive",
			args:    arg,
			wantErr: false,
		},
		{
			name:    "Handle Suse Manager Failure",
			args:    arg2,
			wantErr: true,
		},
		{
			name:    "Handle Suse Manager Fail Resp",
			args:    arg3,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := HandleSuseManagerResponse(tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("HandleSuseManagerResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestSuseManager_GetSystemGroupName(t *testing.T) {
	type fields struct {
		cfg *SumanConfig
	}
	type args struct {
		negName string
	}

	config := &SumanConfig{
		Host:     "test host",
		Password: "test",
		Insecure: true,
		Login:    "test",
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "Get System Group Name",
			fields: fields{
				cfg: config,
			},
			args: args{
				negName: "testNegName",
			},
			want: "a4-loc-testnegname",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SuseManager{
				cfg: tt.fields.cfg,
			}
			if got := s.GetSystemGroupName(tt.args.negName); got != tt.want {
				t.Errorf("SuseManager.GetSystemGroupName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSuseManager_GetAuth(t *testing.T) {
	type fields struct {
		cfg *SumanConfig
	}
	type args struct {
		sessionKey string
	}

	config := &SumanConfig{
		Host:     "test Hostname",
		Password: "test",
		Insecure: true,
		Login:    "test",
	}

	auth := AuthParams{
		SessionKey: "test key",
		Host:       "test Hostname",
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *AuthParams
		wantErr bool
	}{
		{
			name: "Get System Group Name",
			fields: fields{
				cfg: config,
			},
			args: args{
				sessionKey: "test key",
			},
			want:    &auth,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SuseManager{
				cfg: tt.fields.cfg,
			}
			got, err := s.GetAuth(tt.args.sessionKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("SuseManager.GetAuth() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SuseManager.GetAuth() = %v, want %v", got, tt.want)
			}
		})
	}
}
