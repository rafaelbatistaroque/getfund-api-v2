package recoverpasswordeventhandler

import (
	"encoding/json"
	"getfund-api-v2/pkg/eventbus"
)

type recoverPasswordEventHandler struct {
	//usecase para acao
}

func New() eventbus.Handler {
	return &recoverPasswordEventHandler{}
}

// Implementa o Handler genérico para RecoverPasswordStartedEvent
func (h *recoverPasswordEventHandler) Handle(event eventbus.Event) {
	var payload string
	json.Unmarshal(event.GetPayload(), &payload)
	//Handler
	//TODO: recover data cached by key received
	//TODO: build a email template with params to replace
	//TODO: replace specific
	//TODO: send email
}
