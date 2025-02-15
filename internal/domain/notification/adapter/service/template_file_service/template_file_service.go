package template_file_service

import (
	"errors"
	"getfund-api-v2/internal/settings"
	"os"
)

const (
	_TEMPLATE_NOTFOUND_MESSAGE = "template does not exist"
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
	template, err := os.ReadFile(t.settings.GetTemplateDir() + "/recovery_password_template.html")
	if err != nil {
		return "", errors.New(_TEMPLATE_NOTFOUND_MESSAGE)
	}

	return string(template), nil
}

func (t *templateFileService) GetActivationAccountTemplate() (string, error) {
	template, err := os.ReadFile(t.settings.GetTemplateDir() + "/activation_account_template.html")
	if err != nil {
		return "", errors.New(_TEMPLATE_NOTFOUND_MESSAGE)
	}

	return string(template), nil
}
