package biz

import (
	"context"
	"esim/component/tokenprovider"
	"esim/config"
	modelEmail "esim/modules/email/model"
	modelRefreshToken "esim/modules/refreshToken/model"
	"esim/modules/user/model"
	"esim/worker"
)

type UserBiz interface {
	FindUser(ctx context.Context, cond map[string]interface{}) (*model.User, error)
	CreateUser(ctx context.Context, data *model.CreateUser) error
	UpdateUser(ctx context.Context, data *model.UpdateUser) error
	ChangePassword(ctx context.Context, userId int, pass string) error
	FindDevice(ctx context.Context, device string, userID int) (*model.Devide, error)
	CreateDeviceId(ctx context.Context, data *model.CreateDevide) error
}
type Hasher interface {
	Hash(str string) string
}
type RegisterCommon struct {
	user UserBiz
	hash Hasher
}

func NewRegisterCommon(user UserBiz, hash Hasher) *RegisterCommon {
	return &RegisterCommon{user: user, hash: hash}
}

type ForgotCommon struct {
	user            UserBiz
	taskDistributor worker.TaskDistributor
}

func NewForgotCommon(user UserBiz, taskDistributor worker.TaskDistributor) *ForgotCommon {
	return &ForgotCommon{user: user, taskDistributor: taskDistributor}
}

type UserCommon struct {
	user UserBiz
}

func NewUserCommon(user UserBiz) *UserCommon {
	return &UserCommon{user: user}
}

type LoginBiz struct {
	store           UserBiz
	refresh         RefreshTokenBiz
	hash            Hasher
	cfg             *config.Configuration
	provider        tokenprovider.TokenProvider
	taskDistributor worker.TaskDistributor
}

func NewLoginBiz(store UserBiz, refresh RefreshTokenBiz, hash Hasher, cfg *config.Configuration, provider tokenprovider.TokenProvider, taskDistributor worker.TaskDistributor) *LoginBiz {
	return &LoginBiz{store: store, refresh: refresh, hash: hash, cfg: cfg, provider: provider, taskDistributor: taskDistributor}
}

type RefreshTokenBiz interface {
	CreateRefreshToken(ctx context.Context, refreshToken modelRefreshToken.CreateRefreshToken) error
	FindRefreshToken(ctx context.Context, refreshToken string) (*modelRefreshToken.RefreshToken, error)
}
type RefreshTokenCommon struct {
	store    UserBiz
	refresh  RefreshTokenBiz
	provider tokenprovider.TokenProvider
}

func NewRefreshTokenCommon(store UserBiz, refresh RefreshTokenBiz, provider tokenprovider.TokenProvider) *RefreshTokenCommon {
	return &RefreshTokenCommon{store: store, refresh: refresh, provider: provider}
}

type EmailBiz interface {
	CreateCodeVerify(ctx context.Context, data *modelEmail.CreateVerifyAccount) error
	FindCodeVerify(ctx context.Context, cond map[string]interface{}) (*modelEmail.VerifyAccount, error)
	UpdateVerifyCode(ctx context.Context, cond map[string]interface{}, update map[string]interface{}) error
	UpdateVerifyEmail(ctx context.Context, cond map[string]interface{}) error
	UpdateSendCodeEmail(ctx context.Context, data *modelEmail.CreateVerifyAccount) error
}
type VerifyCommon struct {
	user  UserBiz
	email EmailBiz
	hash  Hasher
}

func NewVerifyCommon(store UserBiz, email EmailBiz, hash Hasher) *VerifyCommon {
	return &VerifyCommon{user: store, email: email, hash: hash}
}
