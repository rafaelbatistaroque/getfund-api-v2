package recoverpasswordapplication

import (
	auth_contract "getfund-api-v2/internal/domain/auth/contract"
	"getfund-api-v2/internal/domain/auth/usecase/recoverpassword"
	"getfund-api-v2/internal/shared/contract/settings"
	"getfund-api-v2/internal/shared/resultapp"
	"getfund-api-v2/internal/shared/security"
	"getfund-api-v2/internal/shared/service/codeservice"
	"getfund-api-v2/pkg/eventbus"
)

type recoverPasswordApplication struct {
	hasher         security.Hasher
	settings       settings.ApplicationSettings
	userRepository auth_contract.UserRepository
	codeService    codeservice.CodeService
	//service to get random code
	eventBus eventbus.EventBus
}

func New(hasher security.Hasher, settings settings.ApplicationSettings, userRepository auth_contract.UserRepository, codeService codeservice.CodeService) recoverpassword.UseCase {

	return &recoverPasswordApplication{
		hasher:         hasher,
		settings:       settings,
		userRepository: userRepository,
		codeService:    codeService,
		eventBus:       nil,
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

	_, errRepo := uc.userRepository.GetByUserName(usernameHashed)
	if errRepo != nil {
		return nil, resultapp.New(resultapp.NOT_FOUND_CODE, errRepo)
	}

	//TODO: invoke a new service to get random code
	_, errCode := uc.codeService.GetRandomCode(8)
	if errCode != nil {
		return nil, resultapp.New(resultapp.SERVER_ERROR_CODE, errCode)
	}
	//TODO: save user_email (username), user_firstname, recovery_link and recovery_code with specific key in cache by a hour
	//TODO: publish event RecoverPasswordStarted with key cache
	//TODO: return success with message

	//Handler
	//TODO: recover data cached by key received
	//TODO: build a email template with params to replace
	//TODO: replace specific
	//TODO: send email
	return nil, nil
}
