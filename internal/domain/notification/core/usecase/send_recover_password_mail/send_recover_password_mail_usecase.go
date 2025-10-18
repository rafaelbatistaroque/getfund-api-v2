package send_recover_password_mail

import (
	shared_error "getfund-api-v2/internal/shared/error"
)

type UseCase = sendRecoverPasswordMail

type sendRecoverPasswordMail interface {
	Execute(input *Input) (*Output, *shared_error.Error)
}
