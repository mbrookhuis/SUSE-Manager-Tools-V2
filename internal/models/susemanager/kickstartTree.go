// Package sumamodels - structs needed for SUSE Manager API Calls
package sumamodels

// KickstartTreeGetDetails distribution details
type KickstartTreeGetDetails struct {
	AbsPath     string `json:"abs_path"`
	ID          int    `json:"id"`
	Label       string `json:"label"`
	ChannelID   int    `json:"channel_id"`
	InstallType struct {
		Name  string `json:"name"`
		ID    int    `json:"id"`
		Label string `json:"label"`
	} `json:"install_type"`
}

type KickstartListProfiles struct {
	UpdateType   string `json:"update_type"`
	Name         string `json:"name"`
	Active       bool   `json:"active"`
	Label        string `json:"label"`
	TreeLabel    string `json:"tree_label"`
	AdvancedMode bool   `json:"advanced_mode"`
	OrgDefault   bool   `json:"org_default"`
}

type KickstartProfileVar struct {
	ProfileName string
	ProfileVars interface{}
}

type ProfileVar struct {
	VarName  string `json:"key"`
	VarValue string `json:"value"`
}
