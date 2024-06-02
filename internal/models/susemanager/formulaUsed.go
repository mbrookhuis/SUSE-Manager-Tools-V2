// Package sumamodels - structs needed for SUSE Manager API Calls
package sumamodels

/*
PODConfig formular
*/

// DtagPodConfig podconfig formular
type DtagPodConfig struct {
	PodCfg PodConfig `json:"podconfig"`
}

// PodConfig - api call info
type PodConfig struct {
	PodType            string `json:"pod_type"`
	PodHostname        string `json:"pod_hostname"`
	PodNegID           string `json:"pod_negId"`
	PodNeID            string `json:"pod_neId"`
	PodNr              string `json:"pod_nr"`
	PodPaoIP           string `json:"pod_pao_ip"`
	PodServer          string `json:"pod_ser"`
	PodNtp             string `json:"pod_ntp"`
	PodIPAddr          string `json:"pod_ip_address_mgt"`
	OsInstalledVersion string `json:"os_installed_version"`
}

/*
Uyunihub Yaml
*/

// Uyunihub uyunihub formular
type Uyunihub struct {
	Hub Hub `json:"hub"`
}

// Hub - api call info
type Hub struct {
	ServerUserName string `json:"server_username"`
	ServerPassword string `json:"server_password"`
	HubOrg         string `json:"hub_org"`
}

/*
MgtsServ formular
*/

// MgtsSrvFormular mgts formular
type MgtsSrvFormular struct {
	Mgntenv Mgntenv `json:"mgntenv"`
}

// Mgntenv - api call info
type Mgntenv struct {
	MgtsHostname string         `json:"mgts_hostname"`
	MgtsNTP      string         `json:"mgts_ntp"`
	MgtsEnv      string         `json:"dtag_env"`
	MgtsRoutes   []MgtsRoutes   `json:"mgts_routes"`
	ArtLinks     []DefArtLinks  `json:"art_links,omitempty"`
	MgtsDNS      string         `json:"mgts_dns"`
	MgtsEth      []MgtsEth      `json:"mgts_eth"`
	MgtsRtTables []MgtsRtTables `json:"mgts_rt_tables"`
	MgtsType     string         `json:"mgts_type"`
}

// MgtsRoutes - api call info
type MgtsRoutes struct {
	MgtsRouteGw   string `json:"mgts_route_gw"`
	MgtsRouteEth  string `json:"mgts_route_eth"`
	MgtsRouteIP   string `json:"mgts_route_ip"`
	MgtsRouteNw   string `json:"mgts_route_nw"`
	MgtsRouteDest string `json:"mgts_route_dest"`
}

// MgtsEth  - api call info
type MgtsEth struct {
	MgtsEthName      string `json:"mgts_eth_name"`
	MgtsEthIP        string `json:"mgts_eth_ip"`
	MgtsEthBootproto string `json:"mgts_eth_bootproto"`
}

// MgtsRtTables  - api call info
type MgtsRtTables struct {
	MgtsRtName   string `json:"mgts_rt_name"`
	MgtsRtNwname string `json:"mgts_rt_nwname"`
	MgtsRtOrder  string `json:"mgts_rt_order,omitempty"`
}

// DefArtLinks  - api call info
type DefArtLinks struct {
	ArtName string `json:"art_name,omitempty"`
	ArtLink string `json:"art_link,omitempty"`
}

/*
SUSE Controller Formular
*/

// SuseCtr suse_c formular
type SuseCtr struct {
	Ctr SuseController `json:"suse_c"`
}

