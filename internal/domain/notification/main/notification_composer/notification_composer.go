package notification_composer

import (
	recover_password_eventhandler "getfund-api-v2/internal/domain/notification/port/event_handler/recover_password"
	"getfund-api-v2/internal/shared/contract/settings"
	"getfund-api-v2/internal/shared/service/cache_service"
	"getfund-api-v2/pkg/eventbus"
)

func SubscribeEventHandlers(settings settings.ApplicationSettings, eventBus eventbus.EventBus, cacheService cache_service.Cache) {
	handlers := map[string]eventbus.Handler{
		"RecoverPasswordStarted": recover_password_eventhandler.New(cacheService),
	}

	for eventName, handler := range handlers {
		eventBus.Subscribe(eventName, handler)
	}
}
