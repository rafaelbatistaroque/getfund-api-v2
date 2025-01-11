package signin_application

import (
	"getfund-api-v2/internal/domain/auth/core/domain_service/auth_service"
	"getfund-api-v2/internal/domain/auth/core/usecase/signin"
	mapper "getfund-api-v2/internal/domain/auth/main/mapper/signin_mapper"
	"getfund-api-v2/internal/shared/result_app"
	"getfund-api-v2/internal/shared/service/session_service"
)

type signinApplication struct {
	authService    auth_service.AuthService
	sessionService session_service.SessionService
	mapper         mapper.SigninMapper
}

func New(authService auth_service.AuthService, sessionService session_service.SessionService, mapper mapper.SigninMapper) signin.UseCase {
	return &signinApplication{
		authService:    authService,
		sessionService: sessionService,
		mapper:         mapper,
	}
}

func (uc *signinApplication) Execute(input *signin.Input) (*signin.Output, *result_app.ApplicationError) {
	validated := input.Validate()
	if validated.IsInvalid() {
		return nil, result_app.New(result_app.BAD_REQUEST_CODE, validated.GetErrors())
	}

	session, authErr := uc.authService.Authenticate(input.UserName, input.Password)
	if authErr != nil {
		return nil, authErr
	}

	token, saveSessionErr := uc.sessionService.SaveSession(session)
	if saveSessionErr != nil {
		return nil, result_app.New(result_app.SERVER_ERROR_CODE, saveSessionErr)
	}

	return uc.mapper.ToOutput(token, session), nil
}
