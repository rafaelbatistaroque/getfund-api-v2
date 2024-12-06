package notificationcomposer

import (
	handler "getfund-api-v2/internal/domain/notification/port/eventhandler/newsletter"
	"getfund-api-v2/internal/shared/contract/settings"
	"getfund-api-v2/pkg/eventbus"
)

func SubscribeEventHandlers(settings settings.ApplicationSettings, eventBus eventbus.EventBus) {
	handlers := map[string]eventbus.Handler{
		"UserSignedEvent": handler.New(),
	}

	for eventName, handler := range handlers {
		eventBus.Subscribe(eventName, handler)
	}
}
