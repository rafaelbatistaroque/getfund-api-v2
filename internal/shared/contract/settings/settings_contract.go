package settings

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
}
