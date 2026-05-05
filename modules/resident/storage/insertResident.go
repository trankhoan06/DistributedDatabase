package storage

import (
	"encoding/xml"
	"main.go/modules/resident/model"
	"os"
)

func InsertOneResident(filePath string, newCitizen model.Citizen, regionID string) error {
	// 1. Đọc dữ liệu hiện tại
	data, _ := os.ReadFile(filePath)
	var census model.CensusData
	xml.Unmarshal(data, &census)

	// 2. Tìm vùng (Region) phù hợp và thêm người
	for i := range census.Regions {
		if census.Regions[i].ID == regionID {
			census.Regions[i].Citizens = append(census.Regions[i].Citizens, newCitizen)
			break
		}
	}

	// 3. Ghi đè lại file XML (Update)
	output, _ := xml.MarshalIndent(census, "", "    ")
	return os.WriteFile(filePath, output, 0644)
}
