package reset_password_application

import (
	"encoding/json"
	"errors"
	auth_contract "getfund-api-v2/internal/domain/auth/core/contract"
	auth_model "getfund-api-v2/internal/domain/auth/core/model"
	"getfund-api-v2/internal/domain/auth/core/usecase/reset_password"
	"getfund-api-v2/internal/shared/result_app"
	"getfund-api-v2/internal/shared/service/cache_service"
)

var (
	_KEY_CACHE_PREFIX = "recovery_password_"
)

type resetPasswordApplication struct {
	cacheService   cache_service.Cache
	userRepository auth_contract.UserRepository
}

func New(cacheService cache_service.Cache, userRepository auth_contract.UserRepository) *resetPasswordApplication {
	return &resetPasswordApplication{
		cacheService:   cacheService,
		userRepository: userRepository,
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
		return nil, result_app.New(result_app.NOT_FOUND_CODE, errCache)
	}

	forgetPasswordModel := &auth_model.ForgetPasswordModel{}
	errUnmarshal := json.Unmarshal([]byte(cacheData), forgetPasswordModel)
	if errUnmarshal != nil {
		return nil, result_app.New(result_app.SERVER_ERROR_CODE, errors.New("error to unmarshal data"))
	}

	_, errGetUser := r.userRepository.GetByUserName(forgetPasswordModel.Username)
	if errGetUser != nil {
		return nil, result_app.New(result_app.NOT_FOUND_CODE, errGetUser)
	}

	return nil, nil
}
