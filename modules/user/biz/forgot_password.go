package biz

import (
	"context"
	"errors"
	"esim/common"
	"esim/modules/user/model"
	"esim/worker"
	"github.com/hibiken/asynq"
	"time"
)

func (biz *ForgotCommon) NewForgotPassword(ctx context.Context, data *model.ForgotPassword) error {
	if data.Email == "" {
		return common.ErrEmail(errors.New("email required"))
	}
	_, err := biz.user.FindUser(ctx, map[string]interface{}{"email": data.Email})
	if err != nil {
		return common.ErrNotExist(errors.New("user don't exist"))
	}

	//send email

	taskPayload := worker.PayloadSendEmailForgotPassword{
		Email: data.Email,
	}
	opts := []asynq.Option{
		asynq.MaxRetry(10),
		asynq.ProcessIn(10 * time.Second),
		asynq.Queue(worker.QueueSendResetCodePassword),
		asynq.Unique(1 * time.Minute),
	}
	_ = biz.taskDistributor.DistributeTaskSendEmailForgotPassword(ctx, &taskPayload, opts...)

	return nil
}
