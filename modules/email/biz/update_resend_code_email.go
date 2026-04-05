package biz

import (
	"context"
	"errors"
	"esim/common"
	"esim/modules/email/model"
	"time"
)

func (biz *SendEMailBiz) NewResendCodeEmail(ctx context.Context, email string, expire int, Type *model.TypeCode) (*model.CreateVerifyAccount, error) {
	//check user exist
	_, err := biz.user.FindUser(ctx, map[string]interface{}{"email": email})
	if err != nil {
		return nil, common.ErrEmail(err)
	}
	email1, err := biz.store.FindCodeVerify(ctx, map[string]interface{}{"email": email})
	now := time.Now().Add(-7 * time.Hour)
	if err == nil {
		if email1.Expire.Add(time.Duration(5) * time.Minute).After(now) {
			return nil, common.ErrResendEmail(errors.New("Send emails every 5 minutes."))
		}
	}

	var verifyEmail model.CreateVerifyAccount
	verifyEmail.Email = email
	verifyEmail.Type = Type
	verifyEmail.Verify = false
	verifyEmail.Code = common.GenerateRandomCode()
	verifyEmail.Expire = now.Add(time.Duration(expire) * time.Second)
	if err := biz.store.UpdateSendCodeEmail(ctx, &verifyEmail); err != nil {
		return nil, common.ErrUpdate(err)
	}
	return &verifyEmail, nil
}
