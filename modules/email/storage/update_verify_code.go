package storage

import (
	"context"
)

func (s *SqlModel) UpdateVerifyCode(ctx context.Context, cond map[string]interface{}, update map[string]interface{}) error {
	if err := s.db.Table("send_code").Where(cond).Updates(update).Error; err != nil {
		return err
	}
	return nil
}
