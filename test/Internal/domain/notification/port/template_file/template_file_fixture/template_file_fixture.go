package template_file_fixture

import (
	notification_contract "getfund-api-v2/internal/domain/notification/adapter/contract"
	"getfund-api-v2/internal/domain/notification/port/template_file"
	"getfund-api-v2/test/helper/settings_spy"
)

type TemplateFileSpy struct {
	SettingsSpy *settings_spy.ApplicationSettingsSpy
}

func NewSUT() (notification_contract.TemplateFile, *TemplateFileSpy) {
	settingsSpy := settings_spy.New()
	return template_file.New(settingsSpy),
		&TemplateFileSpy{
			SettingsSpy: settingsSpy,
		}
}
