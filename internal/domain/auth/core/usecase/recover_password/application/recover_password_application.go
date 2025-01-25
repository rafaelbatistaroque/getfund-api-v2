package recover_password_application

import (
	"getfund-api-v2/internal/domain/auth/core/auth_dto"
	auth_contract "getfund-api-v2/internal/domain/auth/core/contract"
	"getfund-api-v2/internal/domain/auth/core/usecase/recover_password"
	"getfund-api-v2/internal/shared/contract/settings"
	"getfund-api-v2/internal/shared/result_app"
	"getfund-api-v2/internal/shared/security"
	"getfund-api-v2/internal/shared/service/cache_service"
	"getfund-api-v2/pkg/bus"
	"getfund-api-v2/pkg/bus/event"
	"time"
)

var (
	key_cache_prefix = "recovery_password_"
)

type recoverPasswordApplication struct {
	hasher         security.Hasher
	settings       settings.ApplicationSettings
	authRepository auth_contract.AuthRepository
	cacheService   cache_service.Cache
	bus            bus.EventBus
}

func New(
	hasher security.Hasher,
	settings settings.ApplicationSettings,
	authRepository auth_contract.AuthRepository,
	cacheService cache_service.Cache,
	bus bus.EventBus) recover_password.UseCase {

	return &recoverPasswordApplication{
		hasher:         hasher,
		settings:       settings,
		authRepository: authRepository,
		cacheService:   cacheService,
		bus:            bus,
	}
}

func (uc *recoverPasswordApplication) Execute(input *recover_password.Input) (*recover_password.Output, *result_app.ApplicationError) {
	validated := input.Validate()
	if validated.IsInvalid() {
		return nil, result_app.New(result_app.BAD_REQUEST_CODE, validated.GetErrors())
	}

	authenticatedUser, errRepo := uc.authRepository.GetAuthenticatedUserByUsername(input.Username)
	if errRepo != nil {
		return nil, result_app.New(result_app.NOT_FOUND_CODE, errRepo)
	}

	randomCode, errCode := uc.hasher.GetRandomCode(8)
	if errCode != nil {
		return nil, result_app.New(result_app.SERVER_ERROR_CODE, errCode)
	}

	recoveryCode, errHash := uc.hasher.Hash(randomCode, uc.settings.GetServerSalt())
	if errHash != nil {
		return nil, result_app.New(result_app.SERVER_ERROR_CODE, errHash)
	}

	data := auth_dto.ForgetPasswordDto{
		Username:     input.Username,
		FirstName:    authenticatedUser.FirstName,
		RecoveryLink: buildRecoverLink(uc.settings, recoveryCode.Data),
	}

	keyCache := buildKeyCache(recoveryCode.Data)
	cacheErr := uc.cacheService.Set(keyCache, data, time.Hour)
	if cacheErr != nil {
		return nil, result_app.New(result_app.SERVER_ERROR_CODE, cacheErr)
	}

	uc.bus.EmitWithPayload(&event.RecoverPasswordStarted{}, keyCache)

	return &recover_password.Output{Message: "recover password started"}, nil
}

func buildRecoverLink(applicationSettings settings.ApplicationSettings, recoveryCode string) string {
	return applicationSettings.GetBaseUrl() + "/new-password/" + recoveryCode
}

func buildKeyCache(recoveryCode string) string {
	return key_cache_prefix + recoveryCode
}
