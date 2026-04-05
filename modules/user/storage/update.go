package storage

import (
	"context"
	"esim/modules/user/model"
)

func (s *SqlModel) UpdateUser(ctx context.Context, data *model.UpdateUser) error {
	if err := s.db.Where("email=?", data.Email).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
