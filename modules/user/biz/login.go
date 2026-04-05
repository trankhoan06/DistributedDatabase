package biz

import (
	"context"
	"errors"
	"esim/common"
	"esim/component/tokenprovider/jwt"
	modelRefreshToken "esim/modules/refreshToken/model"
	"esim/modules/user/model"
	"esim/worker"
	"github.com/hibiken/asynq"
	"gorm.io/gorm"
	"time"
)

func (biz *LoginBiz) NewLogin(ctx context.Context, data *model.Login, db *gorm.DB) (*jwt.ResponToken, error) {
	if data.Account == "" {
		return nil, common.ErrAccount(errors.New("account is request"))
	}
	user, err := biz.store.FindUser(ctx, map[string]interface{}{"account": data.Account})
	if err != nil {
		return nil, common.ErrLogin(errors.New("account or password wrong"))
	}
	pass := biz.hash.Hash(user.Salt + data.Password)
	if pass != user.Password {
		return nil, common.ErrLogin(errors.New("account or password wrong"))
	}
	dev, err := biz.store.FindDevice(ctx, data.DeviceId, user.Id)
	if err != nil {
		return nil, common.ErrLogin(errors.New("device or password wrong"))
	}
	now := time.Now().Add(-7 * time.Hour)
	if dev.UpdateAt.Before(now) {
		taskPayload := worker.PayloadSendEmailForgotPassword{
			Email: user.Email,
		}
		opts := []asynq.Option{
			asynq.MaxRetry(10),
			asynq.ProcessIn(10 * time.Second),
			asynq.Queue(worker.QueueSendVerifyEmail),
			asynq.Unique(1 * time.Minute),
		}
		_ = biz.taskDistributor.DistributeTaskSendEmailForgotPassword(ctx, &taskPayload, opts...)
		return nil, common.NewAuthorize(nil, "device_id expried", "device_id expried", "ERR_DEVICEID")
	}
	payload := &common.Payload{
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
	if err := biz.refresh.CreateRefreshToken(ctx, createRefre); err != nil {
		return nil, common.ErrRefreshToken(err)
	}
	token := jwt.ResponToken{
		AccessToken:  MyToken,
		RefreshToken: MyToken1,
		User: model.UserSimple{
			Id:    user.Id,
			Role:  user.Role,
			Email: user.Email,
		},
	}
	return &token, nil
}
