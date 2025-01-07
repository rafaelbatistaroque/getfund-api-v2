package recover_password

import "getfund-api-v2/internal/shared/result_app"

type UseCase = recoverPassword

type recoverPassword interface {
	Execute(input *Input) (*Output, *result_app.ApplicationError)
}
