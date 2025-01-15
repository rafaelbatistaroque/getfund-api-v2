package signin_application

import (
	auth_contract "getfund-api-v2/internal/domain/auth/core/contract"
	"getfund-api-v2/internal/domain/auth/core/domain_service/auth_service"
	"getfund-api-v2/internal/domain/auth/core/usecase/signin"
	"getfund-api-v2/internal/shared/result_app"
)

type signinApplication struct {
	authService    auth_service.AuthService
	sessionService auth_contract.SessionService
	mapper         auth_contract.SigninMapper
}

func New(authService auth_service.AuthService, sessionService auth_contract.SessionService, mapper auth_contract.SigninMapper) signin.UseCase {
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
