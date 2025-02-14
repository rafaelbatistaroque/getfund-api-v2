package user_criation_with_coupon_started_event_handler

import (
	"encoding/json"
	"getfund-api-v2/internal/domain/coupon/core/dto/coupon_payload"
	"getfund-api-v2/internal/domain/coupon/core/usecase/validate_coupon"
	"getfund-api-v2/internal/shared/service/cache_service"
	"getfund-api-v2/pkg/bus"
	logger "getfund-api-v2/pkg/log"
	"strings"
	"time"
)

type userCriationWithCouponStartedEventHandler struct {
	usecase validate_coupon.UseCase
	cache   cache_service.Cache
	logger  logger.Logger
}

func New(usecase validate_coupon.UseCase, cache cache_service.Cache) bus.Handler {
	return &userCriationWithCouponStartedEventHandler{
		logger:  *logger.New("userCriationWithCouponStartedEventHandler"),
		cache:   cache,
		usecase: usecase,
	}
}

func (h *userCriationWithCouponStartedEventHandler) Handle(event bus.Event) {
	var payload = &coupon_payload.UserCriationWithCouponPayload{}
	if err := json.Unmarshal(event.GetPayload(), payload); err != nil {
		h.logger.Error("IsOk: False | get payload failed")
		return
	}

	_, err := h.usecase.Execute(&validate_coupon.Input{
		CouponCode: payload.CouponCode,
	})

	if err != nil && strings.HasPrefix(err.Message.Error(), "status:") {
		payload.ErrorStatus = strings.TrimPrefix(err.Message.Error(), "status:")
	}

	key := "user_activation_" + payload.ActivationCode + "_coupon"
	h.cache.Set(key, payload, 24*time.Hour)
}
