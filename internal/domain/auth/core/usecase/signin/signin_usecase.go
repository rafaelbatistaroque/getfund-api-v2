package signin

import (
	shared_error "getfund-api-v2/internal/shared/error"
)

type UseCase = signin

type signin interface {
	Execute(input *Input) (*Output, *shared_error.Error)
}
