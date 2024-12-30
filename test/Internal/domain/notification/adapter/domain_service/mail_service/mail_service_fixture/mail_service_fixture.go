package mail_service_fixture

import "getfund-api-v2/internal/domain/notification/adapter/domain_service/mail_service"

func NewSUT() mail_service.MailService {
	return mail_service.New()
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
		Attachments: nil,
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

func GetFakeEmailParams(options ...Option) (string, string, string, []string) {
	params := defaultEmailParams()
	for _, opt := range options {
		opt(params)
	}
	return params.To, params.Subject, params.Content, params.Attachments
}
