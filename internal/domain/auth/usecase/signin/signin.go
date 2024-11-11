package signin

import appErr "getfund-api-v2/internal/pkg/applicationerror"

type UseCase = signin

type signin interface {
	Execute(input *Input) (*Output, *appErr.ApplicationError)
}
