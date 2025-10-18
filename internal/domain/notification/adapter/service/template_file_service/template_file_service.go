package template_file_service

import (
	"errors"
	"getfund-api-v2/internal/config/env"
	"os"
)

const (
	_TEMPLATE_NOTFOUND_MESSAGE   = "template does not exist"
	_RECOVER_PASSWORD_TEMPLATE   = "/recovery_password_template.html"
	_ACTIVATION_ACCOUNT_TEMPLATE = "/activation_account_template.html"
)

type templateFileService struct {
	env env.Variable
}

func New(env env.Variable) *templateFileService {
	return &templateFileService{
		env: env,
	}
}

func (t *templateFileService) GetRecoveryPasswordTemplate() (string, error) {
	template, err := os.ReadFile(t.env.GetTemplateDir() + _RECOVER_PASSWORD_TEMPLATE)
	if err != nil {
		return "", errors.New(_TEMPLATE_NOTFOUND_MESSAGE)
	}

	return string(template), nil
}

func (t *templateFileService) GetActivationAccountTemplate() (string, error) {
	template, err := os.ReadFile(t.env.GetTemplateDir() + _ACTIVATION_ACCOUNT_TEMPLATE)
	if err != nil {
		return "", errors.New(_TEMPLATE_NOTFOUND_MESSAGE)
	}

	return string(template), nil
}
