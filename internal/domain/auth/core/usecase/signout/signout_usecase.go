package signout

import (
	shared_error "getfund-api-v2/internal/shared/error"
)

type UseCase = signout

type signout interface {
	Execute(input *Input) (*Output, *shared_error.Error)
}
