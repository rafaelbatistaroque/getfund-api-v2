package event

type UserCriationStartedEvent struct {
	payload []byte
}

func (e *UserCriationStartedEvent) GetName() string {
	return "UserCriationStartedEvent"
}

func (e *UserCriationStartedEvent) GetPayload() []byte {
	return e.payload
}

func (e *UserCriationStartedEvent) SetPayload(payload []byte) {
	e.payload = payload
}
