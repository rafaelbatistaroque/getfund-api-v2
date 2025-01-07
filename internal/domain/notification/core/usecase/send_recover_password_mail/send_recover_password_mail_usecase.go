package send_recover_password_mail

import "getfund-api-v2/internal/shared/result_app"

type UseCase = sendRecoverPasswordMail

type sendRecoverPasswordMail interface {
	Execute(input *Input) (*Output, *result_app.ApplicationError)
}
