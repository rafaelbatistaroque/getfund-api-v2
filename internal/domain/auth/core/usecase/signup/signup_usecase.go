package signup

import (
	shared_error "getfund-api-v2/internal/shared/error"
)

type UseCase = Signup

type Signup interface {
	Execute(input *Input) (*Output, *shared_error.Error)
}
