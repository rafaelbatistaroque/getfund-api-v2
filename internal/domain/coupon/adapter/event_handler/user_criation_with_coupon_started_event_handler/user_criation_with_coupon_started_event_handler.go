package user_criation_with_coupon_started_event_handler

import (
	"encoding/json"
	"getfund-api-v2/internal/domain/coupon/core/usecase/validate_coupon"
	"getfund-api-v2/internal/shared/service/cache_service"
	"getfund-api-v2/pkg/bus"
	logger "getfund-api-v2/pkg/log"
)

type UserCriationWithCouponPayloadDto struct {
	CouponCode     string `json:"coupon_code"`
	ActivationCode string `json:"activation_code"`
}

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
	var payload = &UserCriationWithCouponPayloadDto{}
	if err := json.Unmarshal(event.GetPayload(), payload); err != nil {
		h.logger.Error("IsOk: False | get payload failed")
		return
	}

	var input = validate_coupon.Input{}
	h.usecase.Execute(&input)
}
