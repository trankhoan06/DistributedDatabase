package storage

import (
	"context"
	"esim/modules/user/model"
)

func (s *SqlModel) CreateUser(ctx context.Context, data *model.CreateUser) error {
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
