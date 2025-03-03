package activate_user

import "getfund-api-v2/internal/shared/result_app"

type UseCase = activateUser

type activateUser interface {
	Execute(input *Input) (*Output, *result_app.ApplicationError)
}
