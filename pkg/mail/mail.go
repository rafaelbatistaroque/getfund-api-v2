package mail

import (
	"getfund-api-v2/internal/shared/contract/settings"
	logger "getfund-api-v2/pkg/log"

	"gopkg.in/gomail.v2"
)

func New(settings settings.ApplicationSettings) (*gomail.Message, *gomail.Dialer) {
	mail := gomail.NewMessage()
	logger.New("SMTP mail config")
	mail.SetHeader("From", settings.GetSMTPFrom())

	dialer := gomail.NewDialer(
		settings.GetSMTPHost(),
		settings.GetSMTPPort(),
		"",
		"")

	return mail, dialer
}
