package signout

import "getfund-api-v2/internal/shared/resultapp"

type UseCase = signout

type signout interface {
	Execute(input *Input) (*Output, *resultapp.ApplicationError)
}
