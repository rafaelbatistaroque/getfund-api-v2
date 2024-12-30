package mail_service

import (
	inputvalidation "getfund-api-v2/pkg/input_validation"
)

type MailService interface {
	SendMail(to, subject, content string, replyTo []string) error
}

type mailService struct {
	params inputvalidation.InputValidation
}

func New() MailService {
	return &mailService{}
}

func (ms *mailService) SendMail(to, subject, content string, replyTo []string) error {
	ms.params.Required(to, "To")
	ms.params.Required(subject, "Subject")
	ms.params.Required(content, "Content")
	if ms.params.IsInvalid() {
		return ms.params.GetErrors()
	}

	return nil
}
