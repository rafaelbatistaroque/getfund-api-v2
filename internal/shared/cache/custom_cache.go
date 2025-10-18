package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
)

// Contract defines the interface for a cache service.
// It provides methods for setting, getting, and deleting cache entries,
// as well as closing the connection to the cache backend.
type Contract interface {
	// Set stores a value in the cache with a specific key and expiration time.
	// If the value is not a string, it will be serialized to JSON.
	Set(key string, value any, time time.Duration) error
	// Get retrieves a value from the cache by its key.
	// The value is returned as a string.
	Get(key string) (string, error)
	// Delete removes a key and its associated value from the cache.
	Delete(key string) error
	// Close terminates the connection to the cache service.
	Close() error
}

type customCache struct {
	context context.Context
	redis   *redis.Client
}

// New creates a new cache service instance.
// It takes a Redis client and an application context as dependencies.
func New(redis *redis.Client, context context.Context) Contract {
	return &customCache{
		redis:   redis,
		context: context,
	}
}

func (c *customCache) Set(key string, value any, time time.Duration) error {
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

func (c *customCache) Get(key string) (string, error) {
	content, err := c.redis.Get(c.context, key).Result()
	if err != nil {
		return "", err
	}

	return content, nil
}

func (c *customCache) Delete(key string) error {
	_, err := c.redis.Del(c.context, key).Result()
	if err != nil {
		return err
	}

	return nil
}

func (c *customCache) Close() error {
	return c.redis.Close()
}
