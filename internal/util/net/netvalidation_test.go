package netvalidation

import (
	"fmt"
	"net"
	"net/url"
	"reflect"
	"testing"
)

func TestValidateIP(t *testing.T) {
	type args struct {
		host string
	}

	arg := args{host: "172.27.27.27"}
	arg1 := args{host: "123"}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "testcase1",
			args: arg,
			want: false,
		},
		{
			name: "testcase2",
			args: arg1,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ValidateIP(tt.args.host); got != tt.want {
				t.Errorf("ValidateIP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateUrl(t *testing.T) {
	type args struct {
		host string
	}
	arg := args{host: "http://www.google.com"}
	expectedresp := "http://www.google.com"

	arg1 := args{host: "xyz"}
	var u *url.URL
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "positivetestcase",
			args:    arg,
			want:    expectedresp,
			wantErr: false,
		},
		{
			name:    "negativetestcase",
			args:    arg1,
			want:    fmt.Sprint(u),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ValidateURL(tt.args.host)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ValidateURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckDNSLookup(t *testing.T) {
	type args struct {
		host string
	}

	arg := args{
		host: "xyz",
	}
	var res []net.IP
	tests := []struct {
		name    string
		args    args
		want    []net.IP
		wantErr bool
	}{
		{
			name:    "negativetestcase",
			args:    arg,
			want:    res,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CheckDNSLookup(tt.args.host)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckDNSLookup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckDNSLookup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseUrl(t *testing.T) {
	type args struct {
		host string
	}

	arg := args{host: "http://test.com"}
	resp := url.URL{Scheme: "http", Host: "test.com", ForceQuery: false}
	tests := []struct {
		name    string
		args    args
		want    *url.URL
		wantErr bool
	}{
		{
			name:    "positivetestcase",
			args:    arg,
			want:    &resp,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseURL(tt.args.host)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
