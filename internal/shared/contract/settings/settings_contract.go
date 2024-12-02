package settings

type ApplicationSettings interface {
	GetPort() string
	GetApiUrl() string
	GetAddrRedis() string
	GetServerSalt() []byte
	GetSecretKey() []byte
}
