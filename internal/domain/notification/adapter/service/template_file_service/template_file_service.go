package template_file_service

import (
	"errors"
	"getfund-api-v2/internal/settings"
	"os"
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
		return "", errors.New("template does not exist")
	}

	return string(template), nil
}
