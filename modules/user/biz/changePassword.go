package biz

import (
	"context"
	"errors"
	"esim/common"
	"esim/modules/user/model"
)

func (biz *RegisterCommon) NewChangePasswork(ctx context.Context, update model.UpdatePass) error {
	user, err := biz.user.FindUser(ctx, map[string]interface{}{"id": update.UserId})
	if err != nil {
		return common.ErrGet(err)
	}
	update.OldPassword = biz.hash.Hash(user.Salt + update.OldPassword)
	update.NewPassword = biz.hash.Hash(user.Salt + update.NewPassword)
	if update.OldPassword != user.Password {
		return common.ErrInvalid(errors.New("old password wrong"))
	}
	if err := biz.user.ChangePassword(ctx, update.UserId, update.NewPassword); err != nil {
		return common.ErrUpdate(err)
	}
	return nil
}
