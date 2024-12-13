package event

type RecoverPasswordStarted struct {
	payload []byte
}

func (e *RecoverPasswordStarted) GetName() string {
	return "RecoverPasswordStarted"
}

func (e *RecoverPasswordStarted) GetPayload() []byte {
	return e.payload
}

func (e *RecoverPasswordStarted) SetPayload(payload []byte) {
	e.payload = payload
}
