package recoverpasswordapplication

import (
	auth_contract "getfund-api-v2/internal/domain/auth/contract"
	"getfund-api-v2/internal/domain/auth/usecase/recoverpassword"
	"getfund-api-v2/internal/shared/contract/settings"
	"getfund-api-v2/internal/shared/resultapp"
	"getfund-api-v2/internal/shared/security"
	"getfund-api-v2/pkg/eventbus"
)

type recoverPasswordApplication struct {
	hasher         security.Hasher
	settings       settings.ApplicationSettings
	userRepository auth_contract.UserRepository
	//service to get random code
	eventBus eventbus.EventBus
}

func New(hasher security.Hasher, settings settings.ApplicationSettings) recoverpassword.UseCase {

	return &recoverPasswordApplication{
		hasher:         hasher,
		settings:       settings,
		userRepository: nil,
		eventBus:       nil,
	}
}

func (uc *recoverPasswordApplication) Execute(input *recoverpassword.Input) (*recoverpassword.Output, *resultapp.ApplicationError) {
	//fail fast validation
	input.Validate()
	if input.IsInvalid() {
		return nil, resultapp.New(resultapp.BAD_REQUEST_CODE, input.GetErrors())
	}

	uc.hasher.HashWithSalt(input.Username, uc.settings.GetServerSalt())

	//decrypt username with hasher
	//get user in repository with repository
	//invoke a new service to get random code
	//save user_email, user_firstname, recovery_link and recovery_code with specific key in cache by a hour
	//publish event RecoverPasswordStarted with key cache
	//return success with message

	//Handler
	//recover data cached by key received
	//build a email template with params to replace
	//replace specific
	//send email
	return nil, nil
}
