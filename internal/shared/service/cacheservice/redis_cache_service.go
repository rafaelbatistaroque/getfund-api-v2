package cacheservice

import (
	"context"
	"getfund-api-v2/internal/shared/contract/settings"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type Cache interface {
	Set(key string, value string, time time.Duration) error
	Get(key string) (string, error)
	Delete(key string) (string, error)
	Close() error
}

type redisCache struct {
	context  context.Context
	settings settings.ApplicationSettings
	redis    *redis.Client
}

func New(context context.Context, settings settings.ApplicationSettings) Cache {
	client := redis.NewClient(&redis.Options{
		Addr: settings.GetAddrRedis(),
	})

	if _, err := client.Ping(context).Result(); err != nil {
		log.Fatalf("Can't get Redis connection: %v", err)
	}

	log.Printf("Redis connected")

	return &redisCache{
		context:  context,
		settings: settings,
		redis:    client,
	}
}

func (c *redisCache) Set(key string, value string, time time.Duration) error {
	return c.redis.SetEx(c.context, key, value, time).Err()
}

func (c *redisCache) Get(key string) (string, error) {
	content, err := c.redis.Get(c.context, key).Result()
	if err != nil {
		return "", err
	}

	return content, nil
}

func (c *redisCache) Delete(key string) (string, error) {
	return "", nil
}

func (c *redisCache) Close() error {
	return c.redis.Close()
}
