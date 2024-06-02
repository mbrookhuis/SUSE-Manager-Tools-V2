// Package sumamodels - structs needed for SUSE Manager API Calls
package sumamodels

// K3SConfig - api call info
type K3SConfig struct {
	K3sConfig DtagK3s `json:"k3sconfig"`
}

// DtagK3s - api call info
type DtagK3s struct {
	K3sAPI             string      `json:"k3s_api"`
	K3sLocation        string      `json:"k3s_location"`
	K3sPrimaryIP       string      `json:"k3s_primary_ip"`
	K3sServer          []K3sServer `json:"k3s_server"`
	K3sToken           string      `json:"k3s_token"`
	K3sVersion         string      `json:"k3s_version"`
	K3sVip             string      `json:"k3s_vip"`
	RancherVersion     string      `json:"rancher_version"`
	OsInstalledVersion string      `json:"os_pod_version"`
}

// K3sServer - api call info
type K3sServer struct {
	PodIPK3s string `json:"pod_ip_k3s"`
	PodName  string `json:"pod_name"`
	PodRole  string `json:"pod_role"`
}
