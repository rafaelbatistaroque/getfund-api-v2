package eventbus_spy

import (
	"encoding/json"
	"errors"
	"getfund-api-v2/pkg/bus"
	"reflect"
	"strconv"
	"time"
)

type EventBusSpy struct {
	Params        map[string][]any
	CallsCount    map[string]int
	ErrorResult   map[string]error
	SuccessResult map[string]any

	ReferenceResult map[string]any
	timeout         time.Duration
}

func New() *EventBusSpy {
	return &EventBusSpy{
		Params:          make(map[string][]any),
		CallsCount:      make(map[string]int),
		ErrorResult:     make(map[string]error),
		SuccessResult:   make(map[string]any),
		ReferenceResult: make(map[string]any),
	}
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

func (eb *EventBusSpy) Wait(promise *bus.Promise, result any) {
	eb.Params["Wait:promise"] = append(eb.Params["Wait:promise"], promise)
	eb.Params["Wait:result"] = append(eb.Params["Wait:result"], result)

	eb.CallsCount["Wait"]++

	select {
	case rawResult := <-promise.GetChannel():
		if len(rawResult) == 0 {
			promise.SetError(errors.New("empty response"))
			return
		}

		if err := fromByte(rawResult, &result); err != nil {
			promise.SetError(err)
			return
		}

		if result == nil {
			promise.SetError(errors.New("result null"))
		}
	case <-time.After(eb.timeout):
		promise.SetError(errors.New("timeout waiting for event"))
	}
}

func fromByte(rawResult []byte, result any) error {
	switch result.(type) {
	case string:
		result = string(rawResult)
	case int:
		ok, err := strconv.Atoi(string(rawResult))
		if err != nil {
			return errors.New("invalid result")
		}
		result = ok
	default:
		err := json.Unmarshal(rawResult, &result)
		if err != nil {
			return errors.New("invalid result")
		}
	}

	return nil
}

func (eb *EventBusSpy) EmitWithPromise(event bus.Event, payload any) *bus.Promise {
	eb.Params["EmitWithPromise:event"] = append(eb.Params["EmitWithPromise:event"], event)
	eb.Params["EmitWithPromise:payload"] = append(eb.Params["EmitWithPromise:payload"], payload)

	eb.CallsCount["EmitWithPromise"]++

	resultChannel := make(chan []byte, 1)
	eb.EmitWithPayloadAndResponse(event, payload, resultChannel)

	promise := &bus.Promise{}
	eb.SuccessResult["EmitWithPromise"] = &bus.Promise{}

	return promise
}

func (eb *EventBusSpy) EmitAndWaitPromise(event bus.Event, payload any, result any) *bus.Promise {
	eb.Params["EmitAndWaitPromise:event"] = append(eb.Params["EmitAndWaitPromise:event"], event)
	eb.Params["EmitAndWaitPromise:payload"] = append(eb.Params["EmitAndWaitPromise:payload"], payload)
	eb.Params["EmitAndWaitPromise:result"] = append(eb.Params["EmitAndWaitPromise:result"], result)

	eb.CallsCount["EmitAndWaitPromise"]++

	if eb.ReferenceResult["EmitAndWaitPromise"+event.GetName()] != nil {
		newValue := eb.ReferenceResult["EmitAndWaitPromise"+event.GetName()]
		reflect.ValueOf(result).Elem().Set(reflect.ValueOf(newValue))
	}

	success := eb.SuccessResult["EmitAndWaitPromise"]
	if success == nil {
		return nil
	}

	return eb.SuccessResult["EmitAndWaitPromise"].(*bus.Promise)
}

func (eb *EventBusSpy) DefineEmitAndWaitPromiseError() {
	promise := &bus.Promise{}
	promise.SetError(errors.New("error"))
	eb.ErrorResult["EmitAndWaitPromise"] = errors.New("error")
	eb.SuccessResult["EmitAndWaitPromise"] = promise
}

func (eb *EventBusSpy) DefineEmitAndWaitPromiseErrorNull() {
	eb.SuccessResult["EmitAndWaitPromise"] = &bus.Promise{}
}

func (eb *EventBusSpy) DefinePromiseResult(event bus.Event, result any) {
	eb.SuccessResult["EmitAndWaitPromise"] = &bus.Promise{}
	eb.ReferenceResult["EmitAndWaitPromise"+event.GetName()] = result
}
