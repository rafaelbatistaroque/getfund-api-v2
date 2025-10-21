package config_redis_cache

import (
	"context"
	"encoding/json"
	"getfund-api-v2/internal/config/env"
	"getfund-api-v2/internal/shared/cache"
	logger "getfund-api-v2/internal/shared/log"
	"time"

	"github.com/redis/go-redis/v9"
)

type cacheAdapter struct {
	context context.Context
	redis   *redis.Client
}

func (c *cacheAdapter) Set(key string, value any, time time.Duration) error {
	var data []byte
	var err error

	switch typeData := value.(type) {
	case string:
		data = []byte(typeData)
	default:
		data, err = json.Marshal(value)
		if err != nil {
			return err
		}
	}

	return c.redis.SetEx(c.context, key, data, time).Err()
}

func (c *cacheAdapter) Get(key string) (string, error) {
	content, err := c.redis.Get(c.context, key).Result()
	if err != nil {
		return "", err
	}

	return content, nil
}

func (c *cacheAdapter) Delete(key string) error {
	_, err := c.redis.Del(c.context, key).Result()
	if err != nil {
		return err
	}

	return nil
}

func (c *cacheAdapter) Close() error {
	return c.redis.Close()
}

func New(context context.Context, variable env.Variable) cache.Service {
	logger := logger.New("Redis config")

	redis_client := redis.NewClient(&redis.Options{
		Addr: variable.GetAddrRedis(),
	})

	if _, err := redis_client.Ping(context).Result(); err != nil {
		logger.Errorf("Can't get Redis connection: %v", err)
		return nil
	}

	logger.Info("Redis connected")

	return &cacheAdapter{context: context, redis: redis_client}
}
