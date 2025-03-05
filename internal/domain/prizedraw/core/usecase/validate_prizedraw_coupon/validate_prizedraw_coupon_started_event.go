package validate_prizedraw_coupon

import "getfund-api-v2/pkg/bus"

type ValidatePrizeDrawCouponStartedEvent struct {
	bus.EventBase
}

func (e *ValidatePrizeDrawCouponStartedEvent) GetName() string {
	return "ValidateCouponStartedEvent"
}
