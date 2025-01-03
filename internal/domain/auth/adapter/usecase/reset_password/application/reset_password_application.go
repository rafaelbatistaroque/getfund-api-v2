package reset_password_application

import (
	"getfund-api-v2/internal/domain/auth/adapter/usecase/reset_password"
	"getfund-api-v2/internal/shared/result_app"
)

type resetPasswordApplication struct {
}

func New() *resetPasswordApplication {
	return &resetPasswordApplication{}
}

func (r *resetPasswordApplication) Execute(input *reset_password.Input) (*reset_password.Output, *result_app.ApplicationError) {
	validatable := input.Validate()
	if validatable.IsInvalid() {
		return nil, result_app.New(result_app.BAD_REQUEST_CODE, validatable.GetErrors())
	}
	return nil, nil
}
