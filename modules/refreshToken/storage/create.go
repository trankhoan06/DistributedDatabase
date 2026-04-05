package storage

import (
	"context"
	modelRefreshToken "esim/modules/refreshToken/model"
)

func (s *sqlModel) CreateRefreshToken(ctx context.Context, refreshToken modelRefreshToken.CreateRefreshToken) error {
	db := s.db.Begin()
	if err := db.Create(&refreshToken).Error; err != nil {
		db.Rollback()
		return err
	}
	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return err
	}
	return nil
}
