// Package logger - logging configuration
package logger

import (
	"fmt"
	"net/url"

	"strings"
	"time"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Config - configuration for log entries
type Config struct {
	Level              string `envconfig:"LOG_LEVEL"`
	TimestampFormat    string `envconfig:"LOG_TIMESTAMP_FORMAT"`
	StdoutEnabled      bool   `envconfig:"LOG_STDOUT_ENABLED"`
	FilePath           string `envconfig:"LOG_FILE"`
	MaxSize            int    `envconfig:"LOG_FILE_MAX_SIZE"`
	MaxAge             int    `envconfig:"LOG_FILE_MAX_AGE"`
	CompressionEnabled bool   `envconfig:"LOG_COMPRESSION_ENABLED"`
	StacktraceEnabled  bool   `envconfig:"LOG_STACKTRACE_ENABLED"`
	EnableFileLogs     bool   `envconfig:"LOG_ENABLE_FILE"`
}

// Default - logging configuration
//
// param: loglevel
// return: zap.logger
func Default(loglevel string) *zap.Logger {
	zapConfig := zap.NewProductionConfig()
	zapConfig.OutputPaths = []string{"stdout"}
	zapConfig.DisableStacktrace = true
	zapConfig.Sampling = nil
	zapConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	logger, _ := zapConfig.Build()
	logger.Sync()
	zapConfig.Level.UnmarshalText([]byte(loglevel))
	return logger
}

// lumberjackSink - logger data
type lumberjackSink struct {
	*lumberjack.Logger
}

// Sync implements zap.Sink
//
// return: error
func (lumberjackSink) Sync() error { return nil }

// NewConfig - create new logger config
//
// param: config
// return: zap.config
// return: error
func NewConfig(config *Config) (*zap.Config, error) {
	zapConfig := zap.NewProductionConfig()
	zapConfig.Sampling = nil
	zapConfig.DisableStacktrace = !config.StacktraceEnabled

	zapConfig.OutputPaths = make([]string, 0)

	if config.EnableFileLogs {
		zapConfig.OutputPaths = append(zapConfig.OutputPaths, fmt.Sprintf("lumberjack:%s", config.FilePath))
	}

	if config.StdoutEnabled {
		zapConfig.OutputPaths = append(zapConfig.OutputPaths, "stdout")
	}
	if config.TimestampFormat != "" {
		encoderCfg := zap.NewProductionEncoderConfig()
		switch config.TimestampFormat {
		case "RFC3339":
			encoderCfg.EncodeTime = zapcore.RFC3339TimeEncoder
		case "ISO8601":
			encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
		case "EpochMillis":
			encoderCfg.EncodeTime = EpochMillisTimeEncoder
		}
		encoderCfg.TimeKey = "log_time_stamp"
		encoderCfg.LevelKey = "log_level"
		encoderCfg.MessageKey = "log_message"
		encoderCfg.CallerKey = "function_name"
		encoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder
		zapConfig.EncoderConfig = encoderCfg
	}
	if config.Level != "" {
		if err := zapConfig.Level.UnmarshalText([]byte(strings.ToLower(config.Level))); err != nil {
			return nil, errors.Errorf("Unable to parse log level from config %s: %s", config.Level, err.Error())
		}
	}
	return &zapConfig, nil
}

// New - logging started
//
// param: config
// param: zapConfig
// return: zap.logger
// return: function
// return: error
func New(config *Config, zapConfig *zap.Config) (*zap.Logger, func(), error) {
	err := zap.RegisterSink("lumberjack", func(u *url.URL) (zap.Sink, error) {
		return lumberjackSink{
			Logger: &lumberjack.Logger{
				Filename: config.FilePath,
				// in megabytes
				MaxSize:  1,
				MaxAge:   1,
				Compress: config.CompressionEnabled,
			},
		}, nil
	})
	if err != nil {
		return nil, nil, err
	}
	logger, err := zapConfig.Build()
	if err != nil {
		return nil, nil, err
	}

	stop := func() {
		defer func() {
			if err := recover(); err != nil {
				logger.Error("Panic while syncing zap logger")
			}
		}()
		if logger != nil {
			err := logger.Sync()
			if err != nil {
				return
			}
		}
	}
	return logger, stop, nil
}

// EpochMillisTimeEncoder - epoch encoder
//
// param: t
// param: enc
func EpochMillisTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	nanos := t.UnixNano()
	millis := nanos / int64(time.Millisecond)
	enc.AppendInt64(millis)
}
