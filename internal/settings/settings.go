package settings

import (
	"encoding/hex"
	logger "getfund-api-v2/pkg/log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type ApplicationSettings interface {
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

type applicationSettings struct {
	port                 string
	baseUrl              string
	apiUrl               string
	addrRedis            string
	secretKey            []byte
	serverSalt           []byte
	stripeSecretKey      string
	masterToken          string
	smtp                 *smtpData
	db                   *dbData
	templateDir          string
	timeoutResposenEvent int
}

type smtpData struct {
	host      string
	port      int
	from      string
	userName  string
	passsword string
}

type dbData struct {
	host     string
	port     int
	user     string
	password string
	name     string
}

func (s *applicationSettings) GetPort() string              { return s.port }
func (s *applicationSettings) GetApiUrl() string            { return s.apiUrl }
func (s *applicationSettings) GetBaseUrl() string           { return s.baseUrl }
func (s *applicationSettings) GetServerSalt() []byte        { return s.serverSalt }
func (s *applicationSettings) GetSecretKey() []byte         { return s.secretKey }
func (s *applicationSettings) GetAddrRedis() string         { return s.addrRedis }
func (s *applicationSettings) GetSMTPHost() string          { return s.smtp.host }
func (s *applicationSettings) GetSMTPPort() int             { return s.smtp.port }
func (s *applicationSettings) GetSMTPPassword() string      { return s.smtp.passsword }
func (s *applicationSettings) GetSMTPUsername() string      { return s.smtp.userName }
func (s *applicationSettings) GetSMTPFrom() string          { return s.smtp.from }
func (s *applicationSettings) GetTemplateDir() string       { return s.templateDir }
func (s *applicationSettings) GetTimeoutResponseEvent() int { return s.timeoutResposenEvent }
func (s *applicationSettings) GetDBHost() string            { return s.db.host }
func (s *applicationSettings) GetDBPort() int               { return s.db.port }
func (s *applicationSettings) GetDBUser() string            { return s.db.user }
func (s *applicationSettings) GetDBPassword() string        { return s.db.password }
func (s *applicationSettings) GetDBName() string            { return s.db.name }

func Load() ApplicationSettings {
	logger := logger.New("Settings")
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

	return &applicationSettings{
		port:            ":" + getEnv("PORT", ""),
		baseUrl:         getEnv("BASE_URL", ""),
		apiUrl:          getEnv("BASE_URL", "") + "/api/v2",
		addrRedis:       getEnv("ADDR_REDIS", "Redis address not found"),
		secretKey:       secretKey,
		serverSalt:      serverSalt,
		stripeSecretKey: getEnv("STRIPE_SECRET_KEY", "Stripe key not found"),
		masterToken:     getEnv("MASTER_TOKEN", ""),
		smtp: &smtpData{
			host:      getEnv("SMTP_HOST", ""),
			port:      getIntEnv("SMTP_PORT", ""),
			passsword: getEnv("SMTP_PASSWORD", ""),
			userName:  getEnv("SMTP_USERNAME", ""),
			from:      getEnv("SMTP_FROM", ""),
		},
		db: &dbData{
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
