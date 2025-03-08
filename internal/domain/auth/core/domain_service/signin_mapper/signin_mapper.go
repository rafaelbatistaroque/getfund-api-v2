package signin_mapper

import (
	"getfund-api-v2/internal/domain/auth/core/auth_dto"
	"getfund-api-v2/internal/domain/auth/core/usecase/signin"
)

type SigninMapper interface {
	ToOutput(token string, session *auth_dto.SessionDto) *signin.Output
	ToSessionModel(authenticatedUser *auth_dto.AuthenticatedUserDto) *auth_dto.SessionDto
}

type signinMapper struct {
}

// Constructor
func New() SigninMapper {
	return &signinMapper{}
}

func (m *signinMapper) ToOutput(token string, session *auth_dto.SessionDto) *signin.Output {
	return &signin.SigninOutput{
		Token: token,
		Session: signin.SessionOutput{
			ID:        session.ID,
			FirstName: session.FirstName,
			IsAdmin:   session.IsAdmin,
		},
	}
}

func (m *signinMapper) ToSessionModel(authenticatedUser *auth_dto.AuthenticatedUserDto) *auth_dto.SessionDto {
	return &auth_dto.SessionDto{
		ID:        authenticatedUser.Id,
		FirstName: authenticatedUser.FirstName,
		IsAdmin:   authenticatedUser.IsAdmin,
	}
}
