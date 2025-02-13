package create_user

type UserCriationWithCouponStartedEvent struct {
	payload []byte
	channel chan []byte
}

func (e *UserCriationWithCouponStartedEvent) GetName() string {
	return "UserCriationWithCouponStartedEvent"
}

func (e *UserCriationWithCouponStartedEvent) GetPayload() []byte {
	return e.payload
}

func (e *UserCriationWithCouponStartedEvent) SetPayload(payload []byte) {
	e.payload = payload
}

func (e *UserCriationWithCouponStartedEvent) SetChannel(channel chan []byte) {
	e.channel = channel
}

func (e *UserCriationWithCouponStartedEvent) DefineResponse(response []byte) {
	e.channel <- response
}
