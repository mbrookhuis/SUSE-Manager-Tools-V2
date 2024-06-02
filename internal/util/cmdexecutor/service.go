package cmdexecutor

import zapCore "go.uber.org/zap/zapcore"

// ICMDExecutor function definition for all command line scripts
type ICMDExecutor interface {
	ExecuteCommand(binaryName string, args []string, zf ...zapCore.Field) ([]string, error)
	CreateDirectory(path string, zf ...zapCore.Field) error
}
