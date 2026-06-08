package biz

import (
	"main.go/config"
)

type ResidentBiz interface {
	QueryCountry(countryPath string) (int, error)
}
type ResidentCommon struct {
	Resident ResidentBiz
	cfg      *config.Configuration
}

func NewResidentCommon(resident ResidentBiz, cfg *config.Configuration) *ResidentCommon {
	return &ResidentCommon{resident, cfg}
}
