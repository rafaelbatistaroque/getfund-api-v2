package activate_user_with_coupon_confirmed_event_handler_fixture

import (
	"encoding/json"
	"getfund-api-v2/internal/domain/auth/core/usecase/activate_user"
	event_handler "getfund-api-v2/internal/domain/prizedraw/adapter/event_handler/activate_user_with_coupon_confirmed_event_handler"
	"getfund-api-v2/internal/domain/prizedraw/core/dto/prizedraw_dto"
	"getfund-api-v2/pkg/bus"
	"getfund-api-v2/test/helper/cache_spy"
	"getfund-api-v2/test/helper/repository_spy/prizedraw_repository_spy"
	"time"
)

type ActivateUserWithCouponConfirmedEventHandlerFixture struct {
	RepoSpy                    *prizedraw_repository_spy.CouponRepositorySpy
	CacheSpy                   *cache_spy.RedisCacheSpy
	ValidatePrizeDrawCouponSpy *ValidatePrizeDrawCouponApplicationSpy
	ApplyPrizeDrawCouponSpy    *ApplyPrizeDrawCouponApplicationSpy
}

func NewSut() (bus.Handler, *ActivateUserWithCouponConfirmedEventHandlerFixture) {
	repoSpy := prizedraw_repository_spy.New()
	cacheSpy := cache_spy.New()
	validatePrizeDrawCouponSpy := NewValidatePrizeDrawCoupon()
	applyPrizeDrawCouponSpy := NewApplyPrizeDrawCoupon()

	return event_handler.New(repoSpy, validatePrizeDrawCouponSpy, applyPrizeDrawCouponSpy, cacheSpy),
		&ActivateUserWithCouponConfirmedEventHandlerFixture{
			RepoSpy:                    repoSpy,
			CacheSpy:                   cacheSpy,
			ValidatePrizeDrawCouponSpy: validatePrizeDrawCouponSpy,
			ApplyPrizeDrawCouponSpy:    applyPrizeDrawCouponSpy,
		}
}

func GetInvalidActivateUserWithCouponConfirmedEvent() *activate_user.ActivateUserWithCouponConfirmedEvent {
	return &activate_user.ActivateUserWithCouponConfirmedEvent{}
}

func GetValidActivateUserWithCouponConfirmedEvent(couponCode, email string, userId int) *activate_user.ActivateUserWithCouponConfirmedEvent {
	payload, _ := json.Marshal(map[string]any{
		"user_id":     userId,
		"coupon_code": couponCode,
		"email":       email,
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

func GetValidCoupon() *prizedraw_dto.CouponDto {
	minus72Hours := time.Now().Add(-24 * time.Hour).Unix()
	minus24Hours := time.Now().Add(24 * time.Hour).Unix()

	couponTypeApplicability := &prizedraw_dto.CouponTypeApplicabilityDto{
		StartAt: minus72Hours,
		EndAt:   &minus24Hours,
	}
	return &prizedraw_dto.CouponDto{
		Id:                      1,
		Code:                    "fake-coupon",
		PrizeDrawId:             5,
		ProductId:               10,
		CouponTypeApplicability: couponTypeApplicability,
	}
}
