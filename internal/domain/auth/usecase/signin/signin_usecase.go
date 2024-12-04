package signin

import "getfund-api-v2/internal/shared/resultapp"

type UseCase = signin

type signin interface {
	Execute(input *Input) (*Output, *resultapp.ApplicationError)
}
