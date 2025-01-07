package reset_password

import "getfund-api-v2/internal/shared/result_app"

type UseCase = resetPassword

type resetPassword interface {
	Execute(input *Input) (*Output, *result_app.ApplicationError)
}
