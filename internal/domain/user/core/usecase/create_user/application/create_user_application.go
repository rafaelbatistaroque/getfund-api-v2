package create_user_application

import (
	"errors"
	user_contract "getfund-api-v2/internal/domain/user/core/contract"
	payload "getfund-api-v2/internal/domain/user/core/dto/user_payload"
	"getfund-api-v2/internal/domain/user/core/usecase/create_user"
	"getfund-api-v2/internal/settings"
	"getfund-api-v2/internal/shared/result_app"
	"getfund-api-v2/internal/shared/security"
	"getfund-api-v2/internal/shared/service/cache_service"
	"getfund-api-v2/pkg/bus"
	"time"
)

const (
	KEY_CACHE_PREFIX = "user_activation_"
)

type createUserApplication struct {
	repository user_contract.Repository
	hasher     security.Hasher
	cache      cache_service.Cache
	bus        bus.EventBus
	settings   settings.ApplicationSettings
}

func New(repository user_contract.Repository, hasher security.Hasher, cache cache_service.Cache, bus bus.EventBus, settings settings.ApplicationSettings) create_user.UseCase {
	return &createUserApplication{
		repository: repository,
		hasher:     hasher,
		cache:      cache,
		bus:        bus,
		settings:   settings,
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

	activationCode, errCode := buildActivationCode(c.hasher)
	if errCode != nil {
		return nil, result_app.New(result_app.SERVER_ERROR_CODE, errCode)
	}

	if err := c.cache.Set(KEY_CACHE_PREFIX+activationCode, input, 24*time.Hour); err != nil {
		return nil, result_app.New(result_app.SERVER_ERROR_CODE, errors.New("error to save user"))
	}

	emitCreateUserProcessStartedEvent(c.bus, c.settings, activationCode)

	return &create_user.Output{Message: "user creation started"}, nil
}

func validateDuplicatedUser(input *create_user.Input, repository user_contract.Repository) *result_app.ApplicationError {
	userDuplicated, err := repository.UserExistsByUsername(input.Email)
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

	return activationCode, nil
}

func emitCreateUserProcessStartedEvent(bus bus.EventBus, settings settings.ApplicationSettings, activationCode string) {
	payload := &payload.CreateUserProcessPayload{
		ActivationDataKey: KEY_CACHE_PREFIX + activationCode,
		ActivationCode:    activationCode,
		ActivationLink:    settings.GetBaseUrl() + "/user-activation/" + activationCode,
	}

	bus.EmitWithPayload(&create_user.CreateUserProcessStartedEvent{}, payload)
}
