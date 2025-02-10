package event

type RecoverPasswordStarted struct {
	payload []byte
	channel chan []byte
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

func (e *RecoverPasswordStarted) SetChannel(channel chan []byte) {
	e.channel = channel
}

func (e *RecoverPasswordStarted) DefineResponse(response []byte) {
	e.channel <- response
}
