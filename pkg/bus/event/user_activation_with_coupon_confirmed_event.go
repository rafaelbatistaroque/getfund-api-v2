package event

type UserActivationWithCouponConfirmed struct {
	payload []byte
}

func (e *UserActivationWithCouponConfirmed) GetName() string {
	return "UserActivationWithCouponConfirmed"
}

func (e *UserActivationWithCouponConfirmed) GetPayload() []byte {
	return e.payload
}

func (e *UserActivationWithCouponConfirmed) SetPayload(payload []byte) {
	e.payload = payload
}
