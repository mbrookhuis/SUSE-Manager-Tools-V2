package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type SumanInfo struct {
	Server   string `mapstructure:"server"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Timeout  int    `mapstructure:"timeout"`
}

type SMTPInfo struct {
	Sendmail  bool     `mapstructure:"sendmail"`
	Receivers []string `mapstructure:"receivers"`
	Sender    string   `mapstructure:"sender"`
	Server    string   `mapstructure:"server"`
}

type DirsInfo struct {
	LogDir          string `mapstructure:"log_dir"`
	ScriptsDir      string `mapstructure:"scripts_dir"`
	UpdateScriptDir string `mapstructure:"update_script_dir"`
}

type LoglevelInfo struct {
	File   string `mapstructure:"file"`
	Screen string `mapstructure:"screen"`
}

type ErrorHandlingInfo struct {
	Script        string `mapstructure:"script"`
	Update        string `mapstructure:"update"`
	Spmig         string `mapstructure:"spmig"`
	Configupdate  string `mapstructure:"configupdate"`
	Reboot        string `mapstructure:"reboot"`
	TimeoutPassed string `mapstructure:"timeout_passed"`
}

type SpMigrationProjectInfo struct {
	SpMigrationProjectTarget map[string]interface{} `mapstructure:",inline"`
}

type SpMigrationInfo struct {
	SpMigrationTarget string `mapstructure:"link,omitempty"`
}

type ExceptionSpInfo struct {
	ExceptionSpTarget []string `mapstructure:"link,omitempty"`
}

type MaintenanceInfo struct {
	WaitBetweenSystems int                    `mapstructure:"wait_between_systems"`
	ExcludeForPatch    []string               `mapstructure:"exclude_for_patch"`
	SpMigrationProject map[string]interface{} `mapstructure:"sp_migration_project"`
	SpMigration        map[string]interface{} `mapstructure:"sp_migration"`
	ExceptionSp        map[string]interface{} `mapstructure:"exception_sp"`
}

type Config struct {
	Suman         SumanInfo         `mapstructure:"suman"`
	SMTP          SMTPInfo          `mapstructure:"smtp"`
	Dirs          DirsInfo          `mapstructure:"dirs"`
	Loglevel      LoglevelInfo      `mapstructure:"loglevel"`
	ErrorHandling ErrorHandlingInfo `mapstructure:"error_handling"`
	Maintenance   MaintenanceInfo   `mapstructure:"maintenance"`
}

// UpdateScript
type UpdateScript struct {
	BeginScript BeginScript `mapstructure:"begin,omitempty"`
	EndScript   EndScript   `mapstructure:"end,omitempty"`
}

type BeginScript struct {
	TimeOut  int      `mapstructure:"timeout"`
	Commands []string `mapstructure:"commands,omitempty"`
	State    []string `mapstructure:"state,omitempty"`
}

type EndScript struct {
	TimeOut  int      `mapstructure:"timeout"`
	Commands []string `mapstructure:"commands,omitempty"`
	State    []string `mapstructure:"state,omitempty"`
}

var config Config

// LoadFrom - loads configuration from specific path
func LoadFrom(path string, fileName string) {
	viper.SetConfigName(fileName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Config file not found or error reading the file: %e", err)
		os.Exit(1)
	}
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("unable to unmarshall configuration: %e", err)
		os.Exit(1)
	}
	// LoadViper()

}

// GetConfig - Gets loaded configuration
func GetConfig() *Config {
	return &config
}

// LoadAllFrom - loads configuration from local file
func LoadAllFrom(path string, fileName string) {
	LoadFrom(path, fileName)
}
