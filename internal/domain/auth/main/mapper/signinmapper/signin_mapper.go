package signinmapper

import (
	entity "getfund-api-v2/internal/domain/auth/entity/sessionentity"
	"getfund-api-v2/internal/domain/auth/usecase/signin"
)

type SigninMapper interface {
	ToOutput(session entity.Session) *signin.Output
}

type signinMapper struct{}

// Constructor
func New() SigninMapper {
	return &signinMapper{}
}

func (m *signinMapper) ToOutput(session entity.Session) *signin.Output {
	return &signin.SigninOutput{
		Token: session.GetToken(),
		Session: signin.SessionOutput{
			Id:        session.GetID(),
			FirstName: session.GetFirstName(),
			IsAdmin:   session.GetIsAdmin(),
		},
	}
}
