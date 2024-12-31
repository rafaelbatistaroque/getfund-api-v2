package mail

import (
	"getfund-api-v2/internal/shared/contract/settings"
	logger "getfund-api-v2/pkg/log"

	"gopkg.in/gomail.v2"
)

func New(settings settings.ApplicationSettings) (*gomail.Message, *gomail.Dialer) {
	mail := gomail.NewMessage()
	logger := logger.New("SMTP mail config")
	mail.SetHeader("From", settings.GetSMTPFrom())

	dialer := gomail.NewDialer(
		settings.GetSMTPHost(),
		settings.GetSMTPPort(),
		settings.GetSMTPUsername(),
		settings.GetSMTPPassword())

	dialer.SSL = true

	logger.Info("SMTP connection successfully")

	return mail, dialer
}
