// Package sumamodels - structs needed for SUSE Manager API Calls
package sumamodels

// System - api call info
type System struct {
	ID               int        `json:"id"`
	LastChekin       CustomDate `json:"last_checkin"`
	Name             string     `json:"name"`
	OutdatedPkgCount int        `json:"outdated_pkg_count"`
}

// ActiveSystem - api call info
type ActiveSystem struct {
	ID         int        `json:"id"`
	Name       string     `json:"name"`
	LastChekin CustomDate `json:"last_checkin"`
	Created    CustomDate `json:"created"`
	LastBoot   CustomDate `json:"last_boot"`
}

// SubscribedBaseChannel - api call info
type SubscribedBaseChannel struct {
	Summary            string        `json:"summary"`
	MaintainerPhone    string        `json:"maintainer_phone"`
	SupportPolicy      string        `json:"support_policy"`
	EndOfLife          string        `json:"end_of_life"`
	ParentChannelLabel string        `json:"parent_channel_label"`
	GpgKeyURL          string        `json:"gpg_key_url"`
	Description        string        `json:"description"`
	GpgCheck           bool          `json:"gpg_check"`
	Label              string        `json:"label"`
	ArchLabel          string        `json:"arch_label"`
	MaintainerName     string        `json:"maintainer_name"`
	GpgKeyFp           string        `json:"gpg_key_fp"`
	ContentSources     []interface{} `json:"contentSources"`
	ArchName           string        `json:"arch_name"`
	Name               string        `json:"name"`
	GpgKeyID           string        `json:"gpg_key_id"`
	MaintainerEmail    string        `json:"maintainer_email"`
	ID                 int           `json:"id"`
	ChecksumLabel      string        `json:"checksum_label"`
	CloneOriginal      string        `json:"clone_original"`
	LastModified       string        `json:"last_modified"`
}
