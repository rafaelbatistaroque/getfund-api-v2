package send_recover_password_mail_fixture

import (
	"getfund-api-v2/internal/domain/notification/adapter/usecase/send_recover_password_mail"
	send_recover_password_mail_application "getfund-api-v2/internal/domain/notification/adapter/usecase/send_recover_password_mail/application"
	"getfund-api-v2/test/helper/cache_spy"
	"getfund-api-v2/test/helper/settings_spy"
)

type SendRecoverPasswordMailFixture struct {
	SettingsSpy *settings_spy.ApplicationSettingsSpy
	CacheSpy    *cache_spy.RedisCacheSpy
}

func NewSUT() (send_recover_password_mail.UseCase, *SendRecoverPasswordMailFixture) {
	cacheSpy := cache_spy.New()
	return send_recover_password_mail_application.New(cacheSpy),
		&SendRecoverPasswordMailFixture{
			CacheSpy: cacheSpy,
		}

}

func GetInvalidInput() *send_recover_password_mail.Input {
	return &send_recover_password_mail.Input{}
}

func GetValidInput() *send_recover_password_mail.Input {
	return &send_recover_password_mail.Input{KeyCache: "fake-key-cache"}
}
