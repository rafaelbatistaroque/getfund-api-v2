package event

type UserActivationWithCouponConfirmed struct {
	payload []byte
	channel chan []byte
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

func (e *UserActivationWithCouponConfirmed) SetChannel(channel chan []byte) {
	e.channel = channel
}

func (e *UserActivationWithCouponConfirmed) DefineResponse(response []byte) {
	e.channel <- response
}
