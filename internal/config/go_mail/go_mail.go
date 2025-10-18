package config_go_mail

import (
	"getfund-api-v2/internal/config/env"
	logger "getfund-api-v2/internal/shared/log"
	"getfund-api-v2/internal/shared/mail"

	"gopkg.in/gomail.v2"
)

func New(variable env.Variable) mail.Contract {
	logger := logger.New("Go mail config")

	message := gomail.NewMessage()
	message.SetHeader("From", variable.GetSMTPFrom())

	dialer := gomail.NewDialer(
		variable.GetSMTPHost(),
		variable.GetSMTPPort(),
		variable.GetSMTPUsername(),
		variable.GetSMTPPassword())

	dialer.SSL = true

	logger.Info("SMTP connection successfully")

	return mail.New(message, dialer)
}
