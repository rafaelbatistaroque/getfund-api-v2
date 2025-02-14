package create_user

type CreateUserProcessWithCouponStartedEvent struct {
	payload []byte
	channel chan []byte
}

func (e *CreateUserProcessWithCouponStartedEvent) GetName() string {
	return "CreateUserProcessWithCouponStartedEvent"
}

func (e *CreateUserProcessWithCouponStartedEvent) GetPayload() []byte {
	return e.payload
}

func (e *CreateUserProcessWithCouponStartedEvent) SetPayload(payload []byte) {
	e.payload = payload
}

func (e *CreateUserProcessWithCouponStartedEvent) SetChannel(channel chan []byte) {
	e.channel = channel
}

func (e *CreateUserProcessWithCouponStartedEvent) DefineResponse(response []byte) {
	e.channel <- response
}
