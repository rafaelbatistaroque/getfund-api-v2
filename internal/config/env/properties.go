package env

type baseVariable struct {
	port                 string
	baseUrl              string
	apiUrl               string
	addrRedis            string
	secretKey            []byte
	serverSalt           []byte
	stripeSecretKey      string
	masterToken          string
	smtp                 *smtpVariable
	db                   *dbVariable
	templateDir          string
	timeoutResposenEvent int
}

type smtpVariable struct {
	host      string
	port      int
	from      string
	userName  string
	passsword string
}

type dbVariable struct {
	host     string
	port     int
	user     string
	password string
	name     string
}

func (s *baseVariable) GetPort() string              { return s.port }
func (s *baseVariable) GetApiUrl() string            { return s.apiUrl }
func (s *baseVariable) GetBaseUrl() string           { return s.baseUrl }
func (s *baseVariable) GetServerSalt() []byte        { return s.serverSalt }
func (s *baseVariable) GetSecretKey() []byte         { return s.secretKey }
func (s *baseVariable) GetAddrRedis() string         { return s.addrRedis }
func (s *baseVariable) GetSMTPHost() string          { return s.smtp.host }
func (s *baseVariable) GetSMTPPort() int             { return s.smtp.port }
func (s *baseVariable) GetSMTPPassword() string      { return s.smtp.passsword }
func (s *baseVariable) GetSMTPUsername() string      { return s.smtp.userName }
func (s *baseVariable) GetSMTPFrom() string          { return s.smtp.from }
func (s *baseVariable) GetTemplateDir() string       { return s.templateDir }
func (s *baseVariable) GetTimeoutResponseEvent() int { return s.timeoutResposenEvent }
func (s *baseVariable) GetDBHost() string            { return s.db.host }
func (s *baseVariable) GetDBPort() int               { return s.db.port }
func (s *baseVariable) GetDBUser() string            { return s.db.user }
func (s *baseVariable) GetDBPassword() string        { return s.db.password }
func (s *baseVariable) GetDBName() string            { return s.db.name }
