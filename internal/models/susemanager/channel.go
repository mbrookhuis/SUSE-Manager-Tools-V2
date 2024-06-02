// Package sumamodels - structs needed for SUSE Manager API Calls
package sumamodels

// ChannelListSoftwareChannels - api call info
type ChannelListSoftwareChannels struct {
	Label       string `json:"label"`
	Name        string `json:"name"`
	ParentLabel string `json:"parent_label"`
	EndOfLife   string `json:"end_of_life"`
	Arch        string `json:"arch"`
}
