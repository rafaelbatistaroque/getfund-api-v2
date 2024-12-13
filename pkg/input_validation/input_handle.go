package inputvalidation

import (
	"errors"
	"fmt"
	"strings"
)

type InputValidation struct {
	erros []string
}

// Retorna true se existir erros de validação
func (i *InputValidation) IsInvalid() bool {
	return len(i.erros) > 0
}

// Obtêm lista de erros existentes na validação
func (i *InputValidation) GetErrors() error {
	if len(i.erros) == 0 {
		return nil
	}

	return errors.New(strings.Join(i.erros, "\n"))
}

// Adiciona o erro na lista de erros existentes
func (i *InputValidation) AppendError(propertyName, erroMessagem string) {
	i.erros = append(i.erros, fmt.Sprintf(erroMessagem, propertyName))
}
