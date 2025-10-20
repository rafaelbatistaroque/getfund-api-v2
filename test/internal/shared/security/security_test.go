package security_test

import (
	"testing"

	fixture "getfund-api-v2/test/internal/shared/security/fixture"

	"github.com/rafaelbatistaroque/verify/v2"
)

func Test_GivenGetRandomCode_WhenLengthIsZero_ThenEnsureReturnError(t *testing.T) {
	// Arrange
	sut := fixture.NewSut()

	// Act
	_, err := sut.GetRandomCode(0)

	// Assert
	verify.Should(t, err).NotNil()
}

func Test_GivenGetRandomCode_WhenValidLength_ThenEnsureReturnCode(t *testing.T) {
	// Arrange
	sut := fixture.NewSut()

	// Act
	code, err := sut.GetRandomCode(8)

	// Assert
	verify.Should(t, err).Nil()
	verify.Should(t, len(code)).Be(8)
}

func Test_GivenEncrypt_WhenSecretKeyIsNot32Bytes_ThenEnsurePanic(t *testing.T) {
	// Arrange
	sut := fixture.NewSut()

	// Act & Assert
	verify.Should(t, func() {
		sut.Encrypt("any_text", []byte("invalid_key"))
	}).Message("The secret key must have 32 bytes to AES-256").Panic()
}

func Test_GivenEncrypt_WhenValidInput_ThenEnsureReturnEncryptedText(t *testing.T) {
	// Arrange
	sut := fixture.NewSut()
	text := "any_text"

	// Act
	encryptedText := sut.Encrypt(text, fixture.GetSecretKey())

	// Assert
	verify.Should(t, encryptedText).NotEqual(text)
}

func Test_GivenDecryptMerged_WhenValidInput_ThenEnsureReturnDecryptedText(t *testing.T) {
	// Arrange
	sut := fixture.NewSut()
	text := "any_text"
	encryptedText := sut.Encrypt(text, fixture.GetSecretKey())

	// Act
	decryptedText := sut.DecryptMerged(encryptedText, fixture.GetSecretKey())

	// Assert
	verify.Should(t, decryptedText).Be(text)
}

func Test_GivenHashAndMerge_WhenValidInput_ThenEnsureReturnHashedText(t *testing.T) {
	// Arrange
	sut := fixture.NewSut()
	text := "any_text"

	// Act
	hashedText := sut.HashAndMerge(text, fixture.GetServerSalt())

	// Assert
	verify.Should(t, hashedText).NotEqual(text)
}

func Test_GivenIsMatch_WhenValidInput_ThenEnsureReturnTrue(t *testing.T) {
	// Arrange
	sut := fixture.NewSut()
	text := "any_text"
	hashedText := sut.HashAndMerge(text, fixture.GetServerSalt())

	// Act
	isMatch := sut.IsMatch(hashedText, text, fixture.GetServerSalt())

	// Assert
	verify.Should(t, isMatch).BeTrue()
}

func Test_GivenIsMatch_WhenInvalidInput_ThenEnsureReturnFalse(t *testing.T) {
	// Arrange
	sut := fixture.NewSut()
	text := "any_text"
	hashedText := sut.HashAndMerge(text, fixture.GetServerSalt())

	// Act
	isMatch := sut.IsMatch(hashedText, "invalid_text", fixture.GetServerSalt())

	// Assert
	verify.Should(t, isMatch).BeFalse()
}

func Test_GivenHashWithSalt_WhenCalled_ThenEnsureReturnsAHash(t *testing.T) {
	// Arrange
	sut := fixture.NewSut()
	text := "any_text"

	// Act
	hash, err := sut.HashWithSalt(text, fixture.GetServerSalt())

	// Assert
	verify.Should(t, err).Nil()
	verify.Should(t, hash).NotEqual(text)
}

func Test_GivenHash_WhenCalled_ThenEnsureReturnsAHashAndSalt(t *testing.T) {
	// Arrange
	sut := fixture.NewSut()
	text := "any_text"

	// Act
	hashing, err := sut.Hash(text, fixture.GetServerSalt())

	// Assert
	verify.Should(t, err).Nil()
	verify.Should(t, hashing.Data).NotEqual(text)
	verify.Should(t, hashing.Salt).NotNil()
}
