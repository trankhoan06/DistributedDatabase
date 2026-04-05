package biz

import (
	"context"
	"esim/component/tokenprovider"
	"esim/modules/email/model"
	modelRefreshToken "esim/modules/refreshToken/model"
	modelUser "esim/modules/user/model"
)

type EmailBiz interface {
	CreateCodeVerify(ctx context.Context, data *model.CreateVerifyAccount) error
	FindCodeVerify(ctx context.Context, cond map[string]interface{}) (*model.VerifyAccount, error)
	UpdateVerifyCode(ctx context.Context, cond map[string]interface{}, update map[string]interface{}) error
	UpdateVerifyEmail(ctx context.Context, cond map[string]interface{}) error
	UpdateSendCodeEmail(ctx context.Context, data *model.CreateVerifyAccount) error
}
type UserEmailBiz interface {
	FindUser(ctx context.Context, cond map[string]interface{}) (*modelUser.User, error)
	CreateDeviceId(ctx context.Context, data *modelUser.CreateDevide) error
}
type RefreshTokenBiz interface {
	CreateRefreshToken(ctx context.Context, refreshToken modelRefreshToken.CreateRefreshToken) error
	FindRefreshToken(ctx context.Context, refreshToken string) (*modelRefreshToken.RefreshToken, error)
}
type UserCommonBiz struct {
	store EmailBiz
}

func NewUserCommonBiz(store EmailBiz) *UserCommonBiz {
	return &UserCommonBiz{store: store}
}

type SendEMailBiz struct {
	store EmailBiz
	user  UserEmailBiz
}

func NewSendEmailBiz(store EmailBiz, user UserEmailBiz) *SendEMailBiz {
	return &SendEMailBiz{store: store, user: user}
}

type Hasher interface {
	Hash(str string) string
}
type RegisterEmailBiz struct {
	store EmailBiz
	hash  Hasher
}

func NewRegisterEmailBiz(store EmailBiz, hash Hasher) *RegisterEmailBiz {
	return &RegisterEmailBiz{store: store, hash: hash}
}

type LoginBiz struct {
	store        EmailBiz
	user         UserEmailBiz
	provider     tokenprovider.TokenProvider
	hash         Hasher
	RefreshToken RefreshTokenBiz
}

func NewLoginBiz(store EmailBiz, user UserEmailBiz, provider tokenprovider.TokenProvider, hash Hasher, RefreshToken RefreshTokenBiz) *LoginBiz {
	return &LoginBiz{store: store, user: user, provider: provider, hash: hash, RefreshToken: RefreshToken}
}
