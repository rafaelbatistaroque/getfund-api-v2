package signout

import "getfund-api-v2/internal/shared/result_app"

type UseCase = signout

type signout interface {
	Execute(input *Input) (*Output, *result_app.ApplicationError)
}
