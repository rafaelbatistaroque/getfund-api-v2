package template_file_service_fixture

import (
	"getfund-api-v2/internal/domain/notification/adapter/service/template_file_service"
	contract "getfund-api-v2/internal/domain/notification/core/contract"
	"getfund-api-v2/test/helper/settings_spy"
)

type TemplateFileSpy struct {
	SettingsSpy *settings_spy.ApplicationSettingsSpy
}

func NewSUT() (contract.TemplateFileService, *TemplateFileSpy) {
	settingsSpy := settings_spy.New()
	return template_file_service.New(settingsSpy),
		&TemplateFileSpy{
			SettingsSpy: settingsSpy,
		}
}
