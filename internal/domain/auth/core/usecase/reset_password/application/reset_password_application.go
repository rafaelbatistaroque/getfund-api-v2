package reset_password_application

import (
	"encoding/json"
	"errors"
	"getfund-api-v2/internal/domain/auth/core/auth_dto"
	auth_contract "getfund-api-v2/internal/domain/auth/core/contract"
	"getfund-api-v2/internal/domain/auth/core/usecase/reset_password"
	"getfund-api-v2/internal/shared/cache"
	shared_error "getfund-api-v2/internal/shared/error"
)

type resetPasswordApplication struct {
	cache      cache.Contract
	repository auth_contract.Repository
}

func New(cache cache.Contract, repository auth_contract.Repository) *resetPasswordApplication {
	return &resetPasswordApplication{
		cache:      cache,
		repository: repository,
	}
}

func (r *resetPasswordApplication) Execute(input *reset_password.Input) (*reset_password.Output, *shared_error.Error) {
	validatable := input.Validate()
	if validatable.IsInvalid() {
		return nil, shared_error.New(shared_error.BAD_REQUEST_CODE, validatable.GetErrors())
	}

	cacheData, errCache := r.cache.Get(input.RecoveryKey)
	if errCache != nil {
		return nil, shared_error.New(shared_error.NOT_FOUND_CODE, errors.New("recovery code not found"))
	}

	forgetPasswordDto := &auth_dto.ForgetPasswordDto{}
	errUnmarshal := json.Unmarshal([]byte(cacheData), forgetPasswordDto)
	if errUnmarshal != nil {
		return nil, shared_error.New(shared_error.SERVER_ERROR_CODE, errors.New("error on get recovery password data"))
	}

	authenticatedUser, errGetUser := r.repository.GetAuthenticatedUserByUsername(forgetPasswordDto.Username)
	if errGetUser != nil {
		return nil, shared_error.New(shared_error.NOT_FOUND_CODE, errGetUser)
	}

	if errUpdate := r.repository.UpdatePassword(authenticatedUser.Id, input.Password); errUpdate != nil {
		return nil, shared_error.New(shared_error.SERVER_ERROR_CODE, errUpdate)
	}

	defer r.cache.Delete(input.RecoveryKey)

	return &reset_password.Output{Message: "password updated"}, nil
}
