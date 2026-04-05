package biz

import (
	"context"
	"errors"
	"esim/common"
	"esim/modules/user/model"
)

func (biz *RegisterCommon) NewRegister(ctx context.Context, data *model.CreateUser) error {
	if data.Email == "" {
		return common.ErrEmailRequest(errors.New("email is request"))
	}
	if data.Account == "" {
		return common.ErrEmailRequest(errors.New("Account is request"))
	}
	if data.Password == "" {
		return common.ErrEmailRequest(errors.New("password is request"))
	}
	if _, err := biz.user.FindUser(ctx, map[string]interface{}{"email": data.Email}); err == nil {
		return common.ErrEmailExist(errors.New("user already exists"))
	}
	if _, err := biz.user.FindUser(ctx, map[string]interface{}{"account": data.Account}); err == nil {
		return common.ErrAccountExist(errors.New("account already exists"))
	}
	data.Token = biz.hash.Hash(data.Email + data.Salt)
	data.Salt = common.GetSalt(50)
	data.Password = biz.hash.Hash(data.Salt + data.Password)
	if err := biz.user.CreateUser(ctx, data); err != nil {
		return common.ErrInternal(err)
	}

	return nil
}
