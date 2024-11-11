package settings

import (
	"encoding/hex"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type smtpData struct {
	host      string
	port      int
	from      string
	userName  string
	passsword string
}

type applicationSettings struct {
	port            string
	baseUrl         string
	apiUrl          string
	secretKey       []byte
	serverSalt      []byte
	stripeSecretKey string
	masterToken     string
	smtp            *smtpData
}

type ApplicationSettings interface {
}

func (s *applicationSettings) GetPort() string   { return s.port }
func (s *applicationSettings) GetApiUrl() string { return s.apiUrl }

func Load() *applicationSettings {
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

	return &applicationSettings{
		port:            ":" + getEnv("PORT", ""),
		baseUrl:         getEnv("BASE_URL", ""),
		apiUrl:          getEnv("BASE_URL", "") + "/api/v2",
		secretKey:       secretKey,
		serverSalt:      serverSalt,
		stripeSecretKey: getEnv("STRIPE_SECRET_KEY", "Stripe key not found"),
		masterToken:     getEnv("MASTER_TOKEN", ""),
		smtp: &smtpData{
			port:      getIntEnv("SMTP_PORT", ""),
			host:      getEnv("SMTP_HOST", ""),
			from:      getEnv("SMTP_FROM", ""),
			passsword: getEnv("SMTP_PASSWORD", ""),
			userName:  getEnv("SMTP_USERNAME", ""),
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
