package activate_user_with_coupon_confirmed_event_handler

import (
	"encoding/json"
	prizedraw_contract "getfund-api-v2/internal/domain/prizedraw/core/contract"
	"getfund-api-v2/internal/domain/prizedraw/core/usecase/apply_prizedraw_coupon"
	"getfund-api-v2/internal/domain/prizedraw/core/usecase/validate_prizedraw_coupon"
	"getfund-api-v2/internal/shared/result_app"
	"getfund-api-v2/internal/shared/service/cache_service"
	"getfund-api-v2/pkg/bus"
	logger "getfund-api-v2/pkg/log"
)

type activateUserWithCouponConfirmedEventHandler struct {
	repository              prizedraw_contract.Repository
	validatePrizeDrawCoupon validate_prizedraw_coupon.UseCase
	applyPrizeDrawCoupon    apply_prizedraw_coupon.UseCase
	cache                   cache_service.Cache
	logger                  logger.Logger
}

func New(repository prizedraw_contract.Repository, validatePrizeDrawCoupon validate_prizedraw_coupon.UseCase, applyPrizeDrawCoupon apply_prizedraw_coupon.UseCase, cache cache_service.Cache) bus.Handler {
	return &activateUserWithCouponConfirmedEventHandler{
		logger:                  *logger.New("activateUserWithCouponConfirmedEventHandler"),
		repository:              repository,
		cache:                   cache,
		validatePrizeDrawCoupon: validatePrizeDrawCoupon,
		applyPrizeDrawCoupon:    applyPrizeDrawCoupon,
	}
}

var payload struct {
	UserId     int    `json:"user_id"`
	Email      string `json:"email"`
	CouponCode string `json:"coupon_code"`
}

func (h *activateUserWithCouponConfirmedEventHandler) Handle(event bus.Event) {
	var err error
	if err = json.Unmarshal(event.GetPayload(), &payload); err != nil {
		h.logger.Error("IsOk: False | get payload failed")
		return
	}

	couponDto, err := h.repository.GetCouponByCode(payload.CouponCode)
	if err != nil {
		h.logger.Errorf("IsOk: False | Coupon Repository | error: %s", err)
		return
	}

	if couponDto == nil {
		h.logger.Error("IsOk: False | Coupon Repository | coupon null")
		return
	}

	var erroUsecase = &result_app.ApplicationError{}
	_, erroUsecase = h.validatePrizeDrawCoupon.Execute(&validate_prizedraw_coupon.Input{
		SelectedProductId:   couponDto.ProductId,
		SelectedPrizeDrawId: couponDto.PrizeDrawId,
		CouponCode:          payload.CouponCode,
		UserId:              payload.UserId,
		Email:               payload.Email,
	})

	if erroUsecase != nil {
		h.logger.Errorf("IsOk: False | error on validate coupon | code: %s | message: %s", erroUsecase.Code, erroUsecase.Message)
		return
	}

	_, erroUsecase = h.applyPrizeDrawCoupon.Execute(&apply_prizedraw_coupon.Input{
		CouponId:    couponDto.Id,
		PrizeDrawId: couponDto.PrizeDrawId,
		ProductId:   couponDto.ProductId,
		UserId:      payload.UserId,
	})

	if erroUsecase != nil {
		h.logger.Errorf("IsOk: False | error on apply coupon | code: %s | message: %s", erroUsecase.Code, erroUsecase.Message)
		return
	}

}
