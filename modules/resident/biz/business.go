package biz

import "main.go/modules/resident/model"

type ResidentBiz interface {
	FindRegionByID(filePath string, regionID string) (*model.Region, error)
	InsertOneResident(filePath string, newCitizen model.Citizen, regionID string) error
	InsertOneRegion(filePath string, newRegion model.Region) error
	QueryCountry(countryPath string) int
}