// SuseController - suse_c formular
type SuseController struct {
	APIDescription      string  `json:"api_description"`
	APIEnvironment      string  `json:"api_env"`
	APIServer           string  `json:"api_server"`
	APITitle            string  `json:"api_title"`
	APIVersion          string  `json:"api_version"`
	MgmtDomain          string  `json:"domain_mgmr"`
	PodDomain           string  `json:"domain_pod"`
	InventoryDir        string  `json:"inv_dir"`
	LogDir              string  `json:"log_dir"`
	ScriptsDir          string  `json:"scripts_dir"`
	LoglevelF           string  `json:"loglevel_file"`
	LoglevelS           string  `json:"loglevel_screen"`
	SumaHost            string  `json:"suman_server"`
	SumaUser            string  `json:"suman_user"`
	SumaPassword        string  `json:"suman_password"`
	SumaOrg             string  `json:"suman_organization"`
	HighstateTimeout    string  `json:"suman_timeout"`
	MonitoringTimeout   int     `json:"monitoring_timeout"`
	EventsWait          int     `json:"suman_wait"`
	APIUser             string  `json:"suman_apiuser"`
	RancherHost         string  `json:"rancher_server"`
	RancherVip          string  `json:"rancher_vip"`
	K3sProxy            string  `json:"k3s_proxy"`
	AdapterServer       string  `json:"susea_server"`
	AdapterSecure       bool    `json:"susea_secure"`
	AdapterTokenPwd     string  `json:"susea_token_pw"`
	AdapterRealm        string  `json:"susea_token_realm_name"`
	AdapterUser         string  `json:"susea_token_user"`
	AdapterIdpServer    string  `json:"susea_token_server"`
	OSVersion           string  `json:"os_pod_version"`
	SaltVersion         string  `json:"salt_pod_version"`
	RancherVersion      string  `json:"rancher_version"`
	K3sVersion          string  `json:"k3s_version"`
	KeepalivedVer       string  `json:"keep_alived_version"`
	EcpSumaVersion      string  `json:"ecp_suma_ver"`
	APIVersMgmt         string  `json:"api_vers_mgmt"`
	APIManagement       bool    `json:"api_management"`
	SkipBmc             bool    `json:"skip_bmc,omitempty"`
	RegistrationVersion string  `json:"registration_version"`
	CheckFabricVersion  string  `json:"checkfabric_version"`
	DNS                 DNSInfo `json:"dns,omitempty"`
}

type DNSInfo struct {
	DNSIbUser     string `json:"dns_ib_user,omitempty"`
	DNSIbView     string `json:"dns_ib_view,omitempty"`
	DNSIbIPAMView string `json:"dns_ib_ipamview"`
	DNSProvider   string `json:"dns_provider,omitempty"`
	DNSIbServer   string `json:"dns_ib_server,omitempty"`
	DNSIbNsgroup  string `json:"dns_ib_nsgroup,omitempty"`
	DNSIbPassword string `json:"dns_ib_password,omitempty"`
}

// UyunihubYaml formular containing information for uyunihub
type UyunihubYaml struct {
	ServerMgmt MgmtServer            `yaml:"server"`
	AllObjects map[string]ObjectsSrv `yaml:"channels"`
}

// MgmtServer formular containing information for uyunihub
type MgmtServer struct {
	MasterURL  string `yaml:"hubmaster"`
	MasterUser string `yaml:"user"`
	MasterPw   string `yaml:"password"`
}

// ObjectsSrv formular containing information for uyunihub
type ObjectsSrv struct {
	BaseChannels   []string `yaml:"basechannels,omitempty"`
	CLMProjects    []string `yaml:"projects,omitempty"`
	ConfigChannels []string `yaml:"configchannels,omitempty"`
	Formulas       []string `yaml:"formulas,omitempty"`
}

/*
General Formular
*/

// GeneralFormData General formular
type GeneralFormData struct {
	Data GeneralData `json:"general"`
}

// GeneralLdap - General formular
type GeneralLdap struct {
	Environment         string `json:"environment,omitempty"`
	Domain              string `json:"domain"`
	LdapURI             string `json:"ldap_uri"`
	LdapSearchBase      string `json:"ldap_search_base"`
	LdapDefaultBindDn   string `json:"ldap_default_bind_dn"`
	LdapUserSearchBase  string `json:"ldap_user_search_base"`
	LdapGroupSearchBase string `json:"ldap_group_search_base"`
}

