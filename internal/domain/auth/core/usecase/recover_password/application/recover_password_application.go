package recover_password_application

import (
	"getfund-api-v2/internal/config/env"
	auth_contract "getfund-api-v2/internal/domain/auth/core/contract"
	"getfund-api-v2/internal/domain/auth/core/dto"
	"getfund-api-v2/internal/domain/auth/core/usecase/recover_password"
	"getfund-api-v2/internal/domain/auth/core/usecase/recover_password/event"
	shared_bus "getfund-api-v2/internal/shared/bus"
	"getfund-api-v2/internal/shared/cache"
	shared_error "getfund-api-v2/internal/shared/error"
	"getfund-api-v2/internal/shared/security"
	"time"
)

const (
	key_cache_prefix = "recovery_password_"
)

type recoverPasswordApplication struct {
	hasher         security.Hasher
	env            env.Variable
	authRepository auth_contract.Repository
	cache          cache.Service
	bus            shared_bus.EventBus
}

func New(
	hasher security.Hasher,
	env env.Variable,
	authRepository auth_contract.Repository,
	cacheService cache.Service,
	bus shared_bus.EventBus) recover_password.UseCase {

	return &recoverPasswordApplication{
		hasher:         hasher,
		env:            env,
		authRepository: authRepository,
		cache:          cacheService,
		bus:            bus,
	}
}

func (uc *recoverPasswordApplication) Execute(input *recover_password.Input) (*recover_password.Output, *shared_error.Error) {
	validated := input.Validate()
	if validated.IsInvalid() {
		return nil, shared_error.New(shared_error.BAD_REQUEST_CODE, validated.GetErrors())
	}

	authenticatedUser, errRepo := uc.authRepository.GetAuthenticatedUserByUsername(input.Username)
	if errRepo != nil {
		return nil, shared_error.New(shared_error.NOT_FOUND_CODE, errRepo)
	}

	randomCode, errCode := uc.hasher.GetRandomCode(8)
	if errCode != nil {
		return nil, shared_error.New(shared_error.SERVER_ERROR_CODE, errCode)
	}

	recoveryCode, errHash := uc.hasher.Hash(randomCode, uc.env.GetServerSalt())
	if errHash != nil {
		return nil, shared_error.New(shared_error.SERVER_ERROR_CODE, errHash)
	}

	data := dto.ForgetPasswordDto{
		Username:     input.Username,
		FirstName:    authenticatedUser.FirstName,
		RecoveryLink: buildRecoverLink(uc.env, recoveryCode.Data),
	}

	keyCache := buildKeyCache(recoveryCode.Data)
	cacheErr := uc.cache.Set(keyCache, data, time.Hour)
	if cacheErr != nil {
		return nil, shared_error.New(shared_error.SERVER_ERROR_CODE, cacheErr)
	}

	uc.bus.EmitWithPayload(&event.RecoverPasswordStartedEvent{}, keyCache)

	return &recover_password.Output{Message: "recover password started"}, nil
}

func buildRecoverLink(env env.Variable, recoveryCode string) string {
	return env.GetBaseUrl() + "/new-password/" + recoveryCode
}

func buildKeyCache(recoveryCode string) string {
	return key_cache_prefix + recoveryCode
}
