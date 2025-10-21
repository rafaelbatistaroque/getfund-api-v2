package security_fixture

import (
	"getfund-api-v2/internal/shared/security"
)

func NewSut() security.Hasher {
	return security.NewHasher()
}

func GetSecretKey() []byte {
	return []byte("12345678901234567890123456789012")
}

func GetServerSalt() []byte {
	return []byte("any-salt-key")
}
