package storage

import (
	"context"
	modelRefreshToken "esim/modules/refreshToken/model"
)

func (s *sqlModel) FindRefreshToken(ctx context.Context, refreshToken string) (*modelRefreshToken.RefreshToken, error) {
	var data modelRefreshToken.RefreshToken
	if err := s.db.Where("token=?", refreshToken).Last(&data).Error; err != nil {
		return nil, err
	}
	return &data, nil
}
