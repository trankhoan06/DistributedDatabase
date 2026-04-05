package storage

import (
	"context"
	"esim/modules/email/model"
)

func (s *SqlModel) CreateCodeVerify(ctx context.Context, data *model.CreateVerifyAccount) error {
	db := s.db.Begin()
	if err := db.Create(data).Error; err != nil {
		db.Rollback()
		return err
	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return err
	}
	return nil
}
