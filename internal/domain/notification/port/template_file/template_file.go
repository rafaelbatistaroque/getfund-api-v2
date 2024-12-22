package template_file

import (
	"errors"
	"getfund-api-v2/internal/shared/contract/settings"
	"os"
)

type templateFile struct {
	settings settings.ApplicationSettings
}

func New(settings settings.ApplicationSettings) *templateFile {
	return &templateFile{
		settings: settings,
	}
}

func (t *templateFile) GetRecoveryPasswordTemplate() (string, error) {
	_, err := os.ReadFile(t.settings.GetTemplateDir() + "/recovery_password.html")
	if err != nil {
		return "", errors.New("template does not exist")
	}

	return "", nil
}
