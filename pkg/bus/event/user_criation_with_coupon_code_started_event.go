package event

type UserCriationWithCouponCodeStartedEvent struct {
	payload []byte
}

func (e *UserCriationWithCouponCodeStartedEvent) GetName() string {
	return "UserCriationStartedEvent"
}

func (e *UserCriationWithCouponCodeStartedEvent) GetPayload() []byte {
	return e.payload
}

func (e *UserCriationWithCouponCodeStartedEvent) SetPayload(payload []byte) {
	e.payload = payload
}
