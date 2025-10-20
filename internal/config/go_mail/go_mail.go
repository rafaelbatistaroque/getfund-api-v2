package config_go_mail

import (
	"fmt"
	"getfund-api-v2/internal/config/env"
	logger "getfund-api-v2/internal/shared/log"
	"getfund-api-v2/internal/shared/mail"

	"gopkg.in/gomail.v2"
)

// --- Message Adapter ---
type gomailMessageAdapter struct {
	*gomail.Message
}

func (g *gomailMessageAdapter) SetHeader(key string, value ...string) {
	g.Message.SetHeader(key, value...)
}

func (g *gomailMessageAdapter) SetBody(contentType string, body string) {
	g.Message.SetBody(contentType, body)
}

// --- Dialer Adapter ---
type gomailDialerAdapter struct {
	*gomail.Dialer
}

func (g *gomailDialerAdapter) DialAndSend(messages ...mail.Message) error {
	rawMessages := make([]*gomail.Message, len(messages))
	for i, msg := range messages {
		adapter, ok := msg.(*gomailMessageAdapter)
		if !ok {
			return fmt.Errorf("invalid message type: expected *gomailMessageAdapter, got %T", msg)
		}
		rawMessages[i] = adapter.Message
	}
	return g.Dialer.DialAndSend(rawMessages...)
}

func New(variable env.Variable) mail.Contract {
	log := logger.New("Go mail config")

	message := gomail.NewMessage()
	message.SetHeader("From", variable.GetSMTPFrom())

	dialer := gomail.NewDialer(
		variable.GetSMTPHost(),
		variable.GetSMTPPort(),
		variable.GetSMTPUsername(),
		variable.GetSMTPPassword())

	dialer.SSL = true

	log.Info("SMTP connection successfully")

	messageAdapter := &gomailMessageAdapter{message}
	dialerAdapter := &gomailDialerAdapter{dialer}

	return mail.New(messageAdapter, dialerAdapter)
}
