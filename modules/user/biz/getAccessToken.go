package biz

import (
	"context"
	"errors"
	"esim/common"
	"esim/component/tokenprovider/jwt"
	modelRefreshToken "esim/modules/refreshToken/model"
	"esim/modules/user/model"
	"time"
)

func (biz *RefreshTokenCommon) NewGetAccessToken(ctx context.Context, data modelRefreshToken.GetRefreshToken) (*jwt.ResponToken, error) {
	ref, err := biz.refresh.FindRefreshToken(ctx, data.Token)
	if err != nil {
		return nil, common.ErrFound(errors.New("the refresh token does not exist"))
	}
	now := time.Now().Add(-7 * time.Hour)
	if ref.ExpireAt.Before(now) {
		return nil, common.ErrExpire(errors.New("refresh token is expired"))
	}
	user, err1 := biz.store.FindUser(ctx, map[string]interface{}{"id": ref.UserId})
	if err1 != nil {
		return nil, common.ErrExist(err1)
	}
	payload := common.Payload{
		Role: user.Role,
		UId:  user.Id,
	}
	accessToken, err := biz.provider.GenerateAccessToken(&payload)
	if err != nil {
		return nil, common.ErrInternal(err)
	}
	token := jwt.ResponToken{
		AccessToken:  accessToken,
		RefreshToken: data.Token,
		User: model.UserSimple{
			Id:    user.Id,
			Role:  user.Role,
			Email: user.Email,
		},
	}
	return &token, nil
}
