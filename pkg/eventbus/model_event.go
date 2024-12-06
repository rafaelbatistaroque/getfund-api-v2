package eventbus

type UserSignedEvent struct {
	payload []byte
}

func (e *UserSignedEvent) GetName() string {
	return "UserSignedEvent"
}

func (e *UserSignedEvent) GetPayload() []byte {
	return e.payload
}

func (e *UserSignedEvent) SetPayload(payload []byte) {
	e.payload = payload
}
