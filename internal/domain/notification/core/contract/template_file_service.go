package notification_contract

type TemplateFileContract interface {
	GetRecoveryPasswordTemplate() (string, error)
	GetActivationAccountTemplate() (string, error)
}
