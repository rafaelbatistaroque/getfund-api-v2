package config_redis_cache

import (
	"context"
	"getfund-api-v2/internal/config/env"
	"getfund-api-v2/internal/shared/cache"
	logger "getfund-api-v2/internal/shared/log"

	"github.com/redis/go-redis/v9"
)

func New(context context.Context, variable env.Variable) cache.Contract {
	logger := logger.New("Redis config")

	redis_client := redis.NewClient(&redis.Options{
		Addr: variable.GetAddrRedis(),
	})

	if _, err := redis_client.Ping(context).Result(); err != nil {
		logger.Errorf("Can't get Redis connection: %v", err)
		return nil
	}

	logger.Info("Redis connected")

	return cache.New(redis_client, context)
}
