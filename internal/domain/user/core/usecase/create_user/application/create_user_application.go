package create_user_application

import (
	"errors"
	user_contract "getfund-api-v2/internal/domain/user/core/contract"
	"getfund-api-v2/internal/domain/user/core/usecase/create_user"
	"getfund-api-v2/internal/shared/result_app"
)

type createUserApplication struct {
	repository user_contract.Repository
}

func New(repository user_contract.Repository) create_user.UseCase {
	return &createUserApplication{
		repository: repository,
	}
}

func (c *createUserApplication) Execute(input *create_user.Input) (*create_user.Output, *result_app.ApplicationError) {
	validated := input.Validate()
	if validated.IsInvalid() {
		return nil, result_app.New(result_app.BAD_REQUEST_CODE, validated.GetErrors())
	}

	userDuplicated, err := c.repository.GetUserByUsername(input.Email)
	if err != nil {
		return nil, result_app.New(result_app.NOT_FOUND_CODE, err)
	}

	if userDuplicated != nil {
		return nil, result_app.New(result_app.DUPLICATED_ENTRY_CODE, errors.New("User already exists"))
	}

	return nil, nil
}
