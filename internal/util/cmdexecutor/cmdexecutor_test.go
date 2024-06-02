package cmdexecutor

import (
	"os/exec"
	"testing"

	"ecp-golang-cm/pkg/util/logger"

	"go.uber.org/zap"
	zapCore "go.uber.org/zap/zapcore"
)

func TestNewCMDExecutor(t *testing.T) {
	type args struct {
		logger *zap.Logger
	}
	testLogger := logger.Default("debug")
	wantCMDExecutor := NewCMDExecutor(testLogger)
	tests := []struct {
		name string
		args args
		want ICMDExecutor
	}{
		{
			name: "New OS Use Case",
			args: args{
				logger: testLogger,
			},
			want: wantCMDExecutor,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewCMDExecutor(tt.args.logger)
			if got == nil {
				t.Errorf("NewOS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_os_ExecuteCommand(t *testing.T) {
	type fields struct {
		logger      *zap.Logger
		execCommand func(name string, arg ...string) *exec.Cmd
	}
	type args struct {
		name string
		args []string
		zf   []zapCore.Field
	}

	execCMDSuccess := func(name string, arg ...string) *exec.Cmd {
		return exec.Command("hostname")
	}

	execCMDErr := func(name string, arg ...string) *exec.Cmd {
		return exec.Command("commandNotPresent")
	}

	execCMDStdErr := func(name string, arg ...string) *exec.Cmd {
		return exec.Command("hostname -rr")
	}

	testLogger := logger.Default("debug")
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Execute Command",
			fields: fields{
				logger:      testLogger,
				execCommand: execCMDSuccess,
			},
			args: args{
				name: "test binary",
				args: []string{"Test", "Args"},
				zf:   []zapCore.Field{zap.Any("Test", "TestZap")},
			},
			//want:    []string{"Test Args", ""},
			wantErr: false,
		},
		{
			name: "Execute Command Negative",
			fields: fields{
				logger:      testLogger,
				execCommand: execCMDErr,
			},
			args: args{
				name: "test binary",
				args: []string{"Test", "Args"},
				zf:   []zapCore.Field{zap.Any("Test", "TestZap")},
			},
			wantErr: true,
		},
		{
			name: "Execute Command StdErr Negative",
			fields: fields{
				logger:      testLogger,
				execCommand: execCMDStdErr,
			},
			args: args{
				name: "test binary",
				args: []string{"Test", "Args"},
				zf:   []zapCore.Field{zap.Any("Test", "TestZap")},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os := &cmdUtils{
				logger:      tt.fields.logger,
				execCommand: tt.fields.execCommand,
			}
			_, err := os.ExecuteCommand(tt.args.name, tt.args.args, tt.args.zf...)
			if (err != nil) != tt.wantErr {
				t.Errorf("os.ExecuteCommand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_cmdUtils_CreateDirectory(t *testing.T) {
	type fields struct {
		logger      *zap.Logger
		execCommand func(name string, arg ...string) *exec.Cmd
	}
	type args struct {
		path string
		zf   []zapCore.Field
	}
	execCMDSuccess := func(name string, arg ...string) *exec.Cmd {
		return exec.Command("hostname")
	}

	execCMDErr := func(name string, arg ...string) *exec.Cmd {
		return exec.Command("commandNotPresent")
	}

	testLogger := logger.Default("debug")
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Create Dir",
			fields: fields{
				logger:      testLogger,
				execCommand: execCMDSuccess,
			},
			args: args{
				path: "/present",
				zf:   []zapCore.Field{zap.Any("Test", "TestZap")},
			},
			wantErr: false,
		},
		{
			name: "Execute Command Negative",
			fields: fields{
				logger:      testLogger,
				execCommand: execCMDErr,
			},
			args: args{
				path: "",
				zf:   []zapCore.Field{zap.Any("Test", "TestZap")},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmdUtils := &cmdUtils{
				logger:      tt.fields.logger,
				execCommand: tt.fields.execCommand,
			}
			if err := cmdUtils.CreateDirectory(tt.args.path, tt.args.zf...); (err != nil) != tt.wantErr {
				t.Errorf("CreateDirectory() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
