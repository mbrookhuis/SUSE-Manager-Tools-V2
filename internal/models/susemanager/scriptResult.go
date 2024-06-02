// Package sumamodels - structs needed for SUSE Manager API Calls
package sumamodels

// ScriptResult - api call info
type ScriptResult struct {
	ServerID    string     `json:"serverId"`
	SatrtDate   CustomDate `json:"startDate"`
	StopDate    CustomDate `json:"stopDate"`
	ResturnCode CustomDate `json:"returnCode"`
	Output      string     `json:"output"`
}
