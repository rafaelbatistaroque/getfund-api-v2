package security

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"

	"github.com/google/uuid"
)

const (
	SIZE_IV              = 16
	SIZE_SALT            = 16
	ENCRYPTION_ALGORITHM = "aes-256-cbc"
	HASH_LENGTH          = 64
	MERGE_HASHING_LENGTH = HASH_LENGTH + (SIZE_SALT * 2)
	GUID_LENGTH          = 36
)

type Encryption struct {
	IV   string
	Data string
}

type Hashing struct {
	Salt string
	Data string
}

// Gera um hash com salt
func hash(inputText string, serverSalt []byte) (*Hashing, error) {
	entrySalt := make([]byte, SIZE_SALT)
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
func EncryptAndMerge(input string, secretKey []byte) string {
	if len(secretKey) != 32 {
		panic("a chave secreta deve ter 32 bytes para AES-256")
	}

	// Encriptar o texto de entrada
	encryption, err := encrypt(input, secretKey)
	if err != nil {
		panic(err)
	}

	// Mesclar o IV e os dados encriptados, sem hash
	return mergeEncryption(encryption)
}

// Gera um hash com um salt específico
func HashWithSalt(inputText, saltText string, serverSalt []byte) (string, error) {
	salt, err := hex.DecodeString(saltText)
	if err != nil {
		return "", err
	}

	inputData := []byte(inputText)
	hash := sha256.New()
	hash.Write(append(append(serverSalt, inputData...), salt...))

	return hex.EncodeToString(hash.Sum(nil)), nil
}

// Mescla o salt e os dados do hash
func HashAndMerge(input string, serverSalt []byte) string {
	hashing, err := hash(input, serverSalt)
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
func DecryptMerged(mergedEncryptedData string, secretKey []byte) string {
	ivHex := extractEncryptionIV(mergedEncryptedData)
	dataHex := extractEncryptionData(mergedEncryptedData)

	// Cria a estrutura Encryption com o IV e os dados extraídos
	encryption := &Encryption{
		IV:   ivHex,
		Data: dataHex,
	}

	// Desencripta os dados
	return decrypt(encryption, secretKey)
}

// // Função para extrair o IV do dado encriptado mesclado
func extractEncryptionIV(encryptedData string) string {
	return encryptedData[:SIZE_IV*2] // Extrai o IV em hexadecimal
}

// Função para extrair os dados encriptados do dado encriptado mesclado
func extractEncryptionData(encryptedData string) string {
	return encryptedData[SIZE_IV*2:] // Extrai os dados encriptados em hexadecimal
}

// Testa se o hash mesclado corresponde ao hash de entrada
func TestMergedHashing(mergedHashing, inputText string, serverSalt []byte) bool {
	entrySalt := extractHashingSalt(mergedHashing)
	hashedData := extractHashingData(mergedHashing)
	inputData := []byte(inputText)

	hash := sha256.New()
	hash.Write(append(append(serverSalt, inputData...), entrySalt...))
	return bytes.Equal(hashedData, hash.Sum(nil))
}

// Extrai o salt de um hash mesclado
func extractHashingSalt(mergedHashing string) []byte {
	hexSalt := mergedHashing[:SIZE_SALT*2]
	salt, _ := hex.DecodeString(hexSalt)
	return salt
}

// Extrai os dados de um hash mesclado
func extractHashingData(mergedHashing string) []byte {
	hexData := mergedHashing[SIZE_SALT*2:]
	data, _ := hex.DecodeString(hexData)
	return data
}

func encrypt(inputText string, secretKey []byte) (*Encryption, error) {
	// Gerar IV aleatório
	iv := make([]byte, SIZE_IV)
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
	return &Encryption{
		IV:   hex.EncodeToString(iv),
		Data: hex.EncodeToString(encrypted),
	}, nil
}

func decrypt(encryption *Encryption, secretKey []byte) string {
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

func mergeEncryption(encryption *Encryption) string {
	return encryption.IV + encryption.Data
}

// Função para gerar um UUID (GUID)
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
