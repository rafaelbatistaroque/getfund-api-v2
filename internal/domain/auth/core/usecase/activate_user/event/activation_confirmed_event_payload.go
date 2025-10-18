package event

import shared_bus "getfund-api-v2/internal/shared/bus"

const ACTIVATE_USER_WITH_COUPON_CONFIRMED = "activate.user.with.coupon.confirmed"

// ActivateUserWithCouponConfirmedEvent is dispatched when a user is activated.
type ActivateUserWithCouponConfirmedEvent struct {
	shared_bus.EventBase
}

func (e *ActivateUserWithCouponConfirmedEvent) GetName() string {
	return ACTIVATE_USER_WITH_COUPON_CONFIRMED
}

// ActivateUserWithCouponConfirmedPayload is the data contract for the ActivateUserWithCouponConfirmedEvent.
type ActivateUserWithCouponConfirmedPayload struct {
	UserId     int    `json:"user_id"`
	Email      string `json:"email"`
	CouponCode string `json:"coupon_code"`
}
