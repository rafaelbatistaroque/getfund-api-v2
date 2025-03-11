package apply_prizedraw_coupon

import "getfund-api-v2/pkg/bus"

type ApplyPrizeDrawCouponFailedEvent struct {
	bus.EventBase
}

func (e *ApplyPrizeDrawCouponFailedEvent) GetName() string {
	return "ApplyPrizeDrawCouponFailedEvent"
}
