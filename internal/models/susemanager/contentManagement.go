// Package sumamodels - structs needed for SUSE Manager API Calls
package sumamodels

import "time"

// ContentManagementListProjects ContentManagement ListProjects output
type ContentManagementListProjects struct {
	ID               int        `json:"id"`
	Label            string     `json:"label"`
	Name             string     `json:"name"`
	Description      string     `json:"description"`
	LastBuildDate    CustomDate `json:"lastBuildDate"`
	FirstEnvironment string     `json:"firstEnvironment"`
	OrgID            int        `json:"orgId"`
}

// ContentManagementSource ContentManagement addSource output
type ContentManagementSource struct {
	ContentProjectLabel string `json:"contentProjectLabel"`
	Type                string `json:"type"`
	State               string `json:"state"`
	ChannelLabel        string `json:"channelLabel"`
}

// ContentManagementFilter ContentManagement filter output
type ContentManagementFilter struct {
	EntityType string                          `json:"entityType"`
	Criteria   ContentManagementFilterCriteria `json:"criteria"`
	Name       string                          `json:"name"`
	Rule       string                          `json:"rule"`
	ID         int                             `json:"id"`
	OrgID      int                             `json:"orgId"`
}

// ContentManagementFilterCriteria ContentManagement FilterCriteria output
type ContentManagementFilterCriteria struct {
	Field   string    `json:"field"`
	Matcher string    `json:"matcher"`
	Value   time.Time `json:"value"`
}

// FilterCriteria ContentManagement FilterCriteria output
type FilterCriteria struct {
	Field   string `json:"field"`
	Matcher string `json:"matcher"`
	Value   string `json:"value"`
}

// ContentManagementEnvironment ContentManagement Environment output
type ContentManagementEnvironment struct {
	Name                string `json:"name"`
	ContentProjectLabel string `json:"contentProjectLabel"`
	ID                  int    `json:"id"`
	Label               string `json:"label"`
	Version             int    `json:"version"`
	Status              string `json:"status"`
}
