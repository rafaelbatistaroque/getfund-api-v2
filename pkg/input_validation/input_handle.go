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

// Adiciona o erro na lista de erros existentes se valor é nulo ou vazio
func (i *InputValidation) Required(propetyValue, propertyName string) {
	if isNilOrEmpty(propetyValue) {
		i.erros = append(i.erros, fmt.Sprintf(Err_Msg_PARAMETER_NOT_EMPTY.Error(), propertyName))
	}
}
