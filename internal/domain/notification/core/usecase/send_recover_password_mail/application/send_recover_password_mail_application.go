package send_recover_password_mail_application

import (
	"encoding/json"
	"errors"
	contract "getfund-api-v2/internal/domain/notification/core/contract"
	"getfund-api-v2/internal/domain/notification/core/notification_dto"

	"getfund-api-v2/internal/domain/notification/core/usecase/send_recover_password_mail"
	"strings"

	"getfund-api-v2/internal/shared/contract/settings"
	"getfund-api-v2/internal/shared/result_app"
	"getfund-api-v2/internal/shared/service/cache_service"
)

type sendRecoverPasswordMailApplication struct {
	cacheService cache_service.Cache
	mailService  contract.MailService
	settings     settings.ApplicationSettings
	template     contract.TemplateFileService
}

func New(cacheService cache_service.Cache, mailService contract.MailService, settings settings.ApplicationSettings, template contract.TemplateFileService) send_recover_password_mail.UseCase {
	return &sendRecoverPasswordMailApplication{
		cacheService: cacheService,
		mailService:  mailService,
		settings:     settings,
		template:     template,
	}
}

func (uc *sendRecoverPasswordMailApplication) Execute(input *send_recover_password_mail.Input) (*send_recover_password_mail.Output, *result_app.ApplicationError) {
	validated := input.Validate()
	if validated.IsInvalid() {
		return nil, result_app.New(result_app.BAD_REQUEST_CODE, validated.GetErrors())
	}

	userCached, errCache := uc.cacheService.Get(input.KeyCache)
	if errCache != nil {
		return nil, result_app.New(result_app.SERVER_ERROR_CODE, errCache)
	}

	userToRecoverPasswordMailModel := &notification_dto.RecoverPasswordMailDto{}
	errUnmarshal := json.Unmarshal([]byte(userCached), userToRecoverPasswordMailModel)
	if errUnmarshal != nil {
		return nil, result_app.New(result_app.SERVER_ERROR_CODE, errors.New("error to unmarshal data"))
	}

	recoveryPasswordTemplate, errTemplate := uc.template.GetRecoveryPasswordTemplate()
	if errTemplate != nil {
		return nil, result_app.New(result_app.SERVER_ERROR_CODE, errTemplate)
	}

	recoveryPasswordTemplate = replaceTags(recoveryPasswordTemplate, userToRecoverPasswordMailModel)

	err := uc.mailService.SendMail(
		userToRecoverPasswordMailModel.Username,
		"Password Recovery",
		recoveryPasswordTemplate, nil)

	if err != nil {
		return nil, result_app.New(result_app.SERVER_ERROR_CODE, err)
	}

	return &send_recover_password_mail.Output{Messagem: "Email sent successfully"}, nil
}

func replaceTags(template string, model *notification_dto.RecoverPasswordMailDto) string {
	template = strings.ReplaceAll(template, "{{first_name}}", model.FirstName)
	template = strings.ReplaceAll(template, "{{recovery_link}}", model.RecoveryLink)

	return template
}
