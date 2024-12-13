package cache_service

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

type Cache interface {
	Set(key string, value interface{}, time time.Duration) error
	Get(key string) (string, error)
	Delete(key string) error
	Close() error
}

type cacheService struct {
	context context.Context
	redis   *redis.Client
}

func New(redis *redis.Client, context context.Context) Cache {
	return &cacheService{
		redis:   redis,
		context: context,
	}
}

func (c *cacheService) Set(key string, value interface{}, time time.Duration) error {
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

func (c *cacheService) Get(key string) (string, error) {
	content, err := c.redis.Get(c.context, key).Result()
	if err != nil {
		return "", err
	}

	return content, nil
}

func (c *cacheService) Delete(key string) error {
	_, err := c.redis.Del(c.context, key).Result()
	if err != nil {
		return err
	}

	return nil
}

func (c *cacheService) Close() error {
	return c.redis.Close()
}
