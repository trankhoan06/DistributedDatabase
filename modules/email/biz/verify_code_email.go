package biz

import (
	"context"
	"errors"
	"esim/common"
	"esim/component/tokenprovider/jwt"
	"esim/modules/email/model"
	modelRefreshToken "esim/modules/refreshToken/model"
	modelUser "esim/modules/user/model"
	"time"
)

func (biz LoginBiz) NewVerifyEmail(ctx context.Context, verify *model.VerifyAccountCode) (*model.ResponseVerifyAccount, error) {
	if verify.Email == "" {
		return nil, common.ErrEmail(errors.New("email require"))
	}
	user, err := biz.user.FindUser(ctx, map[string]interface{}{"email": verify.Email})
	if err != nil {
		return nil, common.ErrEmail(err)
	}
	v, errVerify := biz.store.FindCodeVerify(ctx, map[string]interface{}{"email": user.Email, "type": int(model.TypeVerifyEmail)})
	if v.Verify {
		return nil, errors.New("code has been verified")
	}
	if errVerify != nil {
		return nil, common.ErrRequest(errors.New("verify code error"))
	}
	now := time.Now().Add(-7 * time.Hour)
	if v.Expire.Before(now) {
		return nil, common.ErrVerifyCode(errors.New("verify code expire"))
	}
	if v.Code != verify.Code {
		return nil, common.ErrVerifyCode(errors.New("verify code error"))
	}
	if err := biz.store.UpdateVerifyCode(ctx, map[string]interface{}{"id": user.Id}, map[string]interface{}{"verify": 1}); err != nil {
		return nil, common.ErrUpdate(err)
	}
	if err := biz.store.UpdateVerifyEmail(ctx, map[string]interface{}{"id": user.Id}); err != nil {
		return nil, common.ErrUpdate(err)
	}
	deviceId := common.GenerateDeviceID()
	expire := time.Now().Add(-7 * time.Hour)
	expire = expire.Add(common.TimeDevice * time.Hour)
	dev := modelUser.CreateDevide{
		UserId:   user.Id,
		DeviceId: deviceId,
		UpdateAt: expire,
	}
	if err := biz.user.CreateDeviceId(ctx, &dev); err != nil {
		return nil, common.ErrCreate(err)
	}
	var payload = &common.Payload{
		UId:  user.Id,
		Role: user.Role,
	}
	MyToken, err := biz.provider.GenerateAccessToken(payload)
	MyToken1, err := biz.provider.GenerateRefreshToken(payload)
	createRefre := modelRefreshToken.CreateRefreshToken{
		Token:    MyToken1,
		ExpireAt: now.Add(time.Duration(jwt.TimeRefrestToken) * time.Hour),
		UserId:   payload.GetUser(),
	}
	if err := biz.RefreshToken.CreateRefreshToken(ctx, createRefre); err != nil {
		return nil, common.ErrRefreshToken(err)
	}
	token := jwt.ResponToken{
		AccessToken:  MyToken,
		RefreshToken: MyToken1,
		User: modelUser.UserSimple{
			Id:    user.Id,
			Role:  user.Role,
			Email: user.Email,
		},
	}
	return &model.ResponseVerifyAccount{
		token.RefreshToken,
		token.AccessToken,
		deviceId,
	}, nil
}
