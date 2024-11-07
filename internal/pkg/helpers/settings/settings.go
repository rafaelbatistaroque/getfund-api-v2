package settings

import (
	"encoding/hex"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type smtpData struct {
	Host      string
	Port      int
	From      string
	UserName  string
	Passsword string
}

type settings struct {
	Port            string
	BaseUrl         string
	SecretKey       []byte
	ServerSalt      []byte
	StripeSecretKey string
	MasterToken     string
	Smtp            *smtpData
}

func Load() *settings {
	env := os.Getenv("GET_FUND_API_ENV")
	if env != "production" {
		if err := godotenv.Load("../../.env.development"); err != nil {
			panic(err.Error())
		}

		fmt.Println(".env file loaded")
	}

	secretKey, err := hex.DecodeString(os.Getenv("SECRET_KEY"))
	if err != nil {
		panic(err.Error())
	}

	serverSalt, err := hex.DecodeString(os.Getenv("SERVER_SALT"))
	if err != nil {
		panic(err.Error())
	}

	return &settings{
		Port:            ":" + getEnv("PORT", ""),
		BaseUrl:         getEnv("BASE_URL", ""),
		SecretKey:       secretKey,
		ServerSalt:      serverSalt,
		StripeSecretKey: getEnv("STRIPE_SECRET_KEY", "Stripe key not found"),
		MasterToken:     getEnv("MASTER_TOKEN", ""),
		Smtp: &smtpData{
			Port:      getIntEnv("SMTP_PORT", ""),
			Host:      getEnv("SMTP_HOST", ""),
			From:      getEnv("SMTP_FROM", ""),
			Passsword: getEnv("SMTP_PASSWORD", ""),
			UserName:  getEnv("SMTP_USERNAME", ""),
		},
	}
}

func getEnv(key string, errorMessagem string) string {
	value := os.Getenv(key)

	if value == "" {
		if errorMessagem != "" {
			panic(errorMessagem)
		}

		panic("Environment variable" + key + " not found.")
	}

	return value
}

func getIntEnv(key string, errorMessagem string) int {
	value, err := strconv.Atoi(os.Getenv(key))

	if err != nil {
		if errorMessagem != "" {
			panic(errorMessagem)
		}

		panic("Environment variable" + key + " not found.")
	}

	return value
}
