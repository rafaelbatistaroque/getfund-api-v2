package create_user

import "getfund-api-v2/internal/shared/result_app"

type UseCase = createUser

type createUser interface {
	Execute(input *Input) (*Output, *result_app.ApplicationError)
}
