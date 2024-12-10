package eventbusspy

import (
	"getfund-api-v2/pkg/eventbus"
)

type EventBusSpy struct {
	Params        map[string]interface{}
	CallsCount    map[string]int
	ErrorResult   map[string]error
	SuccessResult map[string]interface{}
}

func New() *EventBusSpy {
	return &EventBusSpy{Params: make(map[string]interface{}), CallsCount: make(map[string]int), ErrorResult: make(map[string]error), SuccessResult: make(map[string]interface{})}
}

func (eb *EventBusSpy) Subscribe(eventName string, handler eventbus.Handler) {

}

// Publish dispara um evento para todos os handlers associados
func (eb *EventBusSpy) Publish(event eventbus.Event) {

}

// CreateAndPublish após incluir o payload ao evento dispara para todos os handlers
func (eb *EventBusSpy) CreateAndPublish(event eventbus.Event, payload any) {
	eb.Params["CreateAndPublish:event"] = event
	eb.Params["CreateAndPublish:payload"] = payload

	eb.CallsCount["CreateAndPublish"]++
}
