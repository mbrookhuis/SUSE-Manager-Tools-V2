package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"SUSE-Manager-Tools-V2/internal/config"
	"SUSE-Manager-Tools-V2/internal/systemUpdate"
	"SUSE-Manager-Tools-V2/internal/vars"
	log "github.com/sirupsen/logrus"
)

// Constants needed for logging
const (
	// logfile
	logFileName = "system-update.log"
	// PanicLevel set log level to panic
	PanicLevel = 0
	// FatalLevel set log level to fatal
	FatalLevel = 1
	// ErrorLevel set log level to error
	ErrorLevel = 2
	// WarnLevel set log level to warn
	WarnLevel = 3
	// InfoLevel set log level to info
	InfoLevel = 4
	// DebugLevel set log level to debug
	DebugLevel = 5
	// TraceLevel set log level to trace
	TraceLevel = 6
	// ConfigFile configuration file containing the needed configuration
	ConfigFile = "configsm"
)

// init
func init() {
	var logFile string
	switch runtime.GOOS {
	case "windows":
		path, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		logFile = filepath.Join(path, logFileName)
	case "linux":
		logFile = filepath.Join("/var/log/system-update", logFileName)
	case "default":
		logFile = logFileName
	}
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(file)
}

// main
func main() {
	var params vars.Params
	log.Info("system-update initiated")
	flag.StringVar(&params.Server, "s", "not_set", "name of the server to receive config update. Required")
	flag.BoolVar(&params.NoReboot, "n", false, "Do not reboot server after patching or supportpack upgrade. Default: reboot when needed")
	flag.BoolVar(&params.NoDryRun, "d", false, "Do not run a dry run before performing a SP migration. Default: perform dryrun")
	flag.BoolVar(&params.ForceReboot, "f", false, "Force a reboot server after patching or supportpack upgrade. Default: reboot when needed")
	flag.BoolVar(&params.ApplyConfig, "c", false, "Apply configuration after and before patching. Default: no config update")
	flag.BoolVar(&params.UpdateScript, "u", false, "Execute the server specific _begin and _end scripts. Default: no scripts are being executed")
	flag.StringVar(&params.PostScript, "p", "no_script_given", "Execute given script on the SUSE Manger Server when system_update has finished")
	flag.StringVar(&params.ConfigFile, "i", "not_set", "Configuration file containing the needed configuration. Default: configsm.yaml in the directory where updateServer has been started.")
	flag.Parse()
	var path, fileName string
	if params.ConfigFile == "not_set" {
		path = "."
		fileName = ConfigFile
	} else {
		configFile := strings.Split(params.ConfigFile, "/")
		fileName = configFile[len(configFile)-1]
		path = strings.Replace(params.ConfigFile, fmt.Sprintf("/%s", fileName), "", 1)
		fileName = strings.Split(fileName, ".")[0]
	}
	config.LoadAllFrom(path, fileName)
	log.SetLevel(DebugLevel)
	systemUpdate.Execute(params)

}
