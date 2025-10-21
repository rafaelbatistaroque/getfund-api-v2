package composer

import (
	"getfund-api-v2/internal/config/env"
	config_mail "getfund-api-v2/internal/config/go_mail"
	"getfund-api-v2/internal/domain/notification/adapter/event_handler/recover_password_started_event_handler"
	signup_process_started_event_handler "getfund-api-v2/internal/domain/notification/adapter/event_handler/signup_started_event_handler"
	"getfund-api-v2/internal/domain/notification/adapter/service/template_file_service"
	send_activation_account_mail_application "getfund-api-v2/internal/domain/notification/core/usecase/send_activation_account_mail/application"
	send_recover_password_mail_application "getfund-api-v2/internal/domain/notification/core/usecase/send_recover_password_mail/application"

	shared_bus "getfund-api-v2/internal/shared/bus"
	"getfund-api-v2/internal/shared/cache"
)

func Compose(env env.Variable, eventBus shared_bus.EventBus, cacheService cache.Service) {

	//Services
	mailService := config_mail.New(env)
	templateFileService := template_file_service.New(env)

	//Applications
	sendRecoverPasswordMailApplication := send_recover_password_mail_application.New(mailService, env, templateFileService)
	sendActivationAccountMailApplication := send_activation_account_mail_application.New(mailService, env, templateFileService)

	//Event Handler
	handlers := map[string]shared_bus.Handler{
		"recover.password.started": recover_password_started_event_handler.New(sendRecoverPasswordMailApplication, cacheService),
		"signup.started":           signup_process_started_event_handler.New(sendActivationAccountMailApplication),
	}

	for eventName, handler := range handlers {
		eventBus.Subscribe(eventName, handler)
	}
}
