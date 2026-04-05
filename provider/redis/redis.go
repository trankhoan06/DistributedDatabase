package ProviderRedis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"main.go/config"
)

const DefaultMiddleConnect = 0
const DefaultMaxConnect = 10

func NewRedisClient(cfg *config.Config) (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		Password:     cfg.Redis.Password,
		DB:           cfg.Redis.DB,
		PoolSize:     DefaultMiddleConnect,
		MinIdleConns: DefaultMaxConnect,
	})
	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}
	return redisClient, nil
}
