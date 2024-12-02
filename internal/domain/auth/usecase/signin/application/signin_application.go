package signinapplication

import (
	authservice "getfund-api-v2/internal/domain/auth/domainservice/authservice"
	mapper "getfund-api-v2/internal/domain/auth/main/mapper/signinmapper"
	"getfund-api-v2/internal/domain/auth/usecase/signin"
	"getfund-api-v2/internal/shared/resultapp"
	sessionServ "getfund-api-v2/internal/shared/service/sessionservice"
)

type signinApplication struct {
	authService    authservice.AuthService
	sessionService sessionServ.SessionService
	mapper         mapper.SigninMapper
}

func New(authService authservice.AuthService, sessionService sessionServ.SessionService, mapper mapper.SigninMapper) signin.UseCase {
	return &signinApplication{
		authService:    authService,
		sessionService: sessionService,
		mapper:         mapper,
	}
}

func (uc *signinApplication) Execute(input *signin.Input) (*signin.Output, *resultapp.ApplicationError) {
	input.Validate()
	if input.IsInvalid() {
		return nil, resultapp.New(resultapp.BAD_REQUEST, input.GetErrors())
	}

	session, authErr := uc.authService.Authenticate(input.UserName, input.Password)
	if authErr != nil {
		return nil, authErr
	}

	sessionSerialized, toStringErr := uc.mapper.SessionToString(session)
	if toStringErr != nil {
		return nil, resultapp.New(resultapp.CODE_SERVER_ERROR, toStringErr)
	}

	token, saveSessionErr := uc.sessionService.SaveSession(sessionSerialized)
	if saveSessionErr != nil {
		return nil, resultapp.New(resultapp.CODE_SERVER_ERROR, saveSessionErr)
	}

	return uc.mapper.ToOutput(token, session), nil
}
