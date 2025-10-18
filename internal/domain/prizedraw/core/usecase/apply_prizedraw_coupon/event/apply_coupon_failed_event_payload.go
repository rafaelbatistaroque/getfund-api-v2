package event

import shared_bus "getfund-api-v2/internal/shared/bus"

const APPLY_PRIZEDRAW_COUPON_FAILED = "apply.prizedraw.coupon.failed"

// ApplyPrizeDrawCouponFailedEvent is dispatched when the coupon application fails.
type ApplyPrizeDrawCouponFailedEvent struct {
	shared_bus.EventBase
}

func (e *ApplyPrizeDrawCouponFailedEvent) GetName() string {
	return APPLY_PRIZEDRAW_COUPON_FAILED
}

// ApplyPrizeDrawCouponFailedPayload is the data contract for the ApplyPrizeDrawCouponFailedEvent.
type ApplyPrizeDrawCouponFailedPayload struct {
	PurchaseId int `json:"purchase_id"`
}
