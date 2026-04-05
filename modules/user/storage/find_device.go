package storage

import (
	"context"
	"esim/modules/user/model"
)

func (s *SqlModel) FindDevice(ctx context.Context, device string, userID int) (*model.Devide, error) {
	var dev model.Devide
	if err := s.db.Table("device").Where("userid = ? AND device_id = ?", userID, device).First(&dev).Error; err != nil {
		return nil, err
	}
	return &dev, nil
}
