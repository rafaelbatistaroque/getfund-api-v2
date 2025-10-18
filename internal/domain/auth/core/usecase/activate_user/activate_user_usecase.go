package activate_user

import (
	shared_error "getfund-api-v2/internal/shared/error"
)

type UseCase = activateUser

type activateUser interface {
	Execute(input *Input) (*Output, *shared_error.Error)
}
