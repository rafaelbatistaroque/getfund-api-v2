package cache_spy

import (
	"errors"
	"time"
)

// Redis Spy
type RedisCacheSpy struct {
	Params      map[string]interface{}
	CallsCount  map[string]int
	InvokeOrder []string

	SuccessResult map[string]interface{}
	ErrorResult   map[string]error
}

func New() *RedisCacheSpy {
	return &RedisCacheSpy{Params: make(map[string]interface{}), CallsCount: make(map[string]int), InvokeOrder: []string{}, SuccessResult: make(map[string]interface{}), ErrorResult: make(map[string]error)}
}

func (r *RedisCacheSpy) Set(key string, value interface{}, time time.Duration) error {
	r.Params["Set:key"] = key
	r.Params["Set:value"] = value
	r.Params["Set:time"] = time

	r.CallsCount["Set"]++
	r.incrementOrder("Set")

	return r.ErrorResult["Set"]
}

func (r *RedisCacheSpy) Get(key string) (string, error) {
	r.Params["Get:key"] = key

	r.CallsCount["Get"]++
	r.incrementOrder("Get")

	success := r.SuccessResult["Get"]
	if success != nil {
		return success.(string), r.ErrorResult["Get"]
	}

	return "", r.ErrorResult["Get"]
}

func (r *RedisCacheSpy) Delete(key string) error {
	r.Params["Delete:key"] = key

	r.CallsCount["Delete"]++
	r.incrementOrder("Delete")

	return r.ErrorResult["Delete"]
}

func (r *RedisCacheSpy) Close() error {
	return nil
}

func (r *RedisCacheSpy) DefineCacheSetError()    { r.ErrorResult["Set"] = errors.New("fake-error") }
func (r *RedisCacheSpy) DefineCacheDeleteError() { r.ErrorResult["Delete"] = errors.New("fake-error") }
func (r *RedisCacheSpy) DefineCacheGetError()    { r.ErrorResult["Get"] = errors.New("fake-error") }
func (r *RedisCacheSpy) DefineCacheGetSuccess() {
	r.SuccessResult["Get"] = `{"fakeField": "fake-value"}`
}
func (r *RedisCacheSpy) DefineCacheGetSuccessWithValue(value string) { r.SuccessResult["Get"] = value }
func (r *RedisCacheSpy) incrementOrder(methodName string) {
	r.InvokeOrder = append(r.InvokeOrder, methodName)
}
