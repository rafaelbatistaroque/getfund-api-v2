package event

type UserCriationStarted struct {
	payload []byte
}

func (e *UserCriationStarted) GetName() string {
	return "UserCriationStarted"
}

func (e *UserCriationStarted) GetPayload() []byte {
	return e.payload
}

func (e *UserCriationStarted) SetPayload(payload []byte) {
	e.payload = payload
}
