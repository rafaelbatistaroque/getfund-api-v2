package settingsspy

type ApplicationSettingsSpy struct{}

func (s *ApplicationSettingsSpy) GetPort() string       { return "" }
func (s *ApplicationSettingsSpy) GetApiUrl() string     { return "" }
func (s *ApplicationSettingsSpy) GetAddrRedis() string  { return "localhost:6379" }
func (s *ApplicationSettingsSpy) GetServerSalt() []byte { return []byte("fake-server-salt") }
func (s *ApplicationSettingsSpy) GetSecretKey() []byte  { return []byte("fake-secret-key") }

func New() *ApplicationSettingsSpy {
	return &ApplicationSettingsSpy{}
}
