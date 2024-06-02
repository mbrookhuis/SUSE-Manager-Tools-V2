package systemUpdate

import (
	"os"
	"path/filepath"

	_gc "SUSE-Manager-Tools-V2/internal/getConfig"
	log "SUSE-Manager-Tools-V2/internal/logger"
	_sumanUseCase "SUSE-Manager-Tools-V2/internal/susemanager"
	util "SUSE-Manager-Tools-V2/internal/util/uuid"
	"SUSE-Manager-Tools-V2/internal/vars"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

// const - needed constants
const (
	loglevel          = "debug"
	timestampFormat   = "ISO8601" // "RFC3339"
	stdoutEnabled     = true
	maxSize           = 100
	stacktraceEnabled = false
	logFileName       = "system-update.log"
	EnableFileLogs    = true
	retryCount        = 5
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main().
func Execute(params vars.Params) {
	loggerConfig := log.Config{
		Level:             loglevel,
		TimestampFormat:   timestampFormat,
		StdoutEnabled:     stdoutEnabled,
		FilePath:          "c:\\work\\test\\suse-tools\\log\\blabla",
		MaxSize:           maxSize,
		StacktraceEnabled: stacktraceEnabled,
		EnableFileLogs:    EnableFileLogs,
	}
	// Initialize logger object
	zapConfig, _ := log.NewConfig(&loggerConfig)
	logger, _, _ := log.New(&loggerConfig, zapConfig)
	logger.Info("system-update initiated")
	err := validateParams(params, logger)
	if err != nil {
		os.Exit(1)
	}
	configFileByte, err := _gc.ReadYamlConfigFile(params.ConfigFile)
	if err != nil {
		logger.Error("Unable to read configfile", zap.Any("Error", err))
		os.Exit(1)
	}
	var configFile vars.ConfigInfo
	err = yaml.Unmarshal(configFileByte, &configFile)
	if err != nil {
		logger.Error("Unable to convert configfile", zap.Any("Error", err))
		os.Exit(1)
	}

	loggerConfig = log.Config{
		Level:             loglevel,
		TimestampFormat:   timestampFormat,
		StdoutEnabled:     stdoutEnabled,
		FilePath:          filepath.Join(configFile.Dirs.LogDir, params.Server),
		MaxSize:           maxSize,
		StacktraceEnabled: stacktraceEnabled,
		EnableFileLogs:    EnableFileLogs,
	}
	// Initialize logger object
	zapConfig, _ = log.NewConfig(&loggerConfig)
	logger.Info("system-update starting")

	sumancfg := _sumanUseCase.SumanConfig{
		Host:     configFile.Suman.Server,
		Login:    configFile.Suman.User,
		Password: configFile.Suman.Password,
		Insecure: true,
	}
	requestID := util.GenerateUniqueID()
	suseAPI := _sumanUseCase.NewSuseManagerAPI("rhn/manager/api", true, logger, retryCount)
	sumanProxyUseCase := _sumanUseCase.NewProxy(&sumancfg, suseAPI, logger, retryCount)
	suseUseCase := _sumanUseCase.NewSuseManager(sumanProxyUseCase, &sumancfg, logger)
	systemUpdate := NewSystemUpdate(sumanProxyUseCase, suseUseCase, 120, logger, params, configFile, requestID)
	err = systemUpdate.SystemUpdate()
	if err != nil {
		logger.Error("systemUpdate failed", zap.Any("Server", params.Server), zap.Any("error", err.Error()))
		os.Exit(1)
	}
	logger.Info("Finished systemUpdate", zap.Any("Server", params.Server))

	/*
		fmt.Println(data.Maintenance.WaitBetweenSystems)
		fmt.Println(data.Maintenance.SpMigrationProject)

		for key, value := range data.Maintenance.SpMigrationProject {
			// Step 3: Print the key and value
			fmt.Printf("Key: %s, Value: %s\n", key, value)
		}
	*/

}

// init
func init() {
	// Add flags in root command if required
}
