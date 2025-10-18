package signin_application

import (
	auth_contract "getfund-api-v2/internal/domain/auth/core/contract"
	"getfund-api-v2/internal/domain/auth/core/domain_service/auth_service"
	"getfund-api-v2/internal/domain/auth/core/domain_service/signin_mapper"
	"getfund-api-v2/internal/domain/auth/core/usecase/signin"
	shared_error "getfund-api-v2/internal/shared/error"
)

type signinApplication struct {
	authService    auth_service.AuthService
	sessionService auth_contract.SessionService
	mapper         signin_mapper.SigninMapper
}

func New(authService auth_service.AuthService, sessionService auth_contract.SessionService, mapper signin_mapper.SigninMapper) signin.UseCase {
	return &signinApplication{
		authService:    authService,
		sessionService: sessionService,
		mapper:         mapper,
	}
}

func (uc *signinApplication) Execute(input *signin.Input) (*signin.Output, *shared_error.Error) {
	validated := input.Validate()
	if validated.IsInvalid() {
		return nil, shared_error.New(shared_error.BAD_REQUEST_CODE, validated.GetErrors())
	}

	session, authErr := uc.authService.Authenticate(input.Username, input.Password)
	if authErr != nil {
		return nil, authErr
	}

	token, saveSessionErr := uc.sessionService.SaveSession(session)
	if saveSessionErr != nil {
		return nil, shared_error.New(shared_error.SERVER_ERROR_CODE, saveSessionErr)
	}

	return uc.mapper.ToOutput(token, session), nil
}
