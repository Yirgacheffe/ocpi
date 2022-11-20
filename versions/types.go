package versions

import "time"

type ModuleID string

const (
	CDRs      ModuleID = "cdrs"
	Locations ModuleID = "locations"
	Tariffs   ModuleID = "tariffs"
	Tokens    ModuleID = "tokens"
	Sessions  ModuleID = "sessions"
)

type VersionNumber string

const (
	OCPI_V_UNKNOWN VersionNumber = "Unknown"
	OCPI_V_2_0     VersionNumber = "2.0"
	OCPI_V_2_1_1   VersionNumber = "2.1.1"
	OCPI_V_2_2_1   VersionNumber = "2.2.1"
)

type Version struct {
	Id        int           `json:"-"`
	VerNumber VersionNumber `json:"version"`
	URL       string        `json:"url"`
	PartyId   string        `json:"-"`
	CreateAt  time.Time     `json:"-"`
	UpdateAt  time.Time     `json:"-"`
}
