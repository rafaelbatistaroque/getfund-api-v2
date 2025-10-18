package reset_password

import (
	shared_error "getfund-api-v2/internal/shared/error"
)

type UseCase = resetPassword

type resetPassword interface {
	Execute(input *Input) (*Output, *shared_error.Error)
}
