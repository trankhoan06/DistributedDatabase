package worker

import (
	"context"
	"github.com/hibiken/asynq"
)

type TaskDistributor interface {
	DistributeTaskSendEmailForgotPassword(
		ctx context.Context,
		payload *PayloadSendEmailForgotPassword,
		opts ...asynq.Option,
	) error
	DistributeTaskSendVerifyEmail(
		ctx context.Context,
		payload *PayloadSendVerifyEmail,
		opts ...asynq.Option,
	) error
}

type RedisTaskDistributor struct {
	Client *asynq.Client
}

func NewRedisTaskDistributor(redisOpt *asynq.RedisClientOpt) TaskDistributor {
	client := asynq.NewClient(redisOpt)
	return &RedisTaskDistributor{Client: client}
}
