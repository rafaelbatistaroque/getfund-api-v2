package shared_bus

type EventBus interface {
	// Subscribe subscribes a handler to an event.
	Subscribe(eventName string, handler Handler)
	// Emit emits an event.
	Emit(event Event)
	// EmitWithPayload emits an event with a payload.
	EmitWithPayload(event Event, payload any)
	// EmitWithPayloadAndResponse emits an event with a payload and a response channel.
	EmitWithPayloadAndResponse(event Event, payload any, responseChannel chan []byte)
	// EmitWithPromise emits an event and returns a promise.
	EmitWithPromise(event Event, payload any) *Promise
	// EmitAndWaitPromise emits an event and waits for the promise to be resolved.
	EmitAndWaitPromise(event Event, payload any, result any) *Promise
	// Wait waits for a promise to be resolved.
	Wait(promise *Promise, result any)
}
