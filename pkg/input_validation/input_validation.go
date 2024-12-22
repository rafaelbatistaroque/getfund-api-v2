package inputvalidation

import (
	"strings"
	"unicode"
)

// IsNilOrEmpty verifica se um valor é nulo ou vazio
func isNilOrEmpty[T string | []string](value T) bool {
	switch v := any(value).(type) {
	case string:
		// Trim e verificar se a string está vazia
		return strings.TrimSpace(v) == ""
	case []string:
		// Verificar se o slice está vazio
		return len(v) == 0
	default:
		// Para outros tipos, retornar false
		return false
	}
}

// IsAlphanumeric verifica se um valor é alfanumérico
func IsAlphanumeric(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}
