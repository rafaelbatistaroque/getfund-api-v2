package event

type RecoverPasswordStartedEvent struct {
	payload []byte
}

func (e *RecoverPasswordStartedEvent) GetName() string {
	return "RecoverPasswordStartedEvent"
}

func (e *RecoverPasswordStartedEvent) GetPayload() []byte {
	return e.payload
}

func (e *RecoverPasswordStartedEvent) SetPayload(payload []byte) {
	e.payload = payload
}
