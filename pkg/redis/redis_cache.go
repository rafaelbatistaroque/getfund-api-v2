package redisconfig

import (
	"context"
	"getfund-api-v2/internal/shared/contract/settings"
	applog "getfund-api-v2/pkg/log"

	"github.com/redis/go-redis/v9"
)

func New(context context.Context, settings settings.ApplicationSettings) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: settings.GetAddrRedis(),
	})

	if _, err := client.Ping(context).Result(); err != nil {
		applog.Error.Fatalf("Can't get Redis connection: %v", err)
	}

	applog.Info.Print("Redis connected")

	return client
}
