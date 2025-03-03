package activate_user

import "getfund-api-v2/pkg/bus"

type ActivateUserWithCouponConfirmedEvent struct {
	bus.EventBase
}

func (e *ActivateUserWithCouponConfirmedEvent) GetName() string {
	return "ActivateUserWithCouponConfirmedEvent"
}
