package send_activation_account_mail

import "getfund-api-v2/internal/shared/result_app"

type UseCase = sendActivationAccountMail

type sendActivationAccountMail interface {
	Execute(input *Input) (*Output, *result_app.ApplicationError)
}
