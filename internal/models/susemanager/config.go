// Package sumamodels - structs needed for SUSE Manager API Calls
package sumamodels

// ConfigChannelType - api call info
type ConfigChannelType struct {
	CCId    int    `json:"id"`
	CCLabel string `json:"label"`
	CCName  string `json:"name"`
	CCPrio  int    `json:"priority"`
}

// ConfigChannelListGlobals - api call info
type ConfigChannelListGlobals struct {
	ID          int               `json:"id"`
	OrgID       int               `json:"orgId"`
	Label       string            `json:"label"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Type        string            `json:"type"`
	Arch        string            `json:"arch"`
	ChannelType ConfigChannelType `json:"configChannelType"`
}
