package storage

import (
	"context"
	"esim/modules/user/model"
)

func (s *SqlModel) FindUser(ctx context.Context, cond map[string]interface{}) (*model.User, error) {
	var user model.User
	if err := s.db.Where(cond).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
