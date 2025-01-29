package create_user_application

import (
	"errors"
	user_contract "getfund-api-v2/internal/domain/user/core/contract"
	"getfund-api-v2/internal/domain/user/core/usecase/create_user"
	"getfund-api-v2/internal/domain/user/core/user_dto"
	"getfund-api-v2/internal/shared/result_app"
	"getfund-api-v2/internal/shared/security"
	"getfund-api-v2/internal/shared/service/cache_service"
	"getfund-api-v2/pkg/bus"
	"getfund-api-v2/pkg/bus/event"
	"time"
)

const (
	key_cache_prefix = "user_activation_"
)

type createUserApplication struct {
	repository user_contract.Repository
	hasher     security.Hasher
	cache      cache_service.Cache
	bus        bus.EventBus
}

func New(repository user_contract.Repository, hasher security.Hasher, cache cache_service.Cache, bus bus.EventBus) create_user.UseCase {
	return &createUserApplication{
		repository: repository,
		hasher:     hasher,
		cache:      cache,
		bus:        bus,
	}
}

func (c *createUserApplication) Execute(input *create_user.Input) (*create_user.Output, *result_app.ApplicationError) {
	validated := input.Validate()
	if validated.IsInvalid() {
		return nil, result_app.New(result_app.BAD_REQUEST_CODE, validated.GetErrors())
	}

	if invalidUser := validateDuplicatedUser(input, c.repository); invalidUser != nil {
		return nil, invalidUser
	}

	keyCache, errCode := buildActivationCode(c.hasher)
	if errCode != nil {
		return nil, result_app.New(result_app.SERVER_ERROR_CODE, errCode)
	}

	if err := c.cache.Set(keyCache, input, 24*time.Hour); err != nil {
		return nil, result_app.New(result_app.SERVER_ERROR_CODE, errors.New("error to save user"))
	}

	emitUserCriationStartedEvent(c.bus, input, keyCache)

	if input.CouponCode != "" {
		emitUserCriationWithCouponCodeStartedEvent(c.bus, input, keyCache)
	}

	return nil, nil
}

func validateDuplicatedUser(input *create_user.Input, repository user_contract.Repository) *result_app.ApplicationError {
	userDuplicated, err := repository.GetUserByUsername(input.Email)
	if err != nil {
		return result_app.New(result_app.SERVER_ERROR_CODE, err)
	}

	if userDuplicated != nil {
		return result_app.New(result_app.DUPLICATED_ENTRY_CODE, errors.New("user already exists"))
	}

	return nil
}

func buildActivationCode(hasher security.Hasher) (string, error) {
	activationCode, errCode := hasher.GetRandomCode(20)
	if errCode != nil {
		return "", errors.New("error to save user")
	}

	return key_cache_prefix + activationCode, nil
}

func emitUserCriationStartedEvent(bus bus.EventBus, input *create_user.Input, activationCode string) {
	userCriationStarted := &user_dto.UserCriationStartedDto{
		ActivationCode: activationCode,
		FirstName:      input.FirstName,
		Email:          input.Email,
	}

	bus.EmitWithPayload(&event.UserCriationStartedEvent{}, userCriationStarted)
}

func emitUserCriationWithCouponCodeStartedEvent(bus bus.EventBus, input *create_user.Input, activationCode string) {
	payload := &user_dto.UserCriationWithCouponStardedDto{
		CouponCode:     input.CouponCode,
		ActivationCode: activationCode,
	}

	bus.EmitWithPayload(&event.UserCriationWithCouponCodeStartedEvent{}, payload)
}
