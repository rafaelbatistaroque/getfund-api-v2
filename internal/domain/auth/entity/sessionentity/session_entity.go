package sessionentity

import "errors"

var (
	err_ID_INVALID        = errors.New("session id invalid")
	err_FIRSTNAME_INVALID = errors.New("session firstName invalid")
	err_ROLE_INVALID      = errors.New("session role invalid")
	err_SET_TOKEN_INVALID = errors.New("error on set token")
)

// Interface que encapsula a entidade
type Session interface {
	GetID() string
	GetFirstName() string
	GetIsAdmin() bool
	GetToken() string
	SetToken(token string) error
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
func (s *session) SetToken(token string) error {
	if len(token) == 0 {
		return err_SET_TOKEN_INVALID
	}
	s.token = token

	return nil
}

// Construtor
func New(id string, firstName string, role int) (Session, error) {
	if err := validate(id, firstName, role); err != nil {
		return nil, err
	}

	return &session{
		id:        id,
		firstName: firstName,
		isAdmin:   isAdmin(role),
	}, nil
}

func isAdmin(role int) bool {
	return role == 1
}

func validate(id string, firstName string, role int) error {
	if len(id) == 0 {
		return err_ID_INVALID
	}
	if len(firstName) == 0 {
		return err_FIRSTNAME_INVALID
	}
	if role < 0 {
		return err_ROLE_INVALID
	}
	return nil
}
