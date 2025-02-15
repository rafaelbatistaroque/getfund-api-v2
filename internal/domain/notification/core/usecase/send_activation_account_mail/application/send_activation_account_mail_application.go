package send_activation_account_mail_application

import (
	notification_contract "getfund-api-v2/internal/domain/notification/core/contract"
	"getfund-api-v2/internal/domain/notification/core/usecase/send_activation_account_mail"
	"getfund-api-v2/internal/settings"
	"getfund-api-v2/internal/shared/result_app"
)

type sendActivationAccountMailApplication struct {
	mailService notification_contract.MailService
	settings    settings.ApplicationSettings
	template    notification_contract.TemplateFileService
}

func New(mailService notification_contract.MailService, settings settings.ApplicationSettings, template notification_contract.TemplateFileService) send_activation_account_mail.UseCase {
	return &sendActivationAccountMailApplication{
		mailService: mailService,
		settings:    settings,
		template:    template,
	}
}

func (s *sendActivationAccountMailApplication) Execute(input *send_activation_account_mail.Input) (*send_activation_account_mail.Output, *result_app.ApplicationError) {
	validatable := input.Validate()
	if validatable.IsInvalid() {
		return nil, result_app.New(result_app.UNPROCESSABLE_CONTENT_CODE, validatable.GetErrors())
	}

	return nil, nil
}
