package storage

import (
	"context"
)

func (s *SqlModel) ChangePassword(ctx context.Context, userId int, pass string) error {
	if err := s.db.Table("user").Where("id=?", userId).Update("password", pass).Error; err != nil {
		return err
	}
	return nil
}
