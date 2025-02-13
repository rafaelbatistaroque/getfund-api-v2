package create_user

type UserCriationStartedEvent struct {
	payload []byte
	channel chan []byte
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

func (e *UserCriationStartedEvent) SetChannel(channel chan []byte) {
	e.channel = channel
}

func (e *UserCriationStartedEvent) DefineResponse(response []byte) {
	e.channel <- response
}
