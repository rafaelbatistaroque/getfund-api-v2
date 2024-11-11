package sessionentity

// Interface que encapsula a entidade
type Session interface {
	GetID() string
	GetFirstName() string
	GetIsAdmin() bool
	GetToken() string
	SetToken(token string)
}

// Entidade
type session struct {
	id        string
	firstName string
	isAdmin   bool
	token     string
}

// Getters e Setters
func (s *session) GetID() string        { return s.id }
func (s *session) GetFirstName() string { return s.firstName }
func (s *session) GetIsAdmin() bool     { return s.isAdmin }
func (s *session) GetToken() string     { return s.token }
func (s *session) SetToken(token string) {
	s.token = token
}

// Construtor
func New(id string, firstName string, role int) Session {
	return &session{
		id:        id,
		firstName: firstName,
		isAdmin:   isAdmin(role),
	}
}

func isAdmin(role int) bool {
	return role == 1
}