// TasteOS definition - General formular

type TasteOs struct {
	TasteOsCm             bool   `json:"taste_os_cm,omitempty"`
	TasteOsDir            string `json:"taste_os_dir,omitempty"`
	TasteOsPod            bool   `json:"taste_os_pod,omitempty"`
	TasteOsServer         string `json:"taste_os_server,omitempty"`
	TasteOsStoken         string `json:"taste_os_stoken,omitempty"`
	TasteOsSystemTokenCm  string `json:"taste_os_system_token_cm,omitempty"`
	TasteOsSystemTokenPod string `json:"taste_os_system_token_pod,omitempty"`
}

// GeneralData - General formular
type GeneralData struct {
	Ldap       string        `json:"ldap"`
	Root       string        `json:"root"`
	ArtPw      string        `json:"art_pw"`
	ArtUser    string        `json:"art_user"`
	CertsPw    string        `json:"certs_pw"`
	CertsLink  string        `json:"certs_link,omitempty"`
	FleetUser  string        `json:"fleet_user"`
	FleetPw    string        `json:"fleet_pw"`
	PodDocker  string        `json:"poddocker"`
	CICD       string        `json:"cicd_token_salt"`
	EnvSudoers string        `json:"environment_sudoers,omitempty"`
	EcpVersion string        `json:"ecp_version,omitempty"`
	LdapConfig []GeneralLdap `json:"ldap_conf"`
	TasteOs    []TasteOs     `json:"taste_os,omitempty"`
}

/*
DTAG_A4 file
*/

// DtagA4Yaml description of /opt/ecp_suse_controller/dtag_a4.yaml
type DtagA4Yaml struct {
	Suman    SumanSrv   `yaml:"suman"`
	API      APISrv     `yaml:"api"`
	Rancher  RancherSrv `yaml:"rancher"`
	Version  Version    `yaml:"version"`
	SuseA    SuseASrv   `yaml:"suse_a"`
	OpenID   OpenIDSrv  `yaml:"openid"`
	Domain   Domain     `yaml:"domain"`
	Dirs     Dirs       `yaml:"dirs"`
	LogLevel LogLevel   `yaml:"loglevel"`
}

// SumanSrv information needed for accessing SUSE Manager
type SumanSrv struct {
	Server                 string `yaml:"server"`
	User                   string `yaml:"user"`
	Password               string `yaml:"password"`
	Organization           string `yaml:"organization"`
	Timeout                string `yaml:"timeout"`
	WaitBetweenEventsCheck string `yaml:"wait_between_events_check"`
	APIUser                string `yaml:"api_user"`
}

// APISrv information regarding the API
type APISrv struct {
	Title       string `yaml:"title"`
	Description string `yaml:"description"`
	Version     string `yaml:"version"`
	Server      string `yaml:"server"`
	Env         string `yaml:"env"`
}

// RancherSrv information regarding how to access Rancher or SUSE-Tools
type RancherSrv struct {
	VHost    string `yaml:"vhost"`
	VIP      string `yaml:"vip"`
	K3sProxy string `yaml:"k3s_proxy"`
}

// Version information about the needed versions
type Version struct {
	OsPodVersion      string `yaml:"os_pod_version"`
	SaltPodVersion    string `yaml:"salt_pod_version"`
	RancherVersion    string `yaml:"rancher_version"`
	K3sVersion        string `yaml:"k3s_version"`
	KeepAlivedVersion string `yaml:"keep_alived_version"`
}

// SuseASrv - SUSE Adapter info
type SuseASrv struct {
	Server string `yaml:"server"`
	Secure string `yaml:"secure"`
}

// OpenIDSrv - IDP info
type OpenIDSrv struct {
	Server       string `yaml:"server"`
	Realm        string `yaml:"realm"`
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
}

