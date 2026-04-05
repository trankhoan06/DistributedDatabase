package storage

import (
	"context"
	"esim/modules/email/model"
)

func (s *SqlModel) UpdateSendCodeEmail(ctx context.Context, data *model.CreateVerifyAccount) error {
	if err := s.db.Where("email=?", data.Email).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
