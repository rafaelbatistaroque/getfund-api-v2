package signoutapplication

import (
	"getfund-api-v2/internal/domain/auth/usecase/signout"
	"getfund-api-v2/internal/shared/resultapp"
	"getfund-api-v2/internal/shared/service/sessionservice"
)

type signoutApplication struct {
	sessionService sessionservice.SessionService
}

func New(sessionService sessionservice.SessionService) signout.UseCase {
	return &signoutApplication{
		sessionService: sessionService,
	}
}

func (u *signoutApplication) Execute(input *signout.Input) (*signout.Output, *resultapp.ApplicationError) {
	input.Validate()
	if input.IsInvalid() {
		return nil, resultapp.New(resultapp.UNAUTHORIZED_CODE, input.GetErrors())
	}

	err := u.sessionService.DeleteSession(input.Token)
	if err != nil {
		return nil, resultapp.New(resultapp.SERVER_ERROR_CODE, err)
	}

	return &signout.SignoutOutput{Message: "user disconnected"}, nil
}
