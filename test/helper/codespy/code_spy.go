package codespy

import (
	"errors"
)

type CodeSpy struct {
	Params        map[string]interface{}
	CallsCount    map[string]int
	ErrorResult   map[string]error
	SuccessResult map[string]interface{}
}

func New() *CodeSpy {
	return &CodeSpy{Params: make(map[string]interface{}), CallsCount: make(map[string]int), ErrorResult: make(map[string]error), SuccessResult: make(map[string]interface{})}
}

func (c *CodeSpy) GetRandomCode(length int) (string, error) {
	c.Params["GetRandomCode:length"] = length

	c.CallsCount["GetRandomCode"]++

	success := c.SuccessResult["GetRandomCode"]
	if success != nil {
		return success.(string), nil
	}

	return "", c.ErrorResult["GetRandomCode"]
}

func (c *CodeSpy) DefineGetRandomCodeError() {
	c.ErrorResult["GetRandomCode"] = errors.New("fake-error")
}

func (c *CodeSpy) DefineGetRandomCodeSuccess() {
	c.SuccessResult["GetRandomCode"] = "fake-code"
}
