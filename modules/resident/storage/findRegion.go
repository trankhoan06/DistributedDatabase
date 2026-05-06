package storage

import (
	"encoding/xml"
	"fmt"
	"main.go/modules/resident/model"
	"os"
)

func (sql *ResidentFile) FindRegionByID(filePath string, regionID string) (*model.Region, error) {
	// 1. Đọc file từ folder provider tương ứng
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("không thể mở file: %v", err)
	}

	var census model.CensusData
	err = xml.Unmarshal(data, &census)
	if err != nil {
		return nil, fmt.Errorf("lỗi parse XML: %v", err)
	}

	// 2. Duyệt tìm Region theo ID
	for _, r := range census.Regions {
		if r.ID == regionID {
			return &r, nil
		}
	}

	return nil, fmt.Errorf("không tìm thấy vùng có ID: %s", regionID)
}
