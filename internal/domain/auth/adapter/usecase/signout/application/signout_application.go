package signout_application

import (
	"getfund-api-v2/internal/domain/auth/adapter/usecase/signout"
	"getfund-api-v2/internal/shared/result_app"
	"getfund-api-v2/internal/shared/service/session_service"
)

type signoutApplication struct {
	sessionService session_service.SessionService
}

func New(sessionService session_service.SessionService) signout.UseCase {
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
