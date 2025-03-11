package apply_prizedraw_coupon

import "getfund-api-v2/pkg/bus"

type ApplyPrizeDrawCouponStartedEvent struct {
	bus.EventBase
}

func (e *ApplyPrizeDrawCouponStartedEvent) GetName() string {
	return "ApplyPrizeDrawCouponStartedEvent"
}
