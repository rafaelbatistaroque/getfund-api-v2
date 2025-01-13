package signout_application

import (
	auth_contract "getfund-api-v2/internal/domain/auth/core/contract"
	"getfund-api-v2/internal/domain/auth/core/usecase/signout"
	"getfund-api-v2/internal/shared/result_app"
)

type signoutApplication struct {
	sessionService auth_contract.SessionService
}

func New(sessionService auth_contract.SessionService) signout.UseCase {
	return &signoutApplication{
		sessionService: sessionService,
	}
}

func (u *signoutApplication) Execute(input *signout.Input) (*signout.Output, *result_app.ApplicationError) {
	validated := input.Validate()
	if validated.IsInvalid() {
		return nil, result_app.New(result_app.UNAUTHORIZED_CODE, validated.GetErrors())
	}

	err := u.sessionService.DeleteSession(input.Token)
	if err != nil {
		return nil, result_app.New(result_app.SERVER_ERROR_CODE, err)
	}

	return &signout.SignoutOutput{Message: "user disconnected"}, nil
}
