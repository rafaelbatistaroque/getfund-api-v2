package reset_password_application

import (
	"getfund-api-v2/internal/domain/auth/core/usecase/reset_password"
	"getfund-api-v2/internal/shared/result_app"
	"getfund-api-v2/internal/shared/service/cache_service"
)

var (
	_KEY_CACHE_PREFIX = "recovery_password_"
)

type resetPasswordApplication struct {
	cacheService cache_service.Cache
}

func New(cacheService cache_service.Cache) *resetPasswordApplication {
	return &resetPasswordApplication{
		cacheService: cacheService,
	}
}

func (r *resetPasswordApplication) Execute(input *reset_password.Input) (*reset_password.Output, *result_app.ApplicationError) {
	validatable := input.Validate()
	if validatable.IsInvalid() {
		return nil, result_app.New(result_app.BAD_REQUEST_CODE, validatable.GetErrors())
	}

	keyRecoveryCode := _KEY_CACHE_PREFIX + input.RecoveryCode
	defer r.cacheService.Delete(keyRecoveryCode)
	_, errCache := r.cacheService.Get(keyRecoveryCode)
	if errCache != nil {
		return nil, result_app.New(result_app.NOT_FOUND_CODE, errCache)
	}

	return nil, nil
}
