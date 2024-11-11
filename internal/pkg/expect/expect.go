package expect

import (
	"reflect"
	"testing"
)

// StrictEqual verifica se dois valores são a mesma instância
func StrictEqual(t *testing.T, result, shouldBe interface{}, msg ...string) {
	if result != shouldBe {
		message := "Expected values to be the same instance"
		if len(msg) > 0 {
			message = msg[0]
		}
		t.Errorf("%s: expected %v (type %T), but got %v (type %T)", message, shouldBe, shouldBe, result, result)
	}
}

func Equal(t *testing.T, result, shouldBe interface{}, msg ...string) {
	if result != shouldBe {
		message := "Expected values to be equal"
		if len(msg) > 0 {
			message = msg[0]
		}
		t.Errorf("%s: expected %v, but got %v", message, shouldBe, result)
	}
}

// ExpectNotEqual verifica se dois valores não são iguais
func NotEqual(t *testing.T, result, shouldNotBe interface{}, msg ...string) {
	if reflect.DeepEqual(result, shouldNotBe) {
		message := "Expected values to be different"
		if len(msg) > 0 {
			message = msg[0]
		}
		t.Errorf("%s: expected different from %v (type %T), but got %v (type %T)", message, shouldNotBe, shouldNotBe, result, result)
	}
}

// ExpectNil verifica se um valor é nil
func Nil(t *testing.T, value interface{}, msg ...string) {
	if !isNil(value) {
		message := "Expected value to be nil"
		if len(msg) > 0 {
			message = msg[0]
		}
		t.Errorf("%s: expected nil, but got %v (type %T)", message, value, value)
	}
}

// ExpectNotNil verifica se um valor não é nil
func NotNil(t *testing.T, value interface{}, msg ...string) {
	if isNil(value) {
		message := "Expected value to not be nil"
		if len(msg) > 0 {
			message = msg[0]
		}
		t.Errorf("%s: expected not nil, but got nil", message)
	}
}

// ExpectLen verifica o comprimento de um slice, mapa, canal ou string
func Len(t *testing.T, value interface{}, length int, msg ...string) {
	v := reflect.ValueOf(value)
	if v.Len() != length {
		message := "Expected length to be equal"
		if len(msg) > 0 {
			message = msg[0]
		}
		t.Errorf("%s: expected length %d, but got %d", message, length, v.Len())
	}
}

// ExpectTrue verifica se um valor é true
func True(t *testing.T, value bool, msg ...string) {
	if !value {
		message := "Expected value to be true"
		if len(msg) > 0 {
			message = msg[0]
		}
		t.Errorf("%s: expected true, but got false", message)
	}
}

// ExpectFalse verifica se um valor é false
func False(t *testing.T, value bool, msg ...string) {
	if value {
		message := "Expected value to be false"
		if len(msg) > 0 {
			message = msg[0]
		}
		t.Errorf("%s: expected false, but got true", message)
	}
}

// ExpectPanic verifica se uma função causa um pânico
func Panic(t *testing.T, fn func(), msg ...string) {
	defer func() {
		if r := recover(); r == nil {
			message := "Expected function to panic"
			if len(msg) > 0 {
				message = msg[0]
			}
			t.Errorf("%s: expected panic, but function did not panic", message)
		}
	}()
	fn()
}

// ExpectNotPanic verifica se uma função não causa um pânico
func NotPanic(t *testing.T, fn func(), msg ...string) {
	defer func() {
		if r := recover(); r != nil {
			message := "Expected function to not panic"
			if len(msg) > 0 {
				message = msg[0]
			}
			t.Errorf("%s: expected no panic, but function panicked with %v", message, r)
		}
	}()
	fn()
}

// ExpectError verifica se uma função retorna um erro
func Error(t *testing.T, err error, msg ...string) {
	if err == nil {
		message := "Expected an error"
		if len(msg) > 0 {
			message = msg[0]
		}
		t.Errorf("%s: expected error, but got nil", message)
	}
}

// ExpectNoError verifica se uma função não retorna um erro
func NoError(t *testing.T, err error, msg ...string) {
	if err != nil {
		message := "Expected no error"
		if len(msg) > 0 {
			message = msg[0]
		}
		t.Errorf("%s: expected no error, but got %v", message, err)
	}
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

// Contains verifica se o item esperado está presente em actual
func Contains(t *testing.T, actual []string, expected string, msg ...string) {
	actualMap := make(map[string]bool)
	for _, item := range actual {
		actualMap[item] = true
	}

	if !actualMap[expected] {
		message := "Expected item to be present"
		if len(msg) > 0 {
			message = msg[0]
		}
		t.Errorf("%s: missing item %v", message, expected)
	}
}

// ContainsAll verifica se todos os itens de expected estão presentes em actual
func ContainsAll(t *testing.T, actual, expected []string, msg ...string) {
	actualMap := make(map[string]bool)
	for _, item := range actual {
		actualMap[item] = true
	}

	var missingItems []string
	for _, item := range expected {
		if !actualMap[item] {
			missingItems = append(missingItems, item)
		}
	}

	if len(missingItems) > 0 {
		message := "Expected all items to be present"
		if len(msg) > 0 {
			message = msg[0]
		}
		t.Errorf("%s: missing items %v", message, missingItems)
	}
}
