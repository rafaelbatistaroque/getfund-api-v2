package recover_password

type RecoverPasswordStartedEvent struct {
	payload []byte
	channel chan []byte
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

func (e *RecoverPasswordStartedEvent) SetChannel(channel chan []byte) {
	e.channel = channel
}

func (e *RecoverPasswordStartedEvent) DefineResponse(response []byte) {
	e.channel <- response
}