// Domain - DNS domain names used
type Domain struct {
	Pod  string `yaml:"pod"`
	Mgmt string `yaml:"mgmt"`
}

// Dirs - directory used by SUSE_C
type Dirs struct {
	LogDir     string `yaml:"log_dir"`
	ScriptsDir string `yaml:"scripts_dir"`
	InvDir     string `yaml:"inv_dir"`
}

// LogLevel - loglevel definition SUSE_C
type LogLevel struct {
	Screen string `yaml:"screen"`
	File   string `yaml:"file"`
}

/*
File PODDATA.yaml
*/

// SrvMgmt - poddata.yaml
type SrvMgmt struct {
	Server      map[string]ServerDef      `yaml:"server,omitempty"`
	Environment map[string]EnvironmentDef `yaml:"environment,omitempty"`
	Artifactory map[string]ArtifactoryDef `yaml:"artifactory,omitempty"`
}

// ServerDef  - poddata.yaml
type ServerDef struct {
	SrvName        string `yaml:"srvName"`
	SrvTDCNeth     string `yaml:"srvTDCNeth"`
	SrvType        string `yaml:"srvType"`
	SrvNrNic       string `yaml:"srvNrNic"`
	SrvID          string `yaml:"srvID"`
	SrvLocation    string `yaml:"srvLocation"`
	SrvEnvironment string `yaml:"srvEnvironment"`
	SrvEncServerNr string `yaml:"srvEncServerNr"`
	SrvOSRelease   string `yaml:"srvOSRelease"`
	SrvBiosType    string `yaml:"srvBiosType"`
}

// EnvironmentDef  - poddata.yaml
type EnvironmentDef struct {
	DNSServer     string            `yaml:"DNS_SERVER"`
	NTPServer     string            `yaml:"NTP_SERVER"`
	SrvMgmtServ   string            `yaml:"srvMgmtServ"`
	MgtCentralLog string            `yaml:"MGTS_CENTRAL_LOGGING"`
	Eth           map[string]EthDef `yaml:"eth"`
}

// ArtifactoryDef  - poddata.yaml
type ArtifactoryDef struct {
	ArtLink string `yaml:"link,omitempty"`
}

// EthDef  - poddata.yaml
type EthDef struct {
	Pci     string   `yaml:"pci"`
	Net     string   `yaml:"net"`
	GateWay string   `yaml:"gw"`
	Domain  string   `yaml:"dom"`
	RtName  string   `yaml:"rt_name"`
	RtOrder string   `yaml:"rt_order"`
	Routes  []string `yaml:"routes"`
}

// CalcRoute  - poddata.yaml
type CalcRoute struct {
	EthName  string
	EthRoute string
}

// GeneralInit general_init.yaml
type GeneralInit struct {
	Ldap      string                     `yaml:"ldap"`
	Root      string                     `yaml:"root"`
	ArtPw     string                     `yaml:"art_pw"`
	ArtUser   string                     `yaml:"art_user"`
	CertsPw   string                     `yaml:"certs_pw"`
	FleetUser string                     `yaml:"fleet_user"`
	FleetPw   string                     `yaml:"fleet_pw"`
	PodDocker string                     `yaml:"poddocker"`
	CICD      string                     `yaml:"cicd_token_salt"`
	LdapEnv   map[string]GeneralInitLdap `yaml:"ldap_env"`
}

// GeneralInitLdap  - poddata.yaml
type GeneralInitLdap struct {
	Domain              string `yaml:"domain"`
	LdapURI             string `yaml:"ldap_uri"`
	LdapSearchBase      string `yaml:"ldap_search_base"`
	LdapDefaultBindDn   string `yaml:"ldap_default_bind_dn"`
	LdapUserSearchBase  string `yaml:"ldap_user_search_base"`
	LdapGroupSearchBase string `yaml:"ldap_group_search_base"`
}
