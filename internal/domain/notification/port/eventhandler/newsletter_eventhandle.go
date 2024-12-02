package newslettereventhandle

import (
	"fmt"
	"getfund-api-v2/internal/pkg/eventbus"
)

type newsletterEventHandler struct {
	//usecase para acao
}

func New() eventbus.Handler {
	return &newsletterEventHandler{}
}

// Implementa o Handler genérico para SignedEvent
func (h *newsletterEventHandler) Handle(event eventbus.Event) {
	fmt.Printf("Processando newsletter para SignedEvent com ID: %s\n", event)
}
