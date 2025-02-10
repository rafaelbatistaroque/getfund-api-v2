package event

type UserCriationWithCouponCodeStarted struct {
	payload []byte
	channel chan []byte
}

func (e *UserCriationWithCouponCodeStarted) GetName() string {
	return "UserCriationWithCouponCodeStarted"
}

func (e *UserCriationWithCouponCodeStarted) GetPayload() []byte {
	return e.payload
}

func (e *UserCriationWithCouponCodeStarted) SetPayload(payload []byte) {
	e.payload = payload
}

func (e *UserCriationWithCouponCodeStarted) SetChannel(channel chan []byte) {
	e.channel = channel
}

func (e *UserCriationWithCouponCodeStarted) DefineResponse(response []byte) {
	e.channel <- response
}
