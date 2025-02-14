package create_user

type CreateUserProcessStartedEvent struct {
	payload []byte
	channel chan []byte
}

func (e *CreateUserProcessStartedEvent) GetName() string {
	return "CreateUserProcessStartedEvent"
}

func (e *CreateUserProcessStartedEvent) GetPayload() []byte {
	return e.payload
}

func (e *CreateUserProcessStartedEvent) SetPayload(payload []byte) {
	e.payload = payload
}

func (e *CreateUserProcessStartedEvent) SetChannel(channel chan []byte) {
	e.channel = channel
}

func (e *CreateUserProcessStartedEvent) DefineResponse(response []byte) {
	e.channel <- response
}
