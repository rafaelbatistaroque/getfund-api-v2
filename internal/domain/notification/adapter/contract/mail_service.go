package notification_contract

type MailService interface {
	SendMail(to, subject, content string, replyTo []string) error
}
