package send_recover_password_mail_application

import (
	"encoding/json"
	"errors"
	template_file "getfund-api-v2/internal/domain/notification/adapter/contract"
	"getfund-api-v2/internal/domain/notification/adapter/domain_service/mail_service"
	notification_model "getfund-api-v2/internal/domain/notification/adapter/model"
	"getfund-api-v2/internal/domain/notification/adapter/usecase/send_recover_password_mail"
	"strings"

	"getfund-api-v2/internal/shared/contract/settings"
	"getfund-api-v2/internal/shared/result_app"
	"getfund-api-v2/internal/shared/service/cache_service"
)

type sendRecoverPasswordMailApplication struct {
	cacheService cache_service.Cache
	mailService  mail_service.MailService
	settings     settings.ApplicationSettings
	template     template_file.TemplateFile
	//template service
}

func New(cacheService cache_service.Cache, mailService mail_service.MailService, settings settings.ApplicationSettings, template template_file.TemplateFile) send_recover_password_mail.UseCase {
	return &sendRecoverPasswordMailApplication{
		cacheService: cacheService,
		mailService:  mailService,
		settings:     settings,
		template:     template,
	}
}

func (uc *sendRecoverPasswordMailApplication) Execute(input *send_recover_password_mail.Input) (*send_recover_password_mail.Output, *result_app.ApplicationError) {
	input.Validate()
	if input.IsInvalid() {
		return nil, result_app.New(result_app.BAD_REQUEST_CODE, input.GetErrors())
	}

	userCached, errCache := uc.cacheService.Get(input.KeyCache)
	if errCache != nil {
		return nil, result_app.New(result_app.SERVER_ERROR_CODE, errCache)
	}

	userToRecoverPasswordMailModel := &notification_model.RecoverPasswordMailModel{}
	errUnmarshal := json.Unmarshal([]byte(userCached), userToRecoverPasswordMailModel)
	if errUnmarshal != nil {
		return nil, result_app.New(result_app.SERVER_ERROR_CODE, errors.New("error to unmarshal data"))
	}

	recoveryPasswordTemplate, errTemplate := uc.template.GetRecoveryPasswordTemplate()
	if errTemplate != nil {
		return nil, result_app.New(result_app.SERVER_ERROR_CODE, errTemplate)

	}

	recoveryPasswordTemplate = strings.ReplaceAll(recoveryPasswordTemplate, "{{first_name}}", userToRecoverPasswordMailModel.FirstName)
	recoveryPasswordTemplate = strings.ReplaceAll(recoveryPasswordTemplate, "{{recovery_link}}", userToRecoverPasswordMailModel.RecoveryLink)

	err := uc.mailService.SendMail(
		uc.settings.GetSMTPFrom(),
		userToRecoverPasswordMailModel.Username,
		"Password Recovery",
		recoveryPasswordTemplate, nil)

	if err != nil {
		return nil, result_app.New(result_app.SERVER_ERROR_CODE, err)
	}

	return &send_recover_password_mail.Output{Messagem: "Email sent successfully"}, nil
}
