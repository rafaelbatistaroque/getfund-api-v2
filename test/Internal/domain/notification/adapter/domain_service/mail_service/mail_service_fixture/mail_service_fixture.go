package mail_service_fixture

import (
	"getfund-api-v2/internal/domain/notification/adapter/domain_service/mail_service"

	"gopkg.in/gomail.v2"
)

type mailServiceFixture struct {
	Mail *gomail.Message
}

func NewSUT() (mail_service.MailService, *mailServiceFixture) {
	mail := gomail.NewMessage()
	dialer := gomail.NewDialer("fake-host", 123, "fake-username", "fake-password")
	dialer.DialAndSend()
	return mail_service.New(mail, dialer),
		&mailServiceFixture{
			Mail: mail,
		}
}

type emailParams struct {
	To          string
	Subject     string
	Content     string
	Attachments []string
}

func defaultEmailParams() *emailParams {
	return &emailParams{
		To:          "fake-To",
		Subject:     "fake-Subject",
		Content:     "fake-content",
		Attachments: []string{"fake-attachment"},
	}
}

type Option func(*emailParams)

func WithoutTo() Option {
	return func(params *emailParams) {
		params.To = ""
	}
}

func WithoutSubject() Option {
	return func(params *emailParams) {
		params.Subject = ""
	}
}

func WithoutContent() Option {
	return func(params *emailParams) {
		params.Content = ""
	}
}

func GetFakeEmailParams(options ...Option) *emailParams {
	params := defaultEmailParams()
	for _, opt := range options {
		opt(params)
	}
	return params
}

func (e *emailParams) GetParams() (string, string, string, []string) {
	return e.To, e.Subject, e.Content, e.Attachments
}
