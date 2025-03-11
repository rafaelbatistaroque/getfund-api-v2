package notification_composer

import (
	"getfund-api-v2/internal/domain/notification/adapter/event_handler/create_user_process_started_event_handler"
	"getfund-api-v2/internal/domain/notification/adapter/event_handler/recover_password_started_event_handler"
	"getfund-api-v2/internal/domain/notification/adapter/service/mail_service"
	"getfund-api-v2/internal/domain/notification/adapter/service/template_file_service"
	send_activation_account_mail_application "getfund-api-v2/internal/domain/notification/core/usecase/send_activation_account_mail/application"
	send_recover_password_mail_application "getfund-api-v2/internal/domain/notification/core/usecase/send_recover_password_mail/application"
	"getfund-api-v2/internal/settings"
	"getfund-api-v2/internal/shared/service/cache_service"
	"getfund-api-v2/pkg/bus"
	"getfund-api-v2/pkg/mail"
)

func Compose(settings settings.ApplicationSettings, eventBus bus.EventBus, cacheService cache_service.Cache) {

	//Services
	mailService := mail_service.New(mail.New(settings))
	templateFileService := template_file_service.New(settings)

	//Applications
	sendRecoverPasswordMailApplication := send_recover_password_mail_application.New(mailService, settings, templateFileService)
	sendActivationAccountMailApplication := send_activation_account_mail_application.New(mailService, settings, templateFileService)

	//Event Handler
	handlers := map[string]bus.Handler{
		"RecoverPasswordStartedEvent":   recover_password_started_event_handler.New(sendRecoverPasswordMailApplication, cacheService),
		"CreateUserProcessStartedEvent": create_user_process_started_event_handler.New(cacheService, sendActivationAccountMailApplication),
	}

	for eventName, handler := range handlers {
		eventBus.Subscribe(eventName, handler)
	}
}
