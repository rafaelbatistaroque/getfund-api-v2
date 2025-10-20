package mail_fixture

import (
	"errors"
	"getfund-api-v2/internal/shared/mail"
)

// messageSpy is a spy implementation of the mail.Message interface
type messageSpy struct {
	Headers    map[string][]string
	Body       string
	CallsCount map[string]int
}

func (m *messageSpy) SetHeader(key string, value ...string) {
	m.Headers[key] = value
	m.CallsCount["SetHeader"]++
}

func (m *messageSpy) SetBody(contentType string, body string) {
	m.Body = body
	m.CallsCount["SetBody"]++
}

func (m *messageSpy) GetHeader(key string) []string {
	return m.Headers[key]
}

// dialerSpy is a spy implementation of the mail.Dialer interface
type dialerSpy struct {
	CallsCount map[string]int
	Error      error
	Messages   []mail.Message
}

type MailFixture struct {
	DialerSpy  *dialerSpy
	MessageSpy *messageSpy
}

func NewSut() (mail.Contract, *MailFixture) {
	dialerSpy := &dialerSpy{
		CallsCount: make(map[string]int),
	}
	messageSpy := &messageSpy{
		Headers:    make(map[string][]string),
		CallsCount: make(map[string]int),
	}
	sut := mail.New(messageSpy, dialerSpy)
	return sut, &MailFixture{
		DialerSpy:  dialerSpy,
		MessageSpy: messageSpy,
	}
}

func (s *dialerSpy) DialAndSend(m ...mail.Message) error {
	s.CallsCount["DialAndSend"]++
	s.Messages = m
	return s.Error
}

func (s *dialerSpy) DefineError() {
	s.Error = errors.New("fake-error")
}

func (s *dialerSpy) DefineSuccess() {
	s.Error = nil
}

type EmailParams struct {
	To      string
	Subject string
	Content string
	ReplyTo []string
}

func defaultEmailParams() *EmailParams {
	return &EmailParams{
		To:      "fake-to@example.com",
		Subject: "fake-Subject",
		Content: "<h1>fake-content</h1>",
		ReplyTo: []string{"fake-reply-to@example.com"},
	}
}

type Option func(*EmailParams)

func WithoutTo() Option {
	return func(params *EmailParams) {
		params.To = ""
	}
}

func WithoutSubject() Option {
	return func(params *EmailParams) {
		params.Subject = ""
	}
}

func WithoutContent() Option {
	return func(params *EmailParams) {
		params.Content = ""
	}
}

func GetFakeEmailParams(options ...Option) *EmailParams {
	params := defaultEmailParams()
	for _, opt := range options {
		opt(params)
	}
	return params
}

func (e *EmailParams) GetParams() (string, string, string, []string) {
	return e.To, e.Subject, e.Content, e.ReplyTo
}
