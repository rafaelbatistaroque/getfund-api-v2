package send_recover_password_mail_application

import (
	"getfund-api-v2/internal/config/env"
	contract "getfund-api-v2/internal/domain/notification/core/contract"
	"getfund-api-v2/internal/domain/notification/core/usecase/send_recover_password_mail"
	shared_error "getfund-api-v2/internal/shared/error"
	"getfund-api-v2/internal/shared/mail"
	"getfund-api-v2/internal/shared/replacer"
)

type sendRecoverPasswordMailApplication struct {
	mailService mail.Contract
	env         env.Variable
	template    contract.TemplateFileContract
}

func New(mailService mail.Contract, env env.Variable, template contract.TemplateFileContract) send_recover_password_mail.UseCase {
	return &sendRecoverPasswordMailApplication{
		mailService: mailService,
		env:         env,
		template:    template,
	}
}

func (uc *sendRecoverPasswordMailApplication) Execute(input *send_recover_password_mail.Input) (*send_recover_password_mail.Output, *shared_error.Error) {
	validated := input.Validate()
	if validated.IsInvalid() {
		return nil, shared_error.New(shared_error.BAD_REQUEST_CODE, validated.GetErrors())
	}

	recoveryPasswordTemplate, errTemplate := uc.template.GetRecoveryPasswordTemplate()
	if errTemplate != nil {
		return nil, shared_error.New(shared_error.SERVER_ERROR_CODE, errTemplate)
	}

	replacer.Build(&recoveryPasswordTemplate,
		replacer.Replaceable{Tag: "{{first_name}}", Value: input.FirstName},
		replacer.Replaceable{Tag: "{{recovery_link}}", Value: input.RecoveryLink},
	)

	err := uc.mailService.SendMail(
		input.Username,
		"Password Recovery",
		recoveryPasswordTemplate, nil)

	if err != nil {
		return nil, shared_error.New(shared_error.SERVER_ERROR_CODE, err)
	}

	return &send_recover_password_mail.Output{Messagem: "Email sent successfully"}, nil
}
