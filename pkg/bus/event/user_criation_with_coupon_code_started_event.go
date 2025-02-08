package event

type UserCriationWithCouponCodeStarted struct {
	payload []byte
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
