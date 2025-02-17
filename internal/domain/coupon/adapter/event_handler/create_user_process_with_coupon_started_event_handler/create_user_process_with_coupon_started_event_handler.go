package create_user_process_with_coupon_started_event_handler

import (
	"encoding/json"
	coupon_common "getfund-api-v2/internal/domain/coupon/adapter/common"
	"getfund-api-v2/internal/domain/coupon/core/usecase/validate_coupon"
	"getfund-api-v2/internal/shared/service/cache_service"
	"getfund-api-v2/pkg/bus"
	logger "getfund-api-v2/pkg/log"
	"strings"
	"time"
)

type crateUserProcessWithCouponStartedEventHandler struct {
	validateCoupon validate_coupon.UseCase
	cache          cache_service.Cache
	logger         logger.Logger
}

func New(usecase validate_coupon.UseCase, cache cache_service.Cache) bus.Handler {
	return &crateUserProcessWithCouponStartedEventHandler{
		logger:         *logger.New("crateUserProcessWithCouponStartedEventHandler"),
		cache:          cache,
		validateCoupon: usecase,
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

	_, err := h.validateCoupon.Execute(&validate_coupon.Input{
		CouponCode: payload.CouponCode,
	})

	var cacheData = coupon_common.CacheCouponData{IsValid: true}
	if err != nil {
		if errorCode, errorMessage, ok := parseCouponValidationError(err.Message.Error()); ok {
			cacheData.IsValid = false
			cacheData.ErrorCode = errorCode
			cacheData.ErrorMessage = errorMessage
		}
	}

	key := payload.ActivationDataKey + "_coupon"
	h.cache.Set(key, cacheData, 24*time.Hour)
}

func parseCouponValidationError(errorStr string) (code, message string, valid bool) {
	const tag_prefix = "coupon_validation:"
	if !strings.HasPrefix(errorStr, tag_prefix) {
		return "", "", false
	}

	parts := strings.SplitN(strings.TrimPrefix(errorStr, tag_prefix), "|", 2)

	if len(parts) < 2 {
		return "", "", false
	}

	return parts[0], parts[1], true
}
