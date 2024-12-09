package codeservice

type CodeService interface {
	GetRandomCode(length int) (string, error)
}
