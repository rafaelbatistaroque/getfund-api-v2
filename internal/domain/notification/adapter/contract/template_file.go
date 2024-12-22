package template_file

type TemplateFile interface {
	GetRecoveryPasswordTemplate() (string, error)
}
