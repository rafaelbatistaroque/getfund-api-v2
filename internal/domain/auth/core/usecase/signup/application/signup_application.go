package signup_application

import (
	"errors"
	"getfund-api-v2/internal/config/env"
	auth_contract "getfund-api-v2/internal/domain/auth/core/contract"
	"getfund-api-v2/internal/domain/auth/core/usecase/signup"
	"getfund-api-v2/internal/domain/auth/core/usecase/signup/event"
	shared_bus "getfund-api-v2/internal/shared/bus"
	"getfund-api-v2/internal/shared/cache"
	shared_constant "getfund-api-v2/internal/shared/constant"
	shared_error "getfund-api-v2/internal/shared/error"
	"getfund-api-v2/internal/shared/security"

	"time"
)

type SignupApplication struct {
	repository auth_contract.Repository
	hasher     security.Hasher
	cache      cache.Service
	bus        shared_bus.EventBus
	env        env.Variable
}

func New(repository auth_contract.Repository, hasher security.Hasher, cache cache.Service, bus shared_bus.EventBus, env env.Variable) signup.UseCase {
	return &SignupApplication{
		repository: repository,
		hasher:     hasher,
		cache:      cache,
		bus:        bus,
		env:        env,
	}
}

func (c *SignupApplication) Execute(input *signup.Input) (*signup.Output, *shared_error.Error) {
	validated := input.Validate()
	if validated.IsInvalid() {
		return nil, shared_error.New(shared_error.UNPROCESSABLE_CONTENT_CODE, validated.GetErrors())
	}

	if invalidUser := validateDuplicatedUser(input, c.repository); invalidUser != nil {
		return nil, invalidUser
	}

	activationCode, errCode := buildActivationCode(c.hasher)
	if errCode != nil {
		return nil, shared_error.New(shared_error.SERVER_ERROR_CODE, errCode)
	}

	if err := c.cache.Set(shared_constant.UserActivationCacheKeyPrefix+activationCode, input, 24*time.Hour); err != nil {
		return nil, shared_error.New(shared_error.SERVER_ERROR_CODE, errors.New("error to save user"))
	}

	emitSignupProcessStartedEvent(c.bus, c.env, input.FirstName, input.Username, activationCode)

	return &signup.Output{Message: "user creation started"}, nil
}

func validateDuplicatedUser(input *signup.Input, repository auth_contract.Repository) *shared_error.Error {
	userDuplicated, err := repository.UserExists(input.Username)

	if err != nil {
		return shared_error.New(shared_error.SERVER_ERROR_CODE, err)
	}

	if userDuplicated != nil {
		return shared_error.New(shared_error.DUPLICATED_ENTRY_CODE, errors.New("user already exists"))
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

func emitSignupProcessStartedEvent(bus shared_bus.EventBus, env env.Variable, firstName string, email string, activationCode string) {
	payload := &event.SignupStartedPayload{
		FirstName:      firstName,
		Email:          email,
		ActivationCode: activationCode,
		ActivationLink: env.GetBaseUrl() + "/user-activation/" + activationCode,
	}

	bus.EmitWithPayload(&event.SignupStartedEvent{}, payload)
}
