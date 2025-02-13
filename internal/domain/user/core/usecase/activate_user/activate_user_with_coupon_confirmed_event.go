package activate_user

type ActivateUserWithCouponConfirmedEvent struct {
	payload []byte
	channel chan []byte
}

func (e *ActivateUserWithCouponConfirmedEvent) GetName() string {
	return "ActivateUserWithCouponConfirmedEvent"
}

func (e *ActivateUserWithCouponConfirmedEvent) GetPayload() []byte {
	return e.payload
}

func (e *ActivateUserWithCouponConfirmedEvent) SetPayload(payload []byte) {
	e.payload = payload
}

func (e *ActivateUserWithCouponConfirmedEvent) SetChannel(channel chan []byte) {
	e.channel = channel
}

func (e *ActivateUserWithCouponConfirmedEvent) DefineResponse(response []byte) {
	e.channel <- response
}
