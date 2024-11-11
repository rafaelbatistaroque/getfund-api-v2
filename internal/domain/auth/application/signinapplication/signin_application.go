package signinapplication

import (
	authServ "getfund-api-v2/internal/domain/auth/domainservice/authservice"
	sessionServ "getfund-api-v2/internal/domain/auth/domainservice/sessionservice"
	mapper "getfund-api-v2/internal/domain/auth/main/mapper/signinmapper"
	"getfund-api-v2/internal/domain/auth/usecase/signin"
	appErr "getfund-api-v2/internal/pkg/applicationerror"
	appCode "getfund-api-v2/internal/pkg/helpers/applicationcode"
)

type signinApplication struct {
	authService    authServ.AuthService
	sessionService sessionServ.SessionService
	mapper         mapper.SigninMapper
}

func NewUseCase(authService authServ.AuthService, sessionService sessionServ.SessionService, mapper mapper.SigninMapper) signin.UseCase {
	return &signinApplication{
		authService:    authService,
		sessionService: sessionService,
		mapper:         mapper,
	}
}

func (uc *signinApplication) Execute(input *signin.Input) (*signin.Output, *appErr.ApplicationError) {
	input.Validate()
	if input.IsInvalid() {
		return nil, appErr.New(appCode.BAD_REQUEST, input.GetErrors())
	}

	session, authErr := uc.authService.Authenticate(input.UserName, input.Password)
	if authErr != nil {
		return nil, authErr
	}

	buildTokenErr := uc.sessionService.BuildToken(session)
	if buildTokenErr != nil {
		return nil, appErr.New(appCode.CODE_SERVER_ERROR, buildTokenErr)
	}

	saveSessionErr := uc.sessionService.SaveSession(session)
	if saveSessionErr != nil {
		return nil, appErr.New(appCode.CODE_SERVER_ERROR, saveSessionErr)
	}

	return uc.mapper.ToOutput(session), nil
}
