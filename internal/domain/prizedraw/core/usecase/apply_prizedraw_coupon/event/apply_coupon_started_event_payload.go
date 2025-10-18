package event

import shared_bus "getfund-api-v2/internal/shared/bus"

const APPLY_PRIZEDRAW_COUPON_STARTED = "apply.prizedraw.coupon.started"

// ApplyPrizeDrawCouponStartedEvent is dispatched when the coupon application process starts.
type ApplyPrizeDrawCouponStartedEvent struct {
	shared_bus.EventBase
}

func (e *ApplyPrizeDrawCouponStartedEvent) GetName() string {
	return APPLY_PRIZEDRAW_COUPON_STARTED
}

// ApplyPrizeDrawCouponStartedPayload is the data contract for the ApplyPrizeDrawCouponStartedEvent.
type ApplyPrizeDrawCouponStartedPayload struct {
	UserId       int `json:"user_id"`
	CouponId     int `json:"coupon_id"`
	ProductId    int `json:"product_id"`
	PrizeDrawId  int `json:"prize_draw_id"`
	ItemQuantity int `json:"item_quantity"`
}
