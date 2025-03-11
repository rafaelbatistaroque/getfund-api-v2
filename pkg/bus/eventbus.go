package bus

import (
	"encoding/json"
	"errors"
	logger "getfund-api-v2/pkg/log"
	"reflect"
	"strconv"
	"sync"
	"time"
)

var instance *eventBus
var once sync.Once

type EventBus interface {
	Subscribe(eventName string, handler Handler)
	Emit(event Event)
	EmitWithPayload(event Event, payload any)
	EmitWithPayloadAndResponse(event Event, payload any, responseChannel chan []byte)
	EmitWithPromise(event Event, payload any) *Promise
	EmitAndWaitPromise(event Event, payload any, result any) *Promise
	Wait(promise *Promise, result any)
}

type eventBus struct {
	logger   *logger.Logger
	handlers map[string][]Handler // Armazena handlers
	lock     sync.RWMutex
	timeout  time.Duration
}

func New(timeoutInSecond int) EventBus {
	once.Do(func() {
		instance = &eventBus{
			handlers: make(map[string][]Handler),
			logger:   logger.New("EventBus"),
			timeout:  time.Duration(timeoutInSecond) * time.Second,
		}
	})

	return instance
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
	event.SetPayload(toBytes(payload, eb.logger))

	eb.Emit(event)
}

func (eb *eventBus) EmitWithPayloadAndResponse(event Event, payload any, responseChannel chan []byte) {
	event.SetChannel(responseChannel)

	eb.EmitWithPayload(event, payload)
}

func (eb *eventBus) EmitWithPromise(event Event, payload any) *Promise {
	resultChannel := make(chan []byte, 1)
	promise := &Promise{}
	promise.SetChannel(resultChannel)

	eb.EmitWithPayloadAndResponse(event, payload, resultChannel)

	return promise
}

func (eb *eventBus) EmitAndWaitPromise(event Event, payload any, result any) *Promise {
	promise := eb.EmitWithPromise(event, payload)
	eb.Wait(promise, &result)

	return promise
}

func (eb *eventBus) Wait(promise *Promise, result any) {
	rs := reflect.ValueOf(result)
	if rs.Kind() != reflect.Pointer || rs.IsNil() {
		panic(errors.New("result must be a pointer"))
	}

	select {
	case rawResult := <-promise.GetChannel():
		if len(rawResult) == 0 {
			promise.SetError(errors.New("empty response"))
			return
		}

		if err := fromByte(rawResult, &result, eb.logger); err != nil {
			promise.SetError(err)
			return
		}

		if result == nil {
			promise.SetError(errors.New("result null"))
		}

		result = &rawResult
	case <-time.After(eb.timeout):
		promise.SetError(errors.New("timeout waiting for event"))
	}
}

func toBytes(payload any, logger *logger.Logger) []byte {
	var data []byte
	var err error

	switch typeData := payload.(type) {
	case string:
		data = []byte(typeData)
	default:
		data, err = json.Marshal(payload)
		if err != nil {
			logger.Errorf("Error on serialize the payload: %v", err)
			return nil
		}
	}
	return data
}

func fromByte(rawResult []byte, result *any, logger *logger.Logger) error {
	switch (*result).(type) {
	case string:
		*result = string(rawResult)
	case int:
		ok, err := strconv.Atoi(string(rawResult))
		if err != nil {
			return errors.New("invalid result")
		}
		*result = ok
	case bool:
		*result = string(rawResult) == "true"
	default:
		var tempResult any
		err := json.Unmarshal(rawResult, &tempResult)
		*result = tempResult
		if err != nil {
			logger.Errorf("Error on deserialize the payload: %v", err)
			return errors.New("invalid result")
		}
	}

	return nil
}
