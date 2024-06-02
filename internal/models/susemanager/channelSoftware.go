// Package sumamodels - structs needed for SUSE Manager API Calls
package sumamodels

// ChannelSoftwareContentSource - api call info
type ChannelSoftwareContentSource struct {
	ID        int    `json:"id"`
	Label     string `json:"label"`
	SourceURL string `json:"source_url"`
	Type      string `json:"type"`
}

// ChannelSoftwareListChildren - api call info
type ChannelSoftwareListChildren struct {
	ID                 int                            `json:"id"`
	Name               string                         `json:"name"`
	Label              string                         `json:"label"`
	ArchName           string                         `json:"arch_nane"`
	ArchLabel          string                         `json:"arch_label"`
	Summary            string                         `json:"summary"`
	Description        string                         `json:"description"`
	ChecksumLabel      string                         `json:"checksum_label"`
	LastModified       CustomDate                     `json:"last_modified,omitempty"`
	MaintainerName     string                         `json:"maintainer_name"`
	MaintainerEmail    string                         `json:"maintainer_email"`
	MaintainerPhone    string                         `json:"maintainer_phone"`
	SupportPolicy      string                         `json:"support_policy"`
	GpgKeyURL          string                         `json:"gpg_key_url"`
	GpgKeyID           string                         `json:"gpg_key_id"`
	GpgKeyfp           string                         `json:"gpg_key_fp"`
	GpgCheck           bool                           `json:"gpg_check,omitempty"`
	YumrepoLastSync    CustomDate                     `json:"yumrepo_last_sync,omitempty"`
	EndOfLife          string                         `json:"end_of_life"`
	ParentChannelLabel string                         `json:"parent_channel_label"`
	CloneOriginal      string                         `json:"clone_original"`
	ContentSource      []ChannelSoftwareContentSource `json:"content_source,omitempty"`
}

type ChannelSoftwareCreateRepo struct {
	SourceURL         string `json:"sourceUrl,omitempty"`
	ID                int    `json:"id,omitempty"`
	Label             string `json:"label,omitempty"`
	Type              string `json:"type,omitempty"`
	HasSignedMetadata bool   `json:"hasSignedMetadata,omitempty"`
	SslContentSources []struct {
		SslKeyDesc  string `json:"sslKeyDesc,omitempty"`
		SslCaDesc   string `json:"sslCaDesc,omitempty"`
		SslCertDesc string `json:"sslCertDesc,omitempty"`
	} `json:"sslContentSources,omitempty"`
}
