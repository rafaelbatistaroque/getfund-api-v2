package newslettereventhandler

import (
	"encoding/json"
	"fmt"
	"getfund-api-v2/internal/pkg/eventbus"
)

type newsletterPayload struct {
	Id        string
	FirstName string `json:"first_name"`
}

type newsletterEventHandler struct {
	//usecase para acao
}

func New() eventbus.Handler {
	return &newsletterEventHandler{}
}

// Implementa o Handler genérico para SignedEvent
func (h *newsletterEventHandler) Handle(event eventbus.Event) {
	payload := &newsletterPayload{}
	json.Unmarshal(event.GetPayload(), payload)
	fmt.Printf("Processando newsletter para SignedEvent com ID: %s\n", payload.FirstName)
}
