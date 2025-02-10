package event

type UserCriationStarted struct {
	payload []byte
	channel chan []byte
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

func (e *UserCriationStarted) SetChannel(channel chan []byte) {
	e.channel = channel
}

func (e *UserCriationStarted) DefineResponse(response []byte) {
	e.channel <- response
}
