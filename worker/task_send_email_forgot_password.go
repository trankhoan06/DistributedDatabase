package worker

import (
	"context"
	"encoding/json"
	modelEmail "esim/modules/email/model"
	modelUser "esim/modules/user/model"
	"fmt"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog/log"
)

const TaskForgotPassword = "task:forgot_password"

type PayloadSendEmailForgotPassword struct {
	Email string `json:"email"`
}

func (dis *RedisTaskDistributor) DistributeTaskSendEmailForgotPassword(
	ctx context.Context,
	payload *PayloadSendEmailForgotPassword,
	opts ...asynq.Option,
) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error when marshal payload: %v", err)
	}
	task := asynq.NewTask(TaskForgotPassword, jsonPayload, opts...)
	info, err := dis.Client.EnqueueContext(ctx, task)
	if err != nil {
		return fmt.Errorf("error when enqueue task: %v", err)
	}
	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
		Str("queue", info.Queue).Int("max_retry", info.MaxRetry).Msg("enqueued task")
	return nil
}
func (pro *RedisTaskProcessor) TaskProcessSendEmailForgotPassword(ctx context.Context, task *asynq.Task) error {
	var payload PayloadSendEmailForgotPassword
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return err
	}
	user, err := pro.accountStorage.FindUser(ctx, map[string]interface{}{"email": payload.Email})
	if err != nil {
		return fmt.Errorf("error user: %w", err)
	}
	Type := modelEmail.TypeForgotPassword
	verify, err := pro.accountCase.NewResendCodeEmail(ctx, payload.Email, 5*60, &Type)
	if err != nil {
		return fmt.Errorf("error when update resend verify email: %w", err)
	}
	if err := SendMailToForgotPassword(pro, verify, user); err != nil {
		return fmt.Errorf("error when send verify email: %w", err)
	}

	log.Info().Msg("send verify email success")

	log.Info().Str("type", task.Type()).Bytes("payload", task.Payload()).
		Str("email", payload.Email).Msg("processed task")
	return nil
}
func SendMailToForgotPassword(processor *RedisTaskProcessor, verifyEmail *modelEmail.CreateVerifyAccount, user *modelUser.User) error {
	subject := "Welcome to Motel Room"
	content := fmt.Sprintf(`Hello %s,<br/>
	This is reset code for you: %d<br/>
	`, user.LastName, verifyEmail.Code)
	to := []string{verifyEmail.Email}

	err := processor.Mailer.SendEmail(subject, content, to, nil, nil, "", nil)
	if err != nil {
		return fmt.Errorf("failed to send verify email: %w", err)
	}
	return nil
}
