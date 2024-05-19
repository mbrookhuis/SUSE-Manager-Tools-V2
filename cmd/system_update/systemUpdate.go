package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

// Constants needed for logging
const (
	// logfile
	logFile = "/var/log/ecp-suma/ecp-suma.log"
	// set log level to panic
	PanicLevel = 0
	// set log level to fatal
	FatalLevel = 1
	// set log level to error
	ErrorLevel = 2
	// set log level to warn
	WarnLevel = 3
	// set log level to info
	InfoLevel = 4
	// set log level to debug
	DebugLevel = 5
	// set log level to trace
	TraceLevel = 6
)

// init
func init() {

	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(file)
}

// main
func main() {
	logLevel := "debug"
	switch logLevel {
	case "debug":
		log.SetLevel(DebugLevel)
	case "info":
		log.SetLevel(InfoLevel)
	case "error":
		log.SetLevel(ErrorLevel)
	case "warning":
		log.SetLevel(WarnLevel)
	case "trace":
		log.SetLevel(TraceLevel)
	case "panic":
		log.SetLevel(PanicLevel)
	case "fatal":
		log.SetLevel(FatalLevel)

	}
	// _ecpsuma_controller.Execute()
}
