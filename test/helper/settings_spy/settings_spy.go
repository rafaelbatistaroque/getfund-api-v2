package settings_spy

type ApplicationSettingsSpy struct {
	templateDir string
}

func (s *ApplicationSettingsSpy) GetPort() string        { return "" }
func (s *ApplicationSettingsSpy) GetApiUrl() string      { return "" }
func (s *ApplicationSettingsSpy) GetBaseUrl() string     { return "fake-base-url" }
func (s *ApplicationSettingsSpy) GetAddrRedis() string   { return "localhost:6379" }
func (s *ApplicationSettingsSpy) GetServerSalt() []byte  { return []byte("fake-server-salt") }
func (s *ApplicationSettingsSpy) GetSecretKey() []byte   { return []byte("fake-secret-key") }
func (s *ApplicationSettingsSpy) GetSMTPFrom() string    { return "fake-smtp-from" }
func (s *ApplicationSettingsSpy) GetTemplateDir() string { return s.templateDir }

func New() *ApplicationSettingsSpy {
	return &ApplicationSettingsSpy{}
}

func (s *ApplicationSettingsSpy) SetTemplateDir(templateDir string) {
	s.templateDir = "fake/template/dir"
	if templateDir != "" {
		s.templateDir = templateDir
	}
}
