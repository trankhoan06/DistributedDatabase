package storage

import (
	"context"
)

func (s *SqlModel) UpdateVerifyEmail(ctx context.Context, cond map[string]interface{}) error {
	if err := s.db.Table("user").Where(cond).Update("is_email", 1).Error; err != nil {
		return err
	}
	return nil
}
