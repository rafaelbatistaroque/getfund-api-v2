package activate_user_application

import (
	"getfund-api-v2/internal/domain/user/core/usecase/activate_user"
	"getfund-api-v2/internal/shared/result_app"
	"getfund-api-v2/internal/shared/service/cache_service"
)

const (
	_KEY_USER_ACTIVATION_PREFIX = "user_activation_"
)

type activateUserApplication struct {
	cache cache_service.Cache
}

func New(cache cache_service.Cache) activate_user.UseCase {
	return &activateUserApplication{
		cache: cache,
	}
}

func (a *activateUserApplication) Execute(input *activate_user.Input) (*activate_user.Output, *result_app.ApplicationError) {
	validatable := input.Validate()
	if validatable.IsInvalid() {
		return nil, result_app.New(result_app.UNAUTHORIZED_CODE, validatable.GetErrors())
	}

	keyCache := _KEY_USER_ACTIVATION_PREFIX + input.ActivationCode
	a.cache.Get(keyCache)

	return nil, nil
}
