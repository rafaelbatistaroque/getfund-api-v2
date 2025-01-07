package newsletter_event_handler

import (
	"encoding/json"
	"fmt"
	"getfund-api-v2/pkg/bus"
)

type newsletterPayload struct {
	Id        string
	FirstName string `json:"first_name"`
}

type newsletterEventHandler struct {
	//usecase para acao
}

func New() bus.Handler {
	return &newsletterEventHandler{}
}

// Implementa o Handler genérico para SignedEvent
func (h *newsletterEventHandler) Handle(event bus.Event) {
	payload := &newsletterPayload{}
	json.Unmarshal(event.GetPayload(), payload)
	fmt.Printf("Processando newsletter para SignedEvent com ID: %s\n", payload.FirstName)
}
