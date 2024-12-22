package settings

type ApplicationSettings interface {
	GetPort() string
	GetApiUrl() string
	GetBaseUrl() string
	GetAddrRedis() string
	GetServerSalt() []byte
	GetSecretKey() []byte
	GetSMTPFrom() string
	GetTemplateDir() string
}
