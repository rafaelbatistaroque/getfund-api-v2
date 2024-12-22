package mail_service

type MailService interface {
	SendMail(from, to, subject, content string, replyTo []string) error
}

type mailService struct {
}

func New() MailService {
	return &mailService{}
}

func (ms *mailService) SendMail(from, to, subject, content string, replyTo []string) error {
	return nil
}
