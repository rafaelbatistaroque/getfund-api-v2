package bus

import (
	"encoding/json"
	logger "getfund-api-v2/pkg/log"
	"sync"
)

// Event é a interface genérica para eventos
type Event interface {
	GetName() string
	GetPayload() []byte
	SetPayload(payload []byte)

	SetChannel(channel chan []byte)
	DefineResponse(response []byte)
}

type Handler interface {
	Handle(event Event)
}

type EventBus interface {
	Subscribe(eventName string, handler Handler)
	Emit(event Event)
	EmitWithPayload(event Event, payload any)
	EmitWithPayloadAndResponse(event Event, payload any, responseChannel chan []byte)
}

type eventBus struct {
	logger   *logger.Logger
	handlers map[string][]Handler // Armazena handlers
	lock     sync.RWMutex
}

func New() EventBus {
	return &eventBus{
		handlers: make(map[string][]Handler),
		logger:   logger.New("EventBus"),
	}
}

// Subscribe adiciona um handler para um tipo específico de evento
func (eb *eventBus) Subscribe(eventName string, handler Handler) {
	eb.lock.Lock()
	defer eb.lock.Unlock()

	eb.handlers[eventName] = append(eb.handlers[eventName], handler)
	eb.logger.Infof("Event %s subscribed", eventName)
}

// Emit dispara um evento para todos os handlers associados
func (eb *eventBus) Emit(event Event) {
	eb.lock.RLock()
	defer eb.lock.RUnlock()

	if handlers, ok := eb.handlers[event.GetName()]; ok {
		eb.logger.Infof("Event %v published", event.GetName())
		for _, handle := range handlers {
			go handle.Handle(event)
		}
	}
}

// EmitWithPayload após incluir o payload ao evento dispara para todos os handlers
func (eb *eventBus) EmitWithPayload(event Event, payload any) {
	var data []byte
	var err error

	switch typeData := payload.(type) {
	case string:
		data = []byte(typeData)
	default:
		data, err = json.Marshal(payload)
		if err != nil {
			eb.logger.Errorf("Error on serialize the payload: %v", err)
		}
	}

	event.SetPayload(data)

	eb.Emit(event)
}

func (eb *eventBus) EmitWithPayloadAndResponse(event Event, payload any, responseChannel chan []byte) {
	event.SetChannel(responseChannel)
	eb.EmitWithPayload(event, payload)
}
