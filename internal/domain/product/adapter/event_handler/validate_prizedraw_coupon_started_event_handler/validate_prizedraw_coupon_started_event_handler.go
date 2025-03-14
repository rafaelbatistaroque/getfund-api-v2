package validate_prizedraw_coupon_started_event_handler

import (
	"encoding/json"
	product_contract "getfund-api-v2/internal/domain/product/core/contract"
	"getfund-api-v2/pkg/bus"
	logger "getfund-api-v2/pkg/log"
)

type validatePrizeDrawCouponStartedStartedEventHandler struct {
	logger     logger.Logger
	repository product_contract.Repository
}

func New(repository product_contract.Repository) bus.Handler {
	return &validatePrizeDrawCouponStartedStartedEventHandler{
		logger:     *logger.New("validatePrizeDrawCouponStartedStartedEventHandler"),
		repository: repository,
	}
}

var payload struct {
	ProductId int
}

func (h *validatePrizeDrawCouponStartedStartedEventHandler) Handle(event bus.Event) {
	if err := json.Unmarshal(event.GetPayload(), &payload); err != nil {
		h.logger.Error("IsOk: False | get payload failed")
		return
	}

	h.repository.GetProductById(0)
}
