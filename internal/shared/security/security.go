package security

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	mathrand "math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
)

const (
	_SIZE_IV              = 16
	_SIZE_SALT            = 16
	_HASH_LENGTH          = 64
	_MERGE_HASHING_LENGTH = _HASH_LENGTH + (_SIZE_SALT * 2)
	_BYTES_LENGTH         = 32
	_ENTRANCE_CODE        = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	_CODE_LENGTH          = 8
)

// Hasher defines the interface for security operations like hashing and encryption.
type Hasher interface {
	// HashWithSalt generates a SHA-256 hash from a given text using a server salt and a salt derived from the input text's length.
	// This method is deterministic for a given input and server salt.
	HashWithSalt(inputText string, serverSalt []byte) (string, error)
	// HashWithSaltLegacy is the old implementation of HashWithSalt. It should only be used for migrating existing data.
	// It uses the string representation of the input length as a salt.
	HashWithSaltLegacy(inputText string, serverSalt []byte) (string, error)
	// IsMatch compares a plain text string against a merged hash (salt + hash) to check for a match.
	IsMatch(inputHashed, inputText string, serverSalt []byte) bool
	// Encrypt encrypts a string using AES-256 with the provided secret key.
	// It returns a merged hex string containing the IV and the encrypted data.
	Encrypt(input string, secretKey []byte) string
	// DecryptMerged decrypts a merged hex string (IV + data) using AES-256 and the secret key.
	DecryptMerged(mergedEncryptedData string, secretKey []byte) string
	// HashAndMerge generates a SHA-256 hash with a unique, random salt and merges them into a single hex string.
	HashAndMerge(input string, serverSalt []byte) string
	// GetRandomCode generates a random code of a given length from the ENTRANCE_CODE character set.
	GetRandomCode(length int) (string, error)
	// Hash generates a SHA-256 hash with a unique, random salt, returning the hash and salt as separate fields in a Hashing struct.
	Hash(inputText string, serverSalt []byte) (*Hashing, error)
}

type hasher struct {
}

// NewHasher creates and returns a new instance of the default Hasher implementation.
func NewHasher() Hasher {
	return &hasher{}
}

// encryption holds the Initialization Vector (IV) and encrypted data, both as hex strings.
type encryption struct {
	IV   string
	Data string
}

// Hashing holds the salt and the resulting hash data, both as hex strings.
type Hashing struct {
	Salt string
	Data string
}

// Gera um hash com um salt específico
func (s *hasher) HashWithSalt(inputText string, serverSalt []byte) (string, error) {
	salt := make([]byte, _CODE_LENGTH)
	binary.BigEndian.PutUint64(salt, uint64(len(inputText)))

	inputData := []byte(inputText)
	hash := sha256.New()
	hash.Write(append(append(serverSalt, inputData...), salt...))

	return hex.EncodeToString(hash.Sum(nil)), nil
}

// Gera um hash com um salt específico (LÓGICA ANTIGA/LEGADA)
func (s *hasher) HashWithSaltLegacy(inputText string, serverSalt []byte) (string, error) {
	// ATENÇÃO: Esta é a implementação antiga, mantida apenas para migração.
	salt, err := hex.DecodeString(fmt.Sprint(len(inputText)))
	if err != nil {
		return "", err
	}

	inputData := []byte(inputText)
	hash := sha256.New()
	hash.Write(append(append(serverSalt, inputData...), salt...))
	return hex.EncodeToString(hash.Sum(nil)), nil
}

// Gera um Hash com salt
func (s *hasher) Hash(inputText string, serverSalt []byte) (*Hashing, error) {
	entrySalt := make([]byte, _SIZE_SALT)
	if _, err := io.ReadFull(rand.Reader, entrySalt); err != nil {
		return nil, err
	}

	inputData := []byte(inputText)
	hash := sha256.New()
	hash.Write(append(append(serverSalt, inputData...), entrySalt...))
	hashedData := hash.Sum(nil)

	return &Hashing{
		Salt: hex.EncodeToString(entrySalt),
		Data: hex.EncodeToString(hashedData),
	}, nil
}

// Função principal para encriptar e mesclar sem hashing adicional
func (s *hasher) Encrypt(input string, secretKey []byte) string {
	if len(secretKey) != _BYTES_LENGTH {
		panic("The secret key must have 32 bytes to AES-256")
	}

	// Encriptar o texto de entrada
	encryption, err := encrypt(input, secretKey)
	if err != nil {
		panic(err)
	}

	// Mesclar o IV e os dados encriptados, sem hash
	return mergeEncryption(encryption)
}

// Mescla o salt e os dados do hash
func (s *hasher) HashAndMerge(input string, serverSalt []byte) string {
	hashing, err := s.Hash(input, serverSalt)
	if err != nil {
		panic(err)
	}

	return mergeHashing(hashing)
}

