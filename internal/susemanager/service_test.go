// Package SUSE Manager - SUSE Manager api call and support functions
package susemanager

import (
	"reflect"
	"testing"

	"ecp-golang-cm/pkg/util/logger"

	"go.uber.org/zap"
)

func TestNewProxy(t *testing.T) {
	type args struct {
		s      *SumanConfig
		suse   ISuseManagerAPI
		logger *zap.Logger
	}

	config := &SumanConfig{
		Host:     "test host",
		Password: "test",
		Insecure: true,
		Login:    "test",
	}
	testLogger := logger.Default("debug")

	suseManagerMock := new(MockISuseManagerAPI)
	iproxy := NewProxy(config, suseManagerMock, testLogger, 5)

	tests := []struct {
		name string
		args args
		want IProxy
	}{
		{
			name: "New Proxy",
			args: args{
				s:      config,
				suse:   suseManagerMock,
				logger: testLogger,
			},
			want: iproxy,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewProxy(tt.args.s, tt.args.suse, tt.args.logger, 5); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProxy() = %v, want %v", got, tt.want)
			}
		})
	}
}
