package verify

import (
	"reflect"
	"testing"
)

// Verifier encapsula o valor e o estado do teste
type Verifier struct {
	t       *testing.T
	value   interface{}
	message string
}

// Should inicializa o Verifier
func Should(t *testing.T, value interface{}) *Verifier {
	return &Verifier{t: t, value: value}
}

// Message define uma mensagem personalizada
func (v *Verifier) Message(msg string) *Verifier {
	v.message = msg
	return v
}

// Be verifica igualdade estrita
func (v *Verifier) Be(expected interface{}) *Verifier {
	if !reflect.DeepEqual(v.value, expected) {
		message := "Expected values to have the same properties and values"
		if v.message != "" {
			message = v.message
		}
		v.t.Errorf("%s: expected %v (type %T), but got %v (type %T)", message, expected, expected, v.value, v.value)
	}
	return v
}

// StrictEqual verifica se dois valores têm as mesmas propriedades e valores
func (v *Verifier) StrictEqual(expected interface{}) *Verifier {
	if !reflect.DeepEqual(v.value, expected) {
		message := "Expected values to have the same properties and values"
		if v.message != "" {
			message = v.message
		}
		v.t.Errorf("%s: expected %v (type %T), but got %v (type %T)", message, expected, expected, v.value, v.value)
	}
	return v
}

// NotEqual verifica desigualdade
func (v *Verifier) NotEqual(unexpected interface{}) *Verifier {
	if reflect.DeepEqual(v.value, unexpected) {
		message := "Expected values to be different"
		if v.message != "" {
			message = v.message
		}
		v.t.Errorf("%s: expected not %v (type %T), but got %v (type %T)", message, unexpected, unexpected, v.value, v.value)
	}
	return v
}

// NotEmpty verifica se uma string não está vazia
func (v *Verifier) NotEmpty() *Verifier {
	if str, ok := v.value.(string); ok && str == "" {
		message := "Expected string to be not empty"
		if v.message != "" {
			message = v.message
		}
		v.t.Errorf("%s: string is empty", message)
	}
	return v
}

// Empty verifica se uma string está vazia
func (v *Verifier) Empty() *Verifier {
	if str, ok := v.value.(string); ok && str != "" {
		message := "Expected string to be empty"
		if v.message != "" {
			message = v.message
		}
		v.t.Errorf("%s: string is not empty", message)
	}
	return v
}

// Nil verifica se o valor é nil
func (v *Verifier) Nil() *Verifier {
	if !isNil(v.value) {
		message := "Expected value to be nil"
		if v.message != "" {
			message = v.message
		}
		v.t.Errorf("%s: expected nil, but got %v (type %T)", message, v.value, v.value)
	}
	return v
}

// NotNil verifica se o valor não é nil
func (v *Verifier) NotNil() *Verifier {
	if isNil(v.value) {
		message := "Expected value to not be nil"
		if v.message != "" {
			message = v.message
		}
		v.t.Errorf("%s: expected not nil, but got nil", message)
	}
	return v
}

// isNil verifica se um valor é nil
func isNil(value interface{}) bool {
	if value == nil {
		return true
	}
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return v.IsNil()
	}
	return false
}

// BeTrue verifica se o valor booleano é true
func (v *Verifier) BeTrue() *Verifier {
	if value, ok := v.value.(bool); !ok || !value {
		message := "Expected value to be true"
		if v.message != "" {
			message = v.message
		}
		v.t.Errorf("%s: expected true, but got %v (type %T)", message, v.value, v.value)
	}
	return v
}

// BeFalse verifica se o valor booleano é false
func (v *Verifier) BeFalse() *Verifier {
	if value, ok := v.value.(bool); !ok || value {
		message := "Expected value to be false"
		if v.message != "" {
			message = v.message
		}
		v.t.Errorf("%s: expected false, but got %v (type %T)", message, v.value, v.value)
	}
	return v
}

// Len verifica o comprimento de um slice, mapa, canal, string ou array
func (v *Verifier) Len(expected int) *Verifier {
	// Usando reflection para verificar o comprimento
	val := reflect.ValueOf(v.value)
	if val.Kind() != reflect.Slice && val.Kind() != reflect.Map && val.Kind() != reflect.Chan && val.Kind() != reflect.Array && val.Kind() != reflect.String {
		v.t.Errorf("Expected a slice, map, channel, array, or string, but got %T", v.value)
		return v
	}

	// Comparando o comprimento
	actual := val.Len()
	if actual != expected {
		message := "Expected length to be equal"
		if v.message != "" {
			message = v.message
		}
		v.t.Errorf("%s: expected length %d, but got %d", message, expected, actual)
	}
	return v
}
