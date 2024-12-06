package redisconfig

import (
	"context"
	"getfund-api-v2/internal/shared/contract/settings"
	logger "getfund-api-v2/pkg/log"

	"github.com/redis/go-redis/v9"
)

func New(context context.Context, settings settings.ApplicationSettings) *redis.Client {
	logger := logger.New("Redis config")
	client := redis.NewClient(&redis.Options{
		Addr: settings.GetAddrRedis(),
	})

	if _, err := client.Ping(context).Result(); err != nil {
		logger.Errorf("Can't get Redis connection: %v", err)
	}

	logger.Info("Redis connected")

	return client
}
