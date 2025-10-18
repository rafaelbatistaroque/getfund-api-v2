package env

import (
	"encoding/hex"
	logger "getfund-api-v2/internal/shared/log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Variable interface {
	GetPort() string
	GetApiUrl() string
	GetBaseUrl() string
	GetAddrRedis() string
	GetServerSalt() []byte
	GetSecretKey() []byte
	GetSMTPHost() string
	GetSMTPPort() int
	GetSMTPPassword() string
	GetSMTPUsername() string
	GetSMTPFrom() string
	GetTemplateDir() string
	GetTimeoutResponseEvent() int
	GetDBHost() string
	GetDBPort() int
	GetDBUser() string
	GetDBPassword() string
	GetDBName() string
}

func Load() Variable {
	logger := logger.New("Environments")
	env := os.Getenv("GET_FUND_API_ENV")
	if env != "production" {
		if err := godotenv.Load(".env.development"); err != nil {
			panic(err.Error())
		}

		logger.Info(".env file loaded")
	}

	secretKey, err := hex.DecodeString(os.Getenv("SECRET_KEY"))
	if err != nil {
		panic(err.Error())
	}

	serverSalt, err := hex.DecodeString(os.Getenv("SERVER_SALT"))
	if err != nil {
		panic(err.Error())
	}

	return &baseVariable{
		port:            ":" + getEnv("PORT", ""),
		baseUrl:         getEnv("BASE_URL", ""),
		apiUrl:          getEnv("BASE_URL", "") + "/api/v2",
		addrRedis:       getEnv("ADDR_REDIS", "Redis address not found"),
		secretKey:       secretKey,
		serverSalt:      serverSalt,
		stripeSecretKey: getEnv("STRIPE_SECRET_KEY", "Stripe key not found"),
		masterToken:     getEnv("MASTER_TOKEN", ""),
		smtp: &smtpVariable{
			host:      getEnv("SMTP_HOST", ""),
			port:      getIntEnv("SMTP_PORT", ""),
			passsword: getEnv("SMTP_PASSWORD", ""),
			userName:  getEnv("SMTP_USERNAME", ""),
			from:      getEnv("SMTP_FROM", ""),
		},
		db: &dbVariable{
			host:     getEnv("DB_HOST", ""),
			port:     getIntEnv("DB_PORT", ""),
			user:     getEnv("DB_USER", ""),
			password: getEnv("DB_PASSWORD", ""),
			name:     getEnv("DB_NAME", ""),
		},
		templateDir:          "internal/domain/notification/adapter/template",
		timeoutResposenEvent: getIntEnv("TIME_OUT_RESPONSE_EVENT", ""),
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
