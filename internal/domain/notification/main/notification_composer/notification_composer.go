package notification_composer

import (
	send_recover_password_mail_application "getfund-api-v2/internal/domain/notification/adapter/usecase/send_recover_password_mail/application"
	"getfund-api-v2/internal/domain/notification/port/event_handler/recover_password_started_event_handler"
	"getfund-api-v2/internal/domain/notification/port/service/mail_service"
	"getfund-api-v2/internal/domain/notification/port/service/template_file_service"
	"getfund-api-v2/internal/shared/contract/settings"
	"getfund-api-v2/internal/shared/service/cache_service"
	"getfund-api-v2/pkg/bus"
	"getfund-api-v2/pkg/mail"
)

func SubscribeEventHandlers(settings settings.ApplicationSettings, eventBus bus.EventBus, cacheService cache_service.Cache) {
	handlers := map[string]bus.Handler{
		"RecoverPasswordStarted": recover_password_started_event_handler.New(
			send_recover_password_mail_application.New(
				cacheService,
				mail_service.New(mail.New(settings)),
				settings,
				template_file_service.New(settings))),
	}

	for eventName, handler := range handlers {
		eventBus.Subscribe(eventName, handler)
	}
}
