package codeservice

import (
	"getfund-api-v2/internal/shared/contract/settings"
	"getfund-api-v2/internal/shared/security"
)

type CodeService interface {
	GetRandomCode(length int) (string, error)
}

type codeService struct {
	hasher   security.Hasher
	settings settings.ApplicationSettings
}

func New(hasher security.Hasher, settings settings.ApplicationSettings) CodeService {
	return &codeService{
		hasher:   hasher,
		settings: settings,
	}
}

func (c *codeService) GetRandomCode(length int) (string, error) {
	return "", nil
}
