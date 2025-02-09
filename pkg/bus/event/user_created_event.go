package event

type UserCreated struct {
	payload []byte
}

func (e *UserCreated) GetName() string {
	return "UserCreated"
}

func (e *UserCreated) GetPayload() []byte {
	return e.payload
}

func (e *UserCreated) SetPayload(payload []byte) {
	e.payload = payload
}
