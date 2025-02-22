package activate_user_with_coupon_confirmed_event_handler

import (
	"encoding/json"
	coupon_common "getfund-api-v2/internal/domain/coupon/adapter/common"
	"getfund-api-v2/internal/domain/coupon/core/usecase/apply_coupon"
	"getfund-api-v2/internal/domain/coupon/core/usecase/validate_coupon"
	"getfund-api-v2/internal/shared/result_app"
	"getfund-api-v2/internal/shared/service/cache_service"
	"getfund-api-v2/pkg/bus"
	logger "getfund-api-v2/pkg/log"
)

type activateUserWithCouponConfirmedEventHandler struct {
	validateCoupon validate_coupon.UseCase
	applyCoupon    apply_coupon.UseCase
	cache          cache_service.Cache
	logger         logger.Logger
}

func New(validateCoupon validate_coupon.UseCase, applyCoupon apply_coupon.UseCase, cache cache_service.Cache) bus.Handler {
	return &activateUserWithCouponConfirmedEventHandler{
		logger:         *logger.New("activateUserWithCouponConfirmedEventHandler"),
		cache:          cache,
		validateCoupon: validateCoupon,
		applyCoupon:    applyCoupon,
	}
}

var payload struct {
	UserId            int    `json:"user_id"`
	ActivationDataKey string `json:"activation_data_key"`
}

func (h *activateUserWithCouponConfirmedEventHandler) Handle(event bus.Event) {
	var err error
	if err = json.Unmarshal(event.GetPayload(), &payload); err != nil {
		h.logger.Error("IsOk: False | get payload failed")
		return
	}

	var cacheData string
	if cacheData, err = h.cache.Get(payload.ActivationDataKey + "_coupon"); err != nil {
		h.logger.Error("IsOk: False | get cache failed")
		return
	}

	var couponData = coupon_common.CacheCouponData{}
	if err = json.Unmarshal([]byte(cacheData), &couponData); err != nil {
		h.logger.Errorf("IsOk: False | %v", err)
		return
	}

	if !couponData.IsValid {
		h.logger.Errorf("IsOk: False | Invalid coupon code | code: %s | message: %s", couponData.Error.Code, couponData.Error.Message)
		return
	}

	var erroUsecase = &result_app.ApplicationError{}
	_, erroUsecase = h.validateCoupon.Execute(&validate_coupon.Input{
		CouponCode:          couponData.CouponData.Code,
		SelectedProductId:   couponData.CouponData.ProductId,
		SelectedPrizeDrawId: couponData.CouponData.PrizeDrawId,
		UserId:              payload.UserId,
	})

	if erroUsecase != nil {
		h.logger.Errorf("IsOk: False | error on validate coupon | code: %s | message: %s", erroUsecase.Code, erroUsecase.Message)
		return
	}

	_, erroUsecase = h.applyCoupon.Execute(&apply_coupon.Input{
		Id:          couponData.CouponData.Id,
		Code:        couponData.CouponData.Code,
		PrizeDrawId: couponData.CouponData.PrizeDrawId,
		ProductId:   couponData.CouponData.ProductId,
		StartAt:     couponData.CouponData.StartAt,
		EndAt:       couponData.CouponData.EndAt,
		Discount:    couponData.CouponData.Discount,
		UserId:      payload.UserId,
	})

	if erroUsecase != nil {
		h.logger.Errorf("IsOk: False | error on apply coupon | code: %s | message: %s", erroUsecase.Code, erroUsecase.Message)
		return
	}

}
