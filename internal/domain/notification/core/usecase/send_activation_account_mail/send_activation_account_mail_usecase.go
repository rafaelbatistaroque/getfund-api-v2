package send_activation_account_mail

type UseCase = sendActivationAccountMail

type sendActivationAccountMail interface {
	Execute(input *Input) (*Output, error)
}
