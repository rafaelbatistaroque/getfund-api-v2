package notification_contract

type TemplateFile interface {
	GetRecoveryPasswordTemplate() (string, error)
}
