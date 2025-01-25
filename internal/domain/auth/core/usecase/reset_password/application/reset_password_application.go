package reset_password_application

import (
	"encoding/json"
	"errors"
	"getfund-api-v2/internal/domain/auth/core/auth_dto"
	auth_contract "getfund-api-v2/internal/domain/auth/core/contract"
	"getfund-api-v2/internal/domain/auth/core/usecase/reset_password"
	"getfund-api-v2/internal/shared/result_app"
	"getfund-api-v2/internal/shared/service/cache_service"
)

const (
	_KEY_CACHE_PREFIX = "recovery_password_"
)

type resetPasswordApplication struct {
	cacheService   cache_service.Cache
	authRepository auth_contract.AuthRepository
}

func New(cacheService cache_service.Cache, authRepository auth_contract.AuthRepository) *resetPasswordApplication {
	return &resetPasswordApplication{
		cacheService:   cacheService,
		authRepository: authRepository,
	}
}

func (r *resetPasswordApplication) Execute(input *reset_password.Input) (*reset_password.Output, *result_app.ApplicationError) {
	validatable := input.Validate()
	if validatable.IsInvalid() {
		return nil, result_app.New(result_app.BAD_REQUEST_CODE, validatable.GetErrors())
	}

	keyRecoveryCode := _KEY_CACHE_PREFIX + input.RecoveryCode
	defer r.cacheService.Delete(keyRecoveryCode)
	cacheData, errCache := r.cacheService.Get(keyRecoveryCode)
	if errCache != nil {
		return nil, result_app.New(result_app.NOT_FOUND_CODE, errors.New("recovery code not found"))
	}

	forgetPasswordDto := &auth_dto.ForgetPasswordDto{}
	errUnmarshal := json.Unmarshal([]byte(cacheData), forgetPasswordDto)
	if errUnmarshal != nil {
		return nil, result_app.New(result_app.SERVER_ERROR_CODE, errors.New("error to unmarshal data"))
	}

	authenticatedUser, errGetUser := r.authRepository.GetAuthenticatedUserByUsername(forgetPasswordDto.Username)
	if errGetUser != nil {
		return nil, result_app.New(result_app.NOT_FOUND_CODE, errGetUser)
	}

	if errUpdate := r.authRepository.UpdatePassword(authenticatedUser.Id, input.Password); errUpdate != nil {
		return nil, result_app.New(result_app.SERVER_ERROR_CODE, errUpdate)
	}

	return &reset_password.Output{Message: "password updated"}, nil
}
