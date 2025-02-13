package activate_user

type ActivateUserConfirmedEvent struct {
	payload []byte
	channel chan []byte
}

func (e *ActivateUserConfirmedEvent) GetName() string {
	return "ActivateUserConfirmedEvent"
}

func (e *ActivateUserConfirmedEvent) GetPayload() []byte {
	return e.payload
}

func (e *ActivateUserConfirmedEvent) SetPayload(payload []byte) {
	e.payload = payload
}

func (e *ActivateUserConfirmedEvent) SetChannel(channel chan []byte) {
	e.channel = channel
}

func (e *ActivateUserConfirmedEvent) DefineResponse(response []byte) {
	e.channel <- response
}
