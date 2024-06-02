// Package sumamodels - structs needed for SUSE Manager API Calls
package sumamodels

// InstallablePackage - api call info
type InstallablePackage struct {
	Name      string `json:"name"`
	Version   string `json:"version"`
	Release   string `json:"release"`
	Epoch     string `json:"epoch"`
	ID        int    `json:"id"`
	ArchLabel string `json:"arch_label"`
}

// InstalledPackage - api call info
type InstalledPackage struct {
	PackageID int    `json:"package_id"`
	Name      string `json:"name"`
	Version   string `json:"version"`
	Release   string `json:"release"`
	Epoch     string `json:"epoch"`
	Arch      string `json:"arch"`
	Retracted bool   `json:"retracted"`
}
