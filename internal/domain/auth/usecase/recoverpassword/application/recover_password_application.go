package recoverpasswordapplication

import (
	auth_contract "getfund-api-v2/internal/domain/auth/contract"
	"getfund-api-v2/internal/domain/auth/usecase/recoverpassword"
	"getfund-api-v2/internal/shared/resultapp"
	"getfund-api-v2/internal/shared/security"
	"getfund-api-v2/pkg/eventbus"
)

type recoverPasswordApplication struct {
	hasher         security.Hasher
	userRepository auth_contract.UserRepository
	//service to get random code
	eventBus eventbus.EventBus
}

func (uc *recoverPasswordApplication) Execute(input *recoverpassword.Input) (*recoverpassword.Output, *resultapp.ApplicationError) {
	//fail fast validation
	//decrypt username with hasher
	//get user in repository with repository
	//invoke a new service to get random code
	//save user_email, user_firstname, recovery_link and recovery_code with specific key in cache by a hour
	//publish event RecoverPasswordStarted with key cache
	//return success with message

	//Handler
	//recover data cached received by key
	//build a email template with params to replace
	//replace specific
	//send email
	return nil, nil
}
