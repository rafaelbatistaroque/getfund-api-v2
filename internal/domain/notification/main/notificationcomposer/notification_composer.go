package notificationcomposer

import (
	recoverpasswordeventhandler "getfund-api-v2/internal/domain/notification/port/eventhandler/recoverpassword"
	"getfund-api-v2/internal/shared/contract/settings"
	"getfund-api-v2/pkg/eventbus"
)

func SubscribeEventHandlers(settings settings.ApplicationSettings, eventBus eventbus.EventBus) {
	handlers := map[string]eventbus.Handler{
		"RecoverPasswordStartedEvent": recoverpasswordeventhandler.New(),
	}

	for eventName, handler := range handlers {
		eventBus.Subscribe(eventName, handler)
	}
}
