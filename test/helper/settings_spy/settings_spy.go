package settings_spy

type ApplicationSettingsSpy struct{}

func (s *ApplicationSettingsSpy) GetPort() string        { return "" }
func (s *ApplicationSettingsSpy) GetApiUrl() string      { return "" }
func (s *ApplicationSettingsSpy) GetBaseUrl() string     { return "fake-base-url" }
func (s *ApplicationSettingsSpy) GetAddrRedis() string   { return "localhost:6379" }
func (s *ApplicationSettingsSpy) GetServerSalt() []byte  { return []byte("fake-server-salt") }
func (s *ApplicationSettingsSpy) GetSecretKey() []byte   { return []byte("fake-secret-key") }
func (s *ApplicationSettingsSpy) GetSMTPFrom() string    { return "fake-smtp-from" }
func (s *ApplicationSettingsSpy) GetTemplateDir() string { return "fake/template/dir" }

func New() *ApplicationSettingsSpy {
	return &ApplicationSettingsSpy{}
}
