package models

type ViewMode int

const (
	ViewModeRank = iota
	ViewModeSearch
)

type Site struct {
	// name of site
	SiteName string
	// home page of site
	SiteHomePage string

	// which mode to visit site to get data
	Mode ViewMode
}
