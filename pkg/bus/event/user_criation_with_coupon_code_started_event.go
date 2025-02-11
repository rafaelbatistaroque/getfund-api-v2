package event

type UserCriationWithCouponStarted struct {
	payload []byte
	channel chan []byte
}

func (e *UserCriationWithCouponStarted) GetName() string {
	return "UserCriationWithCouponCodeStarted"
}

func (e *UserCriationWithCouponStarted) GetPayload() []byte {
	return e.payload
}

func (e *UserCriationWithCouponStarted) SetPayload(payload []byte) {
	e.payload = payload
}

func (e *UserCriationWithCouponStarted) SetChannel(channel chan []byte) {
	e.channel = channel
}

func (e *UserCriationWithCouponStarted) DefineResponse(response []byte) {
	e.channel <- response
}
