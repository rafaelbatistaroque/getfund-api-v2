package mail_service

import (
	inputvalidation "getfund-api-v2/pkg/input_validation"
)

type MailService interface {
	SendMail(to, subject, content string, replyTo []string) error
}

type mailService struct {
	inputvalidation.InputValidation
}

func New() MailService {
	return &mailService{}
}

func (ms *mailService) SendMail(to, subject, content string, replyTo []string) error {
	ms.Required(to, "To")
	if ms.IsInvalid() {
		return ms.GetErrors()
	}

	return nil
}
