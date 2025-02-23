package create_user_process_with_coupon_started_event_handler

import (
	"encoding/json"
	coupon_common "getfund-api-v2/internal/domain/prizedraw/adapter/common"
	coupon_contract "getfund-api-v2/internal/domain/prizedraw/core/contract"
	"getfund-api-v2/internal/domain/prizedraw/core/dto/prizedraw_dto"
	"getfund-api-v2/internal/shared/service/cache_service"
	"getfund-api-v2/pkg/bus"
	logger "getfund-api-v2/pkg/log"
	"time"
)

type crateUserProcessWithCouponStartedEventHandler struct {
	couponRepository coupon_contract.Repository
	cache            cache_service.Cache
	logger           logger.Logger
}

func New(couponRepository coupon_contract.Repository, cache cache_service.Cache) bus.Handler {
	return &crateUserProcessWithCouponStartedEventHandler{
		logger:           *logger.New("crateUserProcessWithCouponStartedEventHandler"),
		cache:            cache,
		couponRepository: couponRepository,
	}
}

type CreateUserProcessWithCouponPayload struct {
	CouponCode        string `json:"coupon_code"`
	ActivationDataKey string `json:"activation_data_key"`
}

func (h *crateUserProcessWithCouponStartedEventHandler) Handle(event bus.Event) {
	var payload = &CreateUserProcessWithCouponPayload{}
	if err := json.Unmarshal(event.GetPayload(), payload); err != nil {
		h.logger.Error("IsOk: False | get payload failed")
		return
	}

	var cacheData = &coupon_common.CacheCouponData{
		IsValid: true,
	}

	coupon, err := h.couponRepository.GetCouponByCode(payload.CouponCode)
	if err != nil {
		cacheData.IsValid = false
		cacheData.Error = getErrorDataFrom(err)
	}

	cacheData.CouponData = getCouponDataFrom(coupon)

	key := payload.ActivationDataKey + "_coupon"
	h.cache.Set(key, cacheData, 24*time.Hour)
}

func getErrorDataFrom(err error) *coupon_common.ErrorData {
	if err == nil {
		return nil
	}

	return &coupon_common.ErrorData{
		Code:    "COUPON_REPOSITORY",
		Message: err.Error(),
	}
}

func getCouponDataFrom(coupon *prizedraw_dto.CouponDto) *coupon_common.CouponData {
	if coupon == nil {
		return nil
	}

	return &coupon_common.CouponData{
		Id:          coupon.Id,
		Code:        coupon.Code,
		PrizeDrawId: coupon.PrizeDrawId,
		ProductId:   coupon.ProductId,
		StartAt:     coupon.StartAt,
		EndAt:       coupon.EndAt,
		Discount:    coupon.Discount,
	}
}
