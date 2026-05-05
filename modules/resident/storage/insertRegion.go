package storage

import (
	"encoding/xml"
	"fmt"
	"main.go/modules/resident/model"
	"os"
)

func InsertOneRegion(filePath string, newRegion model.Region) error {
	// 1. Đọc file hiện tại từ folder provider
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var census model.CensusData
	xml.Unmarshal(data, &census)

	// 2. Kiểm tra xem vùng đã tồn tại chưa, nếu chưa thì thêm vào
	for _, r := range census.Regions {
		if r.ID == newRegion.ID {
			return fmt.Errorf("vùng này đã tồn tại")
		}
	}
	census.Regions = append(census.Regions, newRegion)

	// 3. Ghi lại vào file (Update fragment)
	output, _ := xml.MarshalIndent(census, "", "    ")
	return os.WriteFile(filePath, output, 0644)
}
