package recover_password_application

import (
	auth_contract "getfund-api-v2/internal/domain/auth/core/contract"
	recoverpassword "getfund-api-v2/internal/domain/auth/core/usecase/recover_password"
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
	userRepository auth_contract.UserRepository
	cacheService   cache_service.Cache
	eventBus       bus.EventBus
}

func New(
	hasher security.Hasher,
	settings settings.ApplicationSettings,
	userRepository auth_contract.UserRepository,
	cacheService cache_service.Cache,
	eventBus bus.EventBus) recoverpassword.UseCase {

	return &recoverPasswordApplication{
		hasher:         hasher,
		settings:       settings,
		userRepository: userRepository,
		cacheService:   cacheService,
		eventBus:       eventBus,
	}
}

func (uc *recoverPasswordApplication) Execute(input *recoverpassword.Input) (*recoverpassword.Output, *result_app.ApplicationError) {
	validated := input.Validate()
	if validated.IsInvalid() {
		return nil, result_app.New(result_app.BAD_REQUEST_CODE, validated.GetErrors())
	}

	usernameHashed, err := uc.hasher.HashWithSalt(input.Username, uc.settings.GetServerSalt())
	if err != nil {
		return nil, result_app.New(result_app.SERVER_ERROR_CODE, err)
	}

	userModel, errRepo := uc.userRepository.GetByUserName(usernameHashed)
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

	data := map[string]interface{}{
		"username":      input.Username,
		"first_name":    uc.hasher.DecryptMerged(userModel.FirstName, uc.settings.GetSecretKey()),
		"recovery_link": buildRecoverLink(uc.settings, recoveryCode.Data),
	}

	keyCache := buildKeyCache(recoveryCode.Data)
	cacheErr := uc.cacheService.Set(keyCache, data, time.Hour)
	if cacheErr != nil {
		return nil, result_app.New(result_app.SERVER_ERROR_CODE, cacheErr)
	}

	uc.eventBus.CreateAndPublish(&event.RecoverPasswordStarted{}, keyCache)

	return &recoverpassword.RecoverPasswordOutput{Message: "recover password started"}, nil
}

func buildRecoverLink(applicationSettings settings.ApplicationSettings, recoveryCode string) string {
	return applicationSettings.GetBaseUrl() + "/new-password/" + recoveryCode
}

func buildKeyCache(recoveryCode string) string {
	return key_cache_prefix + recoveryCode
}
