package event

type UserCreated struct {
	payload []byte
	channel chan []byte
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

func (e *UserCreated) SetChannel(channel chan []byte) {
	e.channel = channel
}

func (e *UserCreated) DefineResponse(response []byte) {
	e.channel <- response
}
