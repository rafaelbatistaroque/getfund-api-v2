package bus

// Event é a interface genérica para eventos
type Event interface {
	GetName() string
	GetPayload() []byte
	SetPayload(payload []byte)

	SetChannel(channel chan []byte)
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

func (e *EventBase) ResolvePromise(result []byte) {
	if e.channel == nil {
		return
	}

	e.channel <- result

	if len(e.channel) == cap(e.channel)-1 {
		close(e.channel)
	}
}
