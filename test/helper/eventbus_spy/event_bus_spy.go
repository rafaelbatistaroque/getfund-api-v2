package eventbus_spy

import (
	"getfund-api-v2/pkg/bus"
	"time"
)

type EventBusSpy struct {
	Params        map[string][]any
	CallsCount    map[string]int
	ErrorResult   map[string]error
	SuccessResult map[string]any
}

func New() *EventBusSpy {
	return &EventBusSpy{Params: make(map[string][]any), CallsCount: make(map[string]int), ErrorResult: make(map[string]error), SuccessResult: make(map[string]any)}
}

func (eb *EventBusSpy) Subscribe(eventName string, handler bus.Handler) {

}

func (eb *EventBusSpy) Emit(event bus.Event) {
	eb.Params["Publish:event"] = append(eb.Params["Publish:event"], event)

	eb.CallsCount["Publish"]++
}

func (eb *EventBusSpy) EmitWithPayload(event bus.Event, payload any) {
	eb.Params["EmitWithPayload:event"] = append(eb.Params["EmitWithPayload:event"], event)
	eb.Params["EmitWithPayload:payload"] = append(eb.Params["EmitWithPayload:payload"], payload)

	eb.CallsCount["EmitWithPayload"]++
}

func (eb *EventBusSpy) EmitWithPayloadAndResponse(event bus.Event, payload any, responseChannel chan []byte) {
	eb.Params["EmitWithPayloadAndResponse:event"] = append(eb.Params["EmitWithPayloadAndResponse:event"], event)
	eb.Params["EmitWithPayloadAndResponse:payload"] = append(eb.Params["EmitWithPayloadAndResponse:payload"], payload)
	eb.Params["EmitWithPayloadAndResponse:responseChannel"] = append(eb.Params["EmitWithPayloadAndResponse:responseChannel"], responseChannel)

	eb.CallsCount["EmitWithPayloadAndResponse"]++
}

func (f *EventBusSpy) Run(sut func(), responses ...[]byte) {
	go sut()

	time.Sleep(100 * time.Millisecond)
	responseChannel := f.Params["EmitWithPayloadAndResponse:responseChannel"][0].(chan []byte)
	for _, response := range responses {
		responseChannel <- response
	}

	defer close(responseChannel)
}
