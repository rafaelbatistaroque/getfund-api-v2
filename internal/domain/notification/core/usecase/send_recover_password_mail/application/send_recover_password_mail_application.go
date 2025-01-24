package send_recover_password_mail_application

import (
	contract "getfund-api-v2/internal/domain/notification/core/contract"

	"getfund-api-v2/internal/domain/notification/core/usecase/send_recover_password_mail"
	"strings"

	"getfund-api-v2/internal/shared/contract/settings"
	"getfund-api-v2/internal/shared/result_app"
)

type sendRecoverPasswordMailApplication struct {
	mailService contract.MailService
	settings    settings.ApplicationSettings
	template    contract.TemplateFileService
}

func New(mailService contract.MailService, settings settings.ApplicationSettings, template contract.TemplateFileService) send_recover_password_mail.UseCase {
	return &sendRecoverPasswordMailApplication{
		mailService: mailService,
		settings:    settings,
		template:    template,
	}
}

func (uc *sendRecoverPasswordMailApplication) Execute(input *send_recover_password_mail.Input) (*send_recover_password_mail.Output, *result_app.ApplicationError) {
	validated := input.Validate()
	if validated.IsInvalid() {
		return nil, result_app.New(result_app.BAD_REQUEST_CODE, validated.GetErrors())
	}

	recoveryPasswordTemplate, errTemplate := uc.template.GetRecoveryPasswordTemplate()
	if errTemplate != nil {
		return nil, result_app.New(result_app.SERVER_ERROR_CODE, errTemplate)
	}

	recoveryPasswordTemplate = replaceTags(recoveryPasswordTemplate, input)

	err := uc.mailService.SendMail(
		input.Username,
		"Password Recovery",
		recoveryPasswordTemplate, nil)

	if err != nil {
		return nil, result_app.New(result_app.SERVER_ERROR_CODE, err)
	}

	return &send_recover_password_mail.Output{Messagem: "Email sent successfully"}, nil
}

func replaceTags(template string, model *send_recover_password_mail.Input) string {
	template = strings.ReplaceAll(template, "{{first_name}}", model.FirstName)
	template = strings.ReplaceAll(template, "{{recovery_link}}", model.RecoveryLink)

	return template
}
