// Package sumamodels - structs needed for SUSE Manager API Calls
package sumamodels

// SystemGroupGetDetails - api call info
type SystemGroupGetDetails struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	OrgID       int    `json:"org_id"`
	SystemCount int    `json:"system_count"`
}

// SystemGroupListSystemsMinimal - api call info
type SystemGroupListSystemsMinimal struct {
	ID               int        `json:"id"`
	LastChekin       CustomDate `json:"last_checkin"`
	Created          CustomDate `json:"created"`
	LastBoot         CustomDate `json:"last_boot"`
	Name             string     `json:"name"`
	OutdatedPkgCount int        `json:"outdated_pkg_count"`
	ExtraPkgCount    int        `json:"extra_pkg_count"`
}
