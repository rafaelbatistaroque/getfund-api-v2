package eventbus_spy

import (
	"getfund-api-v2/pkg/bus"
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

func (eb *EventBusSpy) Subscribe(eventName string, handler bus.Handler) {

}

func (eb *EventBusSpy) Emit(event bus.Event) {
	eb.Params["Publish:event"] = event

	eb.CallsCount["Publish"]++
}

func (eb *EventBusSpy) EmitWithPayload(event bus.Event, payload any) {
	eb.Params["EmitWithPayload:event"] = event
	eb.Params["EmitWithPayload:payload"] = payload

	eb.CallsCount["EmitWithPayload"]++
}
