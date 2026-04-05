package storage

import (
	"context"
	"esim/modules/user/model"
)

func (s *SqlModel) CreateDeviceId(ctx context.Context, data *model.CreateDevide) error {
	db := s.db.Begin()
	if err := db.Create(&data).Error; err != nil {
		db.Rollback()
		return err
	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return err
	}
	return nil
}
