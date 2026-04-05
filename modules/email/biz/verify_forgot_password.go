package biz

import (
	"context"
	"errors"
	"esim/common"
	"esim/modules/email/model"
	"time"
)

func (biz SendEMailBiz) NewVerifyForgotPassword(ctx context.Context, password *model.VerifyAccountCode, expire int) error {
	if password.Email == "" {
		return common.ErrEmail(errors.New("email required"))
	}
	user, err := biz.user.FindUser(ctx, map[string]interface{}{"email": password.Email})
	if err != nil {
		return common.ErrAccount(errors.New("user not found"))
	}
	v, errVerify := biz.store.FindCodeVerify(ctx, map[string]interface{}{"email": user.Email, "type": int(model.TypeForgotPassword)})
	if errVerify != nil {
		return errVerify
	}
	if v.Verify {
		return common.ErrVerifyCode(errors.New("code has been verified"))
	}
	if v.Code != password.Code {
		return common.ErrVerifyCode(errors.New("code wrong"))
	}
	now := time.Now()
	now = now.Add(-7 * time.Hour)
	if v.Expire.Before(now) {
		return common.ErrVerifyCode(errors.New("expire time"))
	}
	now = now.Add(time.Duration(expire) * time.Second)
	if err := biz.store.UpdateVerifyCode(ctx, map[string]interface{}{"email": user.Email}, map[string]interface{}{"verify": 1, "expire": now}); err != nil {
		return common.ErrUpdate(err)
	}
	return nil
}
