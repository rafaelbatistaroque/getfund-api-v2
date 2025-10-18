package shared_bus

// Event é a interface genérica para eventos
type Event interface {
	// GetName returns the name of the event.
	GetName() string
	// GetPayload returns the payload of the event.
	GetPayload() []byte
	// SetPayload sets the payload of the event.
	SetPayload(payload []byte)

	// SetChannel sets the channel for the event promise.
	SetChannel(channel chan []byte)
	// ResolvePromise resolves the promise with a result.
	ResolvePromise(result []byte)
}

type EventBase struct {
	payload []byte
	channel chan []byte
}

func (e *EventBase) GetPayload() []byte {
	return e.payload
}

func (e *EventBase) SetPayload(payload []byte) {
	e.payload = payload
}

func (e *EventBase) SetChannel(channel chan []byte) {
	e.channel = channel
}

func (e *EventBase) GetChannel() chan []byte {
	return e.channel
}

func (e *EventBase) ResolvePromise(result []byte) {
	if e.channel == nil {
		return
	}

	e.channel <- result

	if len(e.channel) == cap(e.channel)-1 {
		close(e.channel)
	}
}
