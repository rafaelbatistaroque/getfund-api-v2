package eventbus

import "sync"

// Event é a interface genérica para eventos
type Event interface {
	GetName() string
}

type Handler interface {
	Handle(event Event)
}

type EventBus interface {
	Subscribe(eventName string, handler Handler)
	Publish(event Event)
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

	// Registra o handler no map de handlers
	eb.handlers[eventName] = append(eb.handlers[eventName], handler)
}

// Publish dispara um evento para todos os handlers associados
func (eb *eventBus) Publish(event Event) {
	eb.lock.RLock()
	defer eb.lock.RUnlock()

	if handlers, ok := eb.handlers[event.GetName()]; ok {
		for _, handle := range handlers {
			go handle.Handle(event)
		}
	}
}
