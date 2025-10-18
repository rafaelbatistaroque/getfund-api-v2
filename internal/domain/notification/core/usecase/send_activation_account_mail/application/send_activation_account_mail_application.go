package send_activation_account_mail_application

import (
	"getfund-api-v2/internal/config/env"
	notification_contract "getfund-api-v2/internal/domain/notification/core/contract"
	"getfund-api-v2/internal/domain/notification/core/usecase/send_activation_account_mail"
	shared_error "getfund-api-v2/internal/shared/error"
	"getfund-api-v2/internal/shared/mail"
	"getfund-api-v2/internal/shared/replacer"
)

type sendActivationAccountMailApplication struct {
	mail     mail.Contract
	env      env.Variable
	template notification_contract.TemplateFileContract
}

func New(mail mail.Contract, env env.Variable, template notification_contract.TemplateFileContract) send_activation_account_mail.UseCase {
	return &sendActivationAccountMailApplication{
		mail:     mail,
		env:      env,
		template: template,
	}
}

func (s *sendActivationAccountMailApplication) Execute(input *send_activation_account_mail.Input) (*send_activation_account_mail.Output, *shared_error.Error) {
	validatable := input.Validate()
	if validatable.IsInvalid() {
		return nil, shared_error.New(shared_error.UNPROCESSABLE_CONTENT_CODE, validatable.GetErrors())
	}

	activationAccountTemplate, errTemplate := s.template.GetActivationAccountTemplate()
	if errTemplate != nil {
		return nil, shared_error.New(shared_error.SERVER_ERROR_CODE, errTemplate)
	}

	replacer.Build(&activationAccountTemplate,
		replacer.Replaceable{Tag: "{{first_name}}", Value: input.FirstName},
		replacer.Replaceable{Tag: "{{activation_link}}", Value: input.ActivationLink},
	)

	if err := s.mail.SendMail(
		input.Email,
		"Activation Account",
		activationAccountTemplate,
		nil); err != nil {
		return nil, shared_error.New(shared_error.SERVER_ERROR_CODE, err)
	}

	return &send_activation_account_mail.Output{Message: "Email sent successfully"}, nil
}
