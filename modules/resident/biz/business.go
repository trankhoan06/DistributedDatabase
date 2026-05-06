package biz

import (
	"main.go/config"
	"main.go/modules/resident/model"
)

type ResidentBiz interface {
	FindRegionByID(filePath string, regionID string) (*model.Region, error)
	InsertOneResident(filePath string, newCitizen model.Citizen, regionID string) error
	InsertOneRegion(filePath string, newRegion model.Region) error
	QueryCountry(countryPath string) (int, error)
}
type ResidentCommon struct {
	Resident ResidentBiz
	cfg      *config.Configuration
}

func NewResidentCommon(resident ResidentBiz, cfg *config.Configuration) *ResidentCommon {
	return &ResidentCommon{resident, cfg}
}
