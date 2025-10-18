package mail_spy

import "errors"

type MailServiceSpy struct {
	Params     map[string]any
	CallsCount map[string]int

	SuccessResult map[string]any
	ErrorResult   map[string]error
}

func New() *MailServiceSpy {
	return &MailServiceSpy{
		Params:        make(map[string]any),
		CallsCount:    make(map[string]int),
		SuccessResult: make(map[string]any),
		ErrorResult:   make(map[string]error),
	}
}

func (s *MailServiceSpy) SendMail(to, subject, content string, replyTo []string) error {
	s.Params["SendMail:to"] = to
	s.Params["SendMail:subject"] = subject
	s.Params["SendMail:content"] = content
	s.Params["SendMail:replyTo"] = replyTo

	s.CallsCount["SendMail"]++

	return s.ErrorResult["SendMail"]
}

func (s *MailServiceSpy) DefineSendMailError() {
	s.ErrorResult["SendMail"] = errors.New("fake-error")
}
