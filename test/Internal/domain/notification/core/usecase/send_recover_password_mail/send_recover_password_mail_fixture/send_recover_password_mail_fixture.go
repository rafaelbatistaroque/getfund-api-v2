package send_recover_password_mail_fixture

import (
	"encoding/json"
	"getfund-api-v2/internal/domain/notification/core/usecase/send_recover_password_mail"
	send_recover_password_mail_application "getfund-api-v2/internal/domain/notification/core/usecase/send_recover_password_mail/application"
	"getfund-api-v2/test/helper/cache_spy"
	"getfund-api-v2/test/helper/mail_spy"
	"getfund-api-v2/test/helper/settings_spy"
	"getfund-api-v2/test/helper/template_file_spy"
)

type recoverPasswordMailModelTest struct {
	Username     string `json:"username"`
	FirstName    string `json:"first_name"`
	RecoveryLink string `json:"recovery_link"`
}

type SendRecoverPasswordMailFixture struct {
	CacheSpy        *cache_spy.RedisCacheSpy
	SettingsSpy     *settings_spy.ApplicationSettingsSpy
	MailSpy         *mail_spy.MailServiceSpy
	TemplateFileSpy *template_file_spy.TemplateFileSpy
}

func NewSUT() (send_recover_password_mail.UseCase, *SendRecoverPasswordMailFixture) {
	cacheSpy := cache_spy.New()
	mailSpy := mail_spy.New()
	settingsSpy := settings_spy.New()
	templateFileSpy := template_file_spy.New()
	return send_recover_password_mail_application.New(cacheSpy, mailSpy, settingsSpy, templateFileSpy),
		&SendRecoverPasswordMailFixture{
			CacheSpy:        cacheSpy,
			MailSpy:         mailSpy,
			SettingsSpy:     settingsSpy,
			TemplateFileSpy: templateFileSpy,
		}

}

func GetInvalidInput() *send_recover_password_mail.Input {
	return &send_recover_password_mail.Input{}
}

func GetValidInput() *send_recover_password_mail.Input {
	return &send_recover_password_mail.Input{KeyCache: "fake-key-cache"}
}

func GetValidRecoverPasswordModeToSendMail() (*recoverPasswordMailModelTest, string) {
	data := &recoverPasswordMailModelTest{
		Username:     "fake-username",
		FirstName:    "fake-first-name",
		RecoveryLink: "recovery-link",
	}
	serializedData, _ := json.Marshal(data)

	return data, string(serializedData)
}
