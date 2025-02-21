package create_user_process_with_coupon_started_event_handler_fixture

import (
	"encoding/json"
	coupon_common "getfund-api-v2/internal/domain/coupon/adapter/common"
	event_handler "getfund-api-v2/internal/domain/coupon/adapter/event_handler/create_user_process_with_coupon_started_event_handler"
	coupon_dto "getfund-api-v2/internal/domain/coupon/core/dto/coupon_dto"
	"getfund-api-v2/internal/domain/coupon/core/usecase/validate_coupon"
	"getfund-api-v2/internal/domain/user/core/usecase/create_user"
	"getfund-api-v2/pkg/bus"
	"getfund-api-v2/test/helper/cache_spy"
	"getfund-api-v2/test/helper/repository_spy/coupon_repository_spy"
)

type CreateUserWithCouponStartedEventHandlerFixture struct {
	CacheSpy *cache_spy.RedisCacheSpy
	RepoSpy  *coupon_repository_spy.CouponRepositorySpy
}

func NewSut() (bus.Handler, *CreateUserWithCouponStartedEventHandlerFixture) {
	cacheSpy := cache_spy.New()
	couponRepoSpy := coupon_repository_spy.New()

	return event_handler.New(couponRepoSpy, cacheSpy),
		&CreateUserWithCouponStartedEventHandlerFixture{
			CacheSpy: cacheSpy,
			RepoSpy:  couponRepoSpy,
		}
}

func GetInvalidCreateUserProcessWithCouponStartedEvent() *create_user.CreateUserProcessWithCouponStartedEvent {
	return &create_user.CreateUserProcessWithCouponStartedEvent{}
}

func GetValidateCouponInput() *validate_coupon.Input {
	return &validate_coupon.Input{
		CouponCode: "fake-coupon-code",
	}
}

func GetCouponDataFromSuccessDB(couponFromDb *coupon_dto.CouponDto) *coupon_common.CouponData {
	return &coupon_common.CouponData{
		Id:          couponFromDb.Id,
		Code:        couponFromDb.Code,
		PrizeDrawId: couponFromDb.PrizeDrawId,
		ProductId:   couponFromDb.ProductId,
		StartAt:     couponFromDb.StartAt,
		EndAt:       couponFromDb.EndAt,
		Discount:    couponFromDb.Discount,
	}
}

func GetCachedCouponData(errorData *coupon_common.ErrorData, couponData *coupon_common.CouponData) *coupon_common.CacheCouponData {
	isValid := true
	if errorData != nil {
		isValid = false
	}

	return &coupon_common.CacheCouponData{
		IsValid:    isValid,
		CouponData: couponData,
		Error:      errorData,
	}
}

func GetValidCreateUserProcessWithCouponStartedEvent() *create_user.CreateUserProcessWithCouponStartedEvent {
	payload, _ := json.Marshal(map[string]string{
		"coupon_code":         "fake-coupon-code",
		"activation_data_key": "fake-activation-data-key",
	})

	event := &create_user.CreateUserProcessWithCouponStartedEvent{}
	event.SetPayload(payload)

	return event
}
