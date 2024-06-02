// Package sumamodels - structs needed for SUSE Manager API Calls
package sumamodels

// Slaves - api call info
type Slaves struct {
	AllowedAllOrgs bool   `json:"allowAllOrgs"`
	Enabled        bool   `json:"enabled"`
	ID             int    `json:"id"`
	Label          string `json:"label"`
}

// SlavesIssMaster - api call info
type SlavesIssMaster struct {
	Master bool   `json:"isCurrentMaster"`
	CaCert string `json:"caCert"`
	ID     int    `json:"id"`
	Label  string `json:"label"`
}
