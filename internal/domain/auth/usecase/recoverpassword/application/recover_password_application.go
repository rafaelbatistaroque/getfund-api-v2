package recoverpasswordapplication

import (
	auth_contract "getfund-api-v2/internal/domain/auth/contract"
	"getfund-api-v2/internal/domain/auth/usecase/recoverpassword"
	"getfund-api-v2/internal/shared/contract/settings"
	"getfund-api-v2/internal/shared/resultapp"
	"getfund-api-v2/internal/shared/security"
	"getfund-api-v2/internal/shared/service/cacheservice"
	"getfund-api-v2/internal/shared/service/codeservice"
	"getfund-api-v2/pkg/eventbus"
	"getfund-api-v2/pkg/eventbus/event"
	"time"
)

var (
	key_cache_prefix = "recovery_password_"
)

type recoverPasswordApplication struct {
	hasher         security.Hasher
	settings       settings.ApplicationSettings
	userRepository auth_contract.UserRepository
	codeService    codeservice.CodeService
	cacheService   cacheservice.Cache
	eventBus       eventbus.EventBus
}

func New(
	hasher security.Hasher,
	settings settings.ApplicationSettings,
	userRepository auth_contract.UserRepository,
	codeService codeservice.CodeService,
	cacheService cacheservice.Cache,
	eventBus eventbus.EventBus) recoverpassword.UseCase {

	return &recoverPasswordApplication{
		hasher:         hasher,
		settings:       settings,
		userRepository: userRepository,
		codeService:    codeService,
		cacheService:   cacheService,
		eventBus:       eventBus,
	}
}

func (uc *recoverPasswordApplication) Execute(input *recoverpassword.Input) (*recoverpassword.Output, *resultapp.ApplicationError) {
	input.Validate()
	if input.IsInvalid() {
		return nil, resultapp.New(resultapp.BAD_REQUEST_CODE, input.GetErrors())
	}

	usernameHashed, err := uc.hasher.HashWithSalt(input.Username, uc.settings.GetServerSalt())
	if err != nil {
		return nil, resultapp.New(resultapp.SERVER_ERROR_CODE, err)
	}

	userModel, errRepo := uc.userRepository.GetByUserName(usernameHashed)
	if errRepo != nil {
		return nil, resultapp.New(resultapp.NOT_FOUND_CODE, errRepo)
	}

	randomCode, errCode := uc.codeService.GetRandomCode(8)
	if errCode != nil {
		return nil, resultapp.New(resultapp.SERVER_ERROR_CODE, errCode)
	}

	data := map[string]interface{}{
		"username":      input.Username,
		"first_name":    uc.hasher.DecryptMerged(userModel.FirstName, uc.settings.GetSecretKey()),
		"recovery_link": buildRecoverLink(uc.settings, randomCode),
	}

	keyCache := buildKeyCache(randomCode)
	cacheErr := uc.cacheService.Set(keyCache, data, time.Hour)
	if cacheErr != nil {
		return nil, resultapp.New(resultapp.SERVER_ERROR_CODE, cacheErr)
	}

	uc.eventBus.CreateAndPublish(&event.RecoverPasswordStartedEvent{}, keyCache)

	return &recoverpassword.RecoverPasswordOutput{Message: "recover password started"}, nil
}

func buildRecoverLink(applicationSettings settings.ApplicationSettings, randomCode string) string {
	return applicationSettings.GetBaseUrl() + "/new-password/" + randomCode
}

func buildKeyCache(randomCode string) string {
	return key_cache_prefix + randomCode
}
