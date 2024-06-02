// Package sumamodels - structs needed for SUSE Manager API Calls
package sumamodels

// SubscribedChannel - api call info
type SubscribedChannel struct {
	Summary            string        `json:"summary"`
	MaintainerPhone    string        `json:"maintainer_phone"`
	SupportPolicy      string        `json:"support_policy"`
	EndOfLife          string        `json:"end_of_life"`
	ParentChannelLabel string        `json:"parent_channel_label"`
	GpgKeyURL          string        `json:"gpg_key_url"`
	Description        string        `jsoon:"description"`
	GpgCheck           bool          `json:"gpg_check"`
	Label              string        `json:"label"`
	ArchLabel          string        `json:"arch_label"`
	MaintainerName     string        `json:"maintainer_name"`
	GpgKeyFp           string        `json:"gpg_key_fp"`
	YumrepoLastSync    CustomDate    `json:"yumrepo_last_sync"`
	ContentSources     []interface{} `json:"ContentSources"`
	ArchName           string        `json:"arch_name"`
	Name               string        `json:"name"`
	GpgKeyID           string        `json:"gpg_key_id"`
	MaintainerEmail    string        `json:"maintainer_email"`
	ID                 int           `json:"id"`
	ChecksumLabel      string        `json:"checksum_label"`
	CloneOriginal      string        `json:"clone_original"`
	LastModified       CustomDate    `json:"last_modified"`
}
