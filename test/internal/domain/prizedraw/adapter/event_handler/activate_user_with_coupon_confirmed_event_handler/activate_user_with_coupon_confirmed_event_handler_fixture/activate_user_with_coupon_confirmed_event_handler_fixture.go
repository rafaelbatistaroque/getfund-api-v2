package activate_user_with_coupon_confirmed_event_handler_fixture

import (
	"encoding/json"
	coupon_common "getfund-api-v2/internal/domain/prizedraw/adapter/common"
	event_handler "getfund-api-v2/internal/domain/prizedraw/adapter/event_handler/activate_user_with_coupon_confirmed_event_handler"
	"getfund-api-v2/internal/domain/user/core/usecase/activate_user"
	"getfund-api-v2/pkg/bus"
	"getfund-api-v2/test/helper/cache_spy"
)

type ActivateUserWithCouponConfirmedEventHandlerFixture struct {
	CacheSpy          *cache_spy.RedisCacheSpy
	ValidateCouponSpy *ValidateCouponApplicationSpy
	ApplyCouponSpy    *ApplyCouponApplicationSpy
}

func NewSut() (bus.Handler, *ActivateUserWithCouponConfirmedEventHandlerFixture) {
	cacheSpy := cache_spy.New()
	validateCouponSpy := NewValidateCoupon()
	applyCouponSpy := NewApplyCoupon()

	return event_handler.New(validateCouponSpy, applyCouponSpy, cacheSpy),
		&ActivateUserWithCouponConfirmedEventHandlerFixture{
			CacheSpy:          cacheSpy,
			ValidateCouponSpy: validateCouponSpy,
			ApplyCouponSpy:    applyCouponSpy,
		}
}

func GetInvalidActivateUserWithCouponConfirmedEvent() *activate_user.ActivateUserWithCouponConfirmedEvent {
	return &activate_user.ActivateUserWithCouponConfirmedEvent{}
}

func GetValidActivateUserWithCouponConfirmedEvent(activationDataKey string, userId int) *activate_user.ActivateUserWithCouponConfirmedEvent {
	payload, _ := json.Marshal(map[string]any{
		"user_id":             userId,
		"activation_data_key": activationDataKey,
	})

	event := &activate_user.ActivateUserWithCouponConfirmedEvent{}
	event.SetPayload(payload)

	return event
}

func GetCacheDataWithInvalidCoupon() string {
	return `{"is_valid": false, "error": {"code": "EXPIRED", "message": "the coupon code is expired"}}`
}

func GetCacheDataWithValidCoupon() string {
	return `{"is_valid": true, "coupon_data": {"id": 1, "code": "fake-coupon-code", "prize_draw_id": 1, "product_id": 1, "start_at": 1709596800, "end_at": 1709596800, "discount": 200}}`
}

func GetCouponData() coupon_common.CacheCouponData {
	var couponData = coupon_common.CacheCouponData{}

	json.Unmarshal([]byte(GetCacheDataWithValidCoupon()), &couponData)

	return couponData
}
