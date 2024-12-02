package cachespy

import (
	"errors"
	"time"
)

// Redis Spy
type RedisCacheSpy struct {
	Params     map[string]interface{}
	CallsCount map[string]int

	SuccessResult map[string]interface{}
	ErrorResult   map[string]error
}

func New() *RedisCacheSpy {
	return &RedisCacheSpy{Params: make(map[string]interface{}), CallsCount: make(map[string]int), SuccessResult: make(map[string]interface{}), ErrorResult: make(map[string]error)}
}

func (r *RedisCacheSpy) Set(key string, value string, time time.Duration) error {
	r.Params["Set:key"] = key
	r.Params["Set:value"] = value
	r.Params["Set:time"] = time

	r.CallsCount["Set"]++

	return r.ErrorResult["Set"]
}
func (r *RedisCacheSpy) DefineRedisSetError() {
	r.ErrorResult["Set"] = errors.New("fake-error")
}

func (r *RedisCacheSpy) Get(key string) (string, error) {
	return "", nil
}

func (r *RedisCacheSpy) Close() error {
	return nil
}
