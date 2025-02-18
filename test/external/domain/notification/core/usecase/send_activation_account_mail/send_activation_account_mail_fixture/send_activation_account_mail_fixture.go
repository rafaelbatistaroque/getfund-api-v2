package send_activation_account_mail_fixture

import (
	"getfund-api-v2/internal/domain/notification/core/usecase/send_activation_account_mail"
	send_activation_account_mail_application "getfund-api-v2/internal/domain/notification/core/usecase/send_activation_account_mail/application"
	"getfund-api-v2/test/helper/cache_spy"
	"getfund-api-v2/test/helper/mail_spy"
	"getfund-api-v2/test/helper/settings_spy"
	"getfund-api-v2/test/helper/template_file_spy"
)

type SendActivationAccountMailFixture struct {
	CacheSpy        *cache_spy.RedisCacheSpy
	SettingsSpy     *settings_spy.ApplicationSettingsSpy
	MailSpy         *mail_spy.MailServiceSpy
	TemplateFileSpy *template_file_spy.TemplateFileSpy
}

func NewSUT() (send_activation_account_mail.UseCase, *SendActivationAccountMailFixture) {
	mailSpy := mail_spy.New()
	settingsSpy := settings_spy.New()
	templateFileSpy := template_file_spy.New()

	return send_activation_account_mail_application.New(mailSpy, settingsSpy, templateFileSpy),
		&SendActivationAccountMailFixture{
			MailSpy:         mailSpy,
			SettingsSpy:     settingsSpy,
			TemplateFileSpy: templateFileSpy,
		}

}

func GetInvalidInput() *send_activation_account_mail.Input {
	return &send_activation_account_mail.Input{}
}

func GetValidInput() *send_activation_account_mail.Input {
	return &send_activation_account_mail.Input{FirstName: "fake-first-name", Email: "fake@email.com", ActivationLink: "fake-activation-link"}
}