// Mescla o salt e os dados do hash
func mergeHashing(hashing *Hashing) string {
	return hashing.Salt + hashing.Data
}

// Função que combina a extração do IV e dos dados e realiza a decriptação
func (s *hasher) DecryptMerged(mergedEncryptedData string, secretKey []byte) string {
	ivHex := extractEncryptionIV(mergedEncryptedData)
	dataHex := extractEncryptionData(mergedEncryptedData)

	// Cria a estrutura Encryption com o IV e os dados extraídos
	encryption := &encryption{
		IV:   ivHex,
		Data: dataHex,
	}

	// Desencripta os dados
	return decrypt(encryption, secretKey)
}

// // Função para extrair o IV do dado encriptado mesclado
func extractEncryptionIV(encryptedData string) string {
	return encryptedData[:_SIZE_IV*2] // Extrai o IV em hexadecimal
}

// Função para extrair os dados encriptados do dado encriptado mesclado
func extractEncryptionData(encryptedData string) string {
	return encryptedData[_SIZE_IV*2:] // Extrai os dados encriptados em hexadecimal
}

// Testa se o hash mesclado corresponde ao hash de entrada
func (h *hasher) IsMatch(inputHashed, inputText string, serverSalt []byte) bool {
	// Certifique-se de que o hash mesclado tem o comprimento mínimo necessário
	if len(inputHashed) < _MERGE_HASHING_LENGTH {
		return false
	}

	entrySalt := extractHashingSalt(inputHashed)
	hashedData := extractHashingData(inputHashed)

	// Certifique-se de que o salt extraído tem o comprimento esperado
	if len(entrySalt) != _SIZE_SALT {
		return false
	}

	inputData := []byte(inputText)

	hash := sha256.New()
	hash.Write(append(append(serverSalt, inputData...), entrySalt...))
	return bytes.Equal(hashedData, hash.Sum(nil))
}

// Extrai o salt de um hash mesclado
func extractHashingSalt(mergedHashing string) []byte {
	hexSalt := mergedHashing[:_SIZE_SALT*2]
	salt, _ := hex.DecodeString(hexSalt)
	return salt
}

// Extrai os dados de um hash mesclado
func extractHashingData(mergedHashing string) []byte {
	hexData := mergedHashing[_SIZE_SALT*2:]
	data, _ := hex.DecodeString(hexData)
	return data
}

func encrypt(inputText string, secretKey []byte) (*encryption, error) {
	// Gerar IV aleatório
	iv := make([]byte, _SIZE_IV)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, fmt.Errorf("erro ao gerar IV: %v", err)
	}

	// Criar o cipher AES
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar cipher AES: %v", err)
	}

	// Configurar o encriptador CBC com o IV
	encrypter := cipher.NewCBCEncrypter(block, iv)

	// Converter o texto de entrada para bytes e encriptar
	plainText := []byte(inputText)
	paddedText := pad(plainText, aes.BlockSize) // Realizar o padding

	encrypted := make([]byte, len(paddedText))
	encrypter.CryptBlocks(encrypted, paddedText)

	// Retornar o IV e os dados encriptados como hexadecimal
	return &encryption{
		IV:   hex.EncodeToString(iv),
		Data: hex.EncodeToString(encrypted),
	}, nil
}

func decrypt(encryption *encryption, secretKey []byte) string {
	iv, err := hex.DecodeString(encryption.IV)
	if err != nil {
		panic(err)
	}

	encryptedData, err := hex.DecodeString(encryption.Data)
	if err != nil {
		panic(err)
	}

	block, err := aes.NewCipher(secretKey)
	if err != nil {
		panic(err)
	}

	decrypter := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(encryptedData))
	decrypter.CryptBlocks(decrypted, encryptedData)

	unpaddedText, err := unpad(decrypted, aes.BlockSize)
	if err != nil {
		panic(err)
	}

	return string(unpaddedText)
}

func mergeEncryption(encryption *encryption) string {
	return encryption.IV + encryption.Data
}

// Função para gerar um UUID (GUID)
// GetGUID generates a new Version 4 UUID (Universally Unique Identifier).
func GetGUID() string {
	return uuid.New().String()
}

// Função de padding e unpadding
func pad(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padText...)
}

func unpad(src []byte, blockSize int) ([]byte, error) {
	length := len(src)
	padding := int(src[length-1])
	if padding > blockSize || padding == 0 {
		return nil, fmt.Errorf("invalid padding")
	}
	return src[:length-padding], nil
}

func (h *hasher) GetRandomCode(length int) (string, error) {
	if length <= 0 {
		return "", fmt.Errorf("length must be greater than zero")
	}
	var code strings.Builder
	mathrand.NewSource(time.Now().UnixNano())

	for range length {
		randomIndex := mathrand.Intn(len(_ENTRANCE_CODE))
		code.WriteByte(_ENTRANCE_CODE[randomIndex])
	}

	return code.String(), nil
}
