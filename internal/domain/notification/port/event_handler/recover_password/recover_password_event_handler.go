package recover_password_event_handler

import (
	"encoding/json"
	"fmt"
	"getfund-api-v2/internal/shared/service/cache_service"
	"getfund-api-v2/pkg/eventbus"
)

type recoverPasswordEventHandler struct {
	cacheService cache_service.Cache
}

func New(cacheService cache_service.Cache) eventbus.Handler {
	return &recoverPasswordEventHandler{
		cacheService: cacheService,
	}
}

// Implementa o Handler genérico para RecoverPasswordStartedEvent
func (h *recoverPasswordEventHandler) Handle(event eventbus.Event) {
	var keyCache string
	json.Unmarshal(event.GetPayload(), &keyCache)

	result, _ := h.cacheService.Get(keyCache)

	fmt.Printf("teste %v", result)
	//Handler
	//TODO: recover data cached by key received
	//TODO: build a email template with params to replace
	//TODO: replace specific
	//TODO: send email
}
