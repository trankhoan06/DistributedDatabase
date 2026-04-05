package biz

import (
	"context"
	"errors"
	"esim/common"
	"esim/modules/user/model"
	"time"
)

func (biz *VerifyCommon) NewChangePasswordForgot(ctx context.Context, password *model.NewPasswordForgot) error {
	if password.Email == "" {
		return common.ErrEmail(errors.New("email required"))
	}
	user, err := biz.user.FindUser(ctx, map[string]interface{}{"email": password.Email})
	if err != nil {
		return err
	}
	v, errVerify := biz.email.FindCodeVerify(ctx, map[string]interface{}{"email": user.Email})
	if errVerify != nil {
		return errVerify
	}
	if !v.Verify {
		return common.ErrVerifyCode(errors.New("user don't verify code"))
	}
	now := time.Now().Add(-7 * time.Hour)
	if v.Expire.Before(now) {
		return common.ErrVerifyCode(errors.New("expire time"))
	}
	password.NewPassword = biz.hash.Hash(user.Salt + password.NewPassword)
	if err := biz.user.ChangePassword(ctx, user.Id, password.NewPassword); err != nil {
		return err
	}
	if err := biz.email.UpdateVerifyCode(ctx, map[string]interface{}{"email": user.Email}, map[string]interface{}{"expire": now}); err != nil {
		return err
	}

	return nil
}
