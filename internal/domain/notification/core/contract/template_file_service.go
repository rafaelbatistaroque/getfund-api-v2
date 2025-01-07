package notification_contract

type TemplateFileService interface {
	GetRecoveryPasswordTemplate() (string, error)
}
