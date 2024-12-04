package recoverpassword

import "getfund-api-v2/internal/shared/resultapp"

type UseCase = recoverPassword

type recoverPassword interface {
	Execute(input *Input) (*Output, *resultapp.ApplicationError)
}
