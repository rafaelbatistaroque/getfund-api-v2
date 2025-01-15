package signin_mapper

import (
	model "getfund-api-v2/internal/domain/auth/core/auth_dto"
	auth_contract "getfund-api-v2/internal/domain/auth/core/contract"
	"getfund-api-v2/internal/domain/auth/core/usecase/signin"
)

type signinMapper struct {
}

// Constructor
func New() auth_contract.SigninMapper {
	return &signinMapper{}
}

func (m *signinMapper) ToOutput(token string, session *model.SessionDto) *signin.Output {
	return &signin.SigninOutput{
		Token: token,
		Session: signin.SessionOutput{
			ID:        session.ID,
			FirstName: session.FirstName,
			IsAdmin:   session.IsAdmin == 1,
		},
	}
}

func (m *signinMapper) ToSessionModel(authenticatedUser *model.AuthenticatedUserDto) *model.SessionDto {
	return &model.SessionDto{
		ID:        authenticatedUser.Id,
		FirstName: authenticatedUser.FirstName,
		IsAdmin:   authenticatedUser.IsAdmin,
	}
}
