package notification_composer

import (
	"getfund-api-v2/internal/domain/notification/port/event_handler/recover_password_event_handler"
	"getfund-api-v2/internal/shared/contract/settings"
	"getfund-api-v2/pkg/eventbus"
)

func SubscribeEventHandlers(settings settings.ApplicationSettings, eventBus eventbus.EventBus) {
	handlers := map[string]eventbus.Handler{
		"RecoverPasswordStarted": recover_password_event_handler.New(nil),
	}

	for eventName, handler := range handlers {
		eventBus.Subscribe(eventName, handler)
	}
}
