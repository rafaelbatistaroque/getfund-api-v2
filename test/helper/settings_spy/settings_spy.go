package settings_spy

type ApplicationSettingsSpy struct {
	templateDir          string
	timeoutResposenEvent int
}

func (s *ApplicationSettingsSpy) GetPort() string              { return "" }
func (s *ApplicationSettingsSpy) GetApiUrl() string            { return "" }
func (s *ApplicationSettingsSpy) GetBaseUrl() string           { return "fake-base-url" }
func (s *ApplicationSettingsSpy) GetAddrRedis() string         { return "localhost:6379" }
func (s *ApplicationSettingsSpy) GetServerSalt() []byte        { return []byte("fake-server-salt") }
func (s *ApplicationSettingsSpy) GetSecretKey() []byte         { return []byte("fake-secret-key") }
func (s *ApplicationSettingsSpy) GetSMTPHost() string          { return "fake-host" }
func (s *ApplicationSettingsSpy) GetSMTPPort() int             { return 123 }
func (s *ApplicationSettingsSpy) GetSMTPPassword() string      { return "fake-password" }
func (s *ApplicationSettingsSpy) GetSMTPUsername() string      { return "fake-username" }
func (s *ApplicationSettingsSpy) GetSMTPFrom() string          { return "fake-smtp-from" }
func (s *ApplicationSettingsSpy) GetTemplateDir() string       { return s.templateDir }
func (s *ApplicationSettingsSpy) GetTimeoutResponseEvent() int { return s.timeoutResposenEvent }

func New() *ApplicationSettingsSpy {
	return &ApplicationSettingsSpy{}
}

func (s *ApplicationSettingsSpy) SetTemplateDir(templateDir string) {
	s.templateDir = "fake/template/dir"
	if templateDir != "" {
		s.templateDir = templateDir
	}
}

func (s *ApplicationSettingsSpy) SetTimeoutResponseEvent(timeout int) {
	if timeout == 0 {
		s.timeoutResposenEvent = 1
		return
	}

	s.timeoutResposenEvent = timeout
}
