package create_user_application

import (
	"getfund-api-v2/internal/domain/user/core/usercase/create_user"
	"getfund-api-v2/internal/shared/result_app"
)

type createUserApplication struct {
}

func New() create_user.UseCase {
	return &createUserApplication{}
}

func (c *createUserApplication) Execute(input *create_user.Input) (*create_user.Output, *result_app.ApplicationError) {
	validated := input.Validate()
	if validated.IsInvalid() {
		return nil, result_app.New(result_app.BAD_REQUEST_CODE, validated.GetErrors())
	}
	panic("unimplemented")
}
