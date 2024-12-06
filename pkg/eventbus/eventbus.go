package eventbus

import (
	"encoding/json"
	applog "getfund-api-v2/pkg/log"
	"sync"
)

// Event é a interface genérica para eventos
type Event interface {
	GetName() string
	GetPayload() []byte
	SetPayload(payload []byte)
}

type Handler interface {
	Handle(event Event)
}

type EventBus interface {
	Subscribe(eventName string, handler Handler)
	Publish(event Event)
	CreateAndPublish(event Event, payload any)
}

type eventBus struct {
	handlers map[string][]Handler // Armazena handlers
	lock     sync.RWMutex
}

func New() EventBus {
	return &eventBus{
		handlers: make(map[string][]Handler),
	}
}

// Subscribe adiciona um handler para um tipo específico de evento
func (eb *eventBus) Subscribe(eventName string, handler Handler) {
	eb.lock.Lock()
	defer eb.lock.Unlock()

	eb.handlers[eventName] = append(eb.handlers[eventName], handler)
	applog.Info.Printf("Event subscribed %v", eventName)
}

// Publish dispara um evento para todos os handlers associados
func (eb *eventBus) Publish(event Event) {
	eb.lock.RLock()
	defer eb.lock.RUnlock()

	if handlers, ok := eb.handlers[event.GetName()]; ok {
		applog.Info.Printf("Event published %v", event.GetName())
		for _, handle := range handlers {
			go handle.Handle(event)
		}
	}
}

// CreateAndPublish após incluir o payload ao evento dispara para todos os handlers
func (eb *eventBus) CreateAndPublish(event Event, payload any) {
	data, err := json.Marshal(payload)
	if err != nil {
		applog.Info.Printf("error on serialize the payload: %v", err)
	}

	event.SetPayload(data)

	eb.Publish(event)
}
