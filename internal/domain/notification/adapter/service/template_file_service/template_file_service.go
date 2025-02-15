package template_file_service

import (
	"errors"
	"getfund-api-v2/internal/settings"
	"os"
)

const (
	_TEMPLATE_NOTFOUND_MESSAGE   = "template does not exist"
	_RECOVER_PASSWORD_TEMPLATE   = "/recovery_password_template.html"
	_ACTIVATION_ACCOUNT_TEMPLATE = "/activation_account_template.html"
)

type templateFileService struct {
	settings settings.ApplicationSettings
}

func New(settings settings.ApplicationSettings) *templateFileService {
	return &templateFileService{
		settings: settings,
	}
}

func (t *templateFileService) GetRecoveryPasswordTemplate() (string, error) {
	template, err := os.ReadFile(t.settings.GetTemplateDir() + _RECOVER_PASSWORD_TEMPLATE)
	if err != nil {
		return "", errors.New(_TEMPLATE_NOTFOUND_MESSAGE)
	}

	return string(template), nil
}

func (t *templateFileService) GetActivationAccountTemplate() (string, error) {
	template, err := os.ReadFile(t.settings.GetTemplateDir() + _ACTIVATION_ACCOUNT_TEMPLATE)
	if err != nil {
		return "", errors.New(_TEMPLATE_NOTFOUND_MESSAGE)
	}

	return string(template), nil
}
