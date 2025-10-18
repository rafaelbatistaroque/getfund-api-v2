package event

import shared_bus "getfund-api-v2/internal/shared/bus"

const VALIDATE_PRIZEDRAW_COUPON_STARTED = "validate.prizedraw.coupon.started"

// ValidatePrizeDrawCouponStartedEvent is dispatched when the coupon validation process starts.
type ValidatePrizeDrawCouponStartedEvent struct {
	shared_bus.EventBase
}

func (e *ValidatePrizeDrawCouponStartedEvent) GetName() string {
	return VALIDATE_PRIZEDRAW_COUPON_STARTED
}

// ValidatePrizeDrawCouponStartedPayload is the data contract for the ValidatePrizeDrawCouponStartedEvent.
type ValidatePrizeDrawCouponStartedPayload struct {
	ProductId   int    `json:"product_id"`
}