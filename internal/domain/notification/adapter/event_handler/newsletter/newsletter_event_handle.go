package newsletter_event_handler

import (
	"encoding/json"
	"fmt"
	shared_bus "getfund-api-v2/internal/shared/bus"
)

type newsletterPayload struct {
	Id        string
	FirstName string `json:"first_name"`
}

type newsletterEventHandler struct {
	//usecase para acao
}

func New() shared_bus.Handler {
	return &newsletterEventHandler{}
}

// Implementa o Handler genérico para SignedEvent
func (h *newsletterEventHandler) Handle(event shared_bus.Event) {
	payload := &newsletterPayload{}
	json.Unmarshal(event.GetPayload(), payload)
	fmt.Printf("Processando newsletter para SignedEvent com ID: %s\n", payload.FirstName)
}
