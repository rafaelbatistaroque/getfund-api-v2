package signin

import "getfund-api-v2/internal/shared/result_app"

type UseCase = signin

type signin interface {
	Execute(input *Input) (*Output, *result_app.ApplicationError)
}
