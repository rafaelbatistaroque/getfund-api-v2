package send_recover_password_mail_fixture

import (
	"getfund-api-v2/internal/domain/notification/core/usecase/send_recover_password_mail"
	send_recover_password_mail_application "getfund-api-v2/internal/domain/notification/core/usecase/send_recover_password_mail/application"
	"getfund-api-v2/test/helper/cache_spy"
	"getfund-api-v2/test/helper/mail_spy"
	"getfund-api-v2/test/helper/settings_spy"
	"getfund-api-v2/test/helper/template_file_spy"
)

type SendRecoverPasswordMailFixture struct {
	CacheSpy        *cache_spy.RedisCacheSpy
	SettingsSpy     *settings_spy.ApplicationSettingsSpy
	MailSpy         *mail_spy.MailServiceSpy
	TemplateFileSpy *template_file_spy.TemplateFileSpy
}

func NewSUT() (send_recover_password_mail.UseCase, *SendRecoverPasswordMailFixture) {
	mailSpy := mail_spy.New()
	settingsSpy := settings_spy.New()
	templateFileSpy := template_file_spy.New()
	return send_recover_password_mail_application.New(mailSpy, settingsSpy, templateFileSpy),
		&SendRecoverPasswordMailFixture{
			MailSpy:         mailSpy,
			SettingsSpy:     settingsSpy,
			TemplateFileSpy: templateFileSpy,
		}

}

func GetInvalidInput() *send_recover_password_mail.Input {
	return &send_recover_password_mail.Input{}
}

func GetValidInput() *send_recover_password_mail.Input {
	return &send_recover_password_mail.Input{Username: "fake-username", FirstName: "fake-first-name", RecoveryLink: "fake-recovery-link"}
}
