package validate_coupon

import "getfund-api-v2/pkg/bus"

type ValidateCouponStartedEvent struct {
	bus.EventBase
}

func (e *ValidateCouponStartedEvent) GetName() string {
	return "ValidateCouponStartedEvent"
}
