package notificationcomposer

import (
	handler "getfund-api-v2/internal/domain/notification/port/eventhandler"
	"getfund-api-v2/internal/pkg/eventbus"
	"getfund-api-v2/internal/shared/contract/settings"

	"net/http"
)

type NotificationAdapter struct {
	Signin http.HandlerFunc
}

func SubscribeEventHandlers(settings settings.ApplicationSettings, eventBus eventbus.EventBus) {

	eventBus.Subscribe("SignedEvent", handler.New())
}
