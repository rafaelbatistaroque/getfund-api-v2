package signout_application

import (
	auth_contract "getfund-api-v2/internal/domain/auth/core/contract"
	"getfund-api-v2/internal/domain/auth/core/usecase/signout"
	shared_error "getfund-api-v2/internal/shared/error"
)

type signoutApplication struct {
	sessionService auth_contract.SessionService
}

func New(sessionService auth_contract.SessionService) signout.UseCase {
	return &signoutApplication{
		sessionService: sessionService,
	}
}

func (u *signoutApplication) Execute(input *signout.Input) (*signout.Output, *shared_error.Error) {
	validated := input.Validate()
	if validated.IsInvalid() {
		return nil, shared_error.New(shared_error.UNAUTHORIZED_CODE, validated.GetErrors())
	}

	err := u.sessionService.DeleteSession(input.Token)
	if err != nil {
		return nil, shared_error.New(shared_error.SERVER_ERROR_CODE, err)
	}

	return &signout.SignoutOutput{Message: "user disconnected"}, nil
}
