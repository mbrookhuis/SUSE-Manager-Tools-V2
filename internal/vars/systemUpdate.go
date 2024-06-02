package vars

type Params struct {
	Server       string
	NoReboot     bool
	NoDryRun     bool
	ForceReboot  bool
	ApplyConfig  bool
	UpdateScript bool
	PostScript   string
	ConfigFile   string
}

// Config file structure

type SumanInfo struct {
	Server   string `yaml:"server"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Timeout  int    `yaml:"timeout"`
}

type SMTPInfo struct {
	Sendmail  bool     `yaml:"sendmail"`
	Receivers []string `yaml:"receivers"`
	Sender    string   `yaml:"sender"`
	Server    string   `yaml:"server"`
}

type DirsInfo struct {
	LogDir          string `yaml:"log_dir"`
	ScriptsDir      string `yaml:"scripts_dir"`
	UpdateScriptDir string `yaml:"update_script_dir"`
}

type LoglevelInfo struct {
	File   string `yaml:"file"`
	Screen string `yaml:"screen"`
}

type ErrorHandlingInfo struct {
	Script        string `yaml:"script"`
	Update        string `yaml:"update"`
	Spmig         string `yaml:"spmig"`
	Configupdate  string `yaml:"configupdate"`
	Reboot        string `yaml:"reboot"`
	TimeoutPassed string `yaml:"timeout_passed"`
}

type SpMigrationProjectInfo struct {
	SpMigrationProjectTarget map[string]interface{} `yaml:",inline"`
}

type SpMigrationInfo struct {
	SpMigrationTarget string `yaml:"link,omitempty"`
}

type ExceptionSpInfo struct {
	ExceptionSpTarget []string `yaml:"link,omitempty"`
}

type MaintenanceInfo struct {
	WaitBetweenSystems int                    `yaml:"wait_between_systems"`
	ExcludeForPatch    []string               `yaml:"exclude_for_patch"`
	SpMigrationProject map[string]interface{} `yaml:"sp_migration_project"`
	SpMigration        map[string]interface{} `yaml:"sp_migration"`
	ExceptionSp        map[string]interface{} `yaml:"exception_sp"`
}

type ConfigInfo struct {
	Suman         SumanInfo         `yaml:"suman"`
	SMTP          SMTPInfo          `yaml:"smtp"`
	Dirs          DirsInfo          `yaml:"dirs"`
	Loglevel      LoglevelInfo      `yaml:"loglevel"`
	ErrorHandling ErrorHandlingInfo `yaml:"error_handling"`
	Maintenance   MaintenanceInfo   `yaml:"maintenance"`
}

// UpdateScript
type UpdateScript struct {
	BeginScript BeginScript `yaml:"begin,omitempty"`
	EndScript   EndScript   `yaml:"end,omitempty"`
}

type BeginScript struct {
	TimeOut  int      `yaml:"timeout"`
	Commands []string `yaml:"commands,omitempty"`
	State    []string `yaml:"state,omitempty"`
}

type EndScript struct {
	TimeOut  int      `yaml:"timeout"`
	Commands []string `yaml:"commands,omitempty"`
	State    []string `yaml:"state,omitempty"`
}
