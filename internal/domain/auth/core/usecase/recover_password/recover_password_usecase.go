package recover_password

import (
	shared_error "getfund-api-v2/internal/shared/error"
)

type UseCase = recoverPassword

type recoverPassword interface {
	Execute(input *Input) (*Output, *shared_error.Error)
}
