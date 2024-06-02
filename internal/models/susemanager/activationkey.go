// Package sumamodels - structs needed for SUSE Manager API Calls
package sumamodels

// ActivationkeyPackages - packages selection for activation keys
type ActivationkeyPackages struct {
	PackageName string `json:"name"`
	ArchLabel   string `json:"arch"`
}

// ActivationkeyGetDetails - struct with needed information for activation key
type ActivationkeyGetDetails struct {
	Key                string                  `json:"key"`
	Description        string                  `json:"name"`
	UsageLimit         int                     `json:"usage_limit"`
	BaseChannelLabel   string                  `json:"base_channel_label"`
	ChildChannelLabels []string                `json:"child_channel_labels"`
	EntitlementLabel   []string                `json:"entitlements"`
	ServerGroupIds     []int                   `json:"server_group_ids"`
	PackageNames       []string                `json:"package-names"`
	Packages           []ActivationkeyPackages `json:"packages"`
	UniversalDefault   bool                    `json:"universal-default"`
	Disabled           bool                    `json:"disabled"`
	ContactMethod      string                  `json:"contact_method"`
}
