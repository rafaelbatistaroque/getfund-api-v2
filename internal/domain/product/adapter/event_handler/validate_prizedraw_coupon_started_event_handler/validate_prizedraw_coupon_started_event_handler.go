package validate_prizedraw_coupon_started_event_handler

import (
	"encoding/json"
	product_contract "getfund-api-v2/internal/domain/product/core/contract"
	"getfund-api-v2/internal/domain/product/core/dto"
	shared_bus "getfund-api-v2/internal/shared/bus"
	shared_constant "getfund-api-v2/internal/shared/constant"
	shared_logger "getfund-api-v2/internal/shared/log"
)

type validatePrizeDrawCouponStartedEventHandler struct {
	logger     shared_logger.Logger
	repository product_contract.Repository
}

func New(repository product_contract.Repository) shared_bus.Handler {
	return &validatePrizeDrawCouponStartedEventHandler{
		logger:     *shared_logger.New("validatePrizeDrawCouponStartedEventHandler"),
		repository: repository,
	}
}

var payload struct {
	ProductId int `json:"product_id"`
}

var promise struct {
	Id       int  `json:"id"`
	IsActive bool `json:"is_active"`
}

func (h *validatePrizeDrawCouponStartedEventHandler) Handle(event shared_bus.Event) {
	var err error
	if err = json.Unmarshal(event.GetPayload(), &payload); err != nil {
		h.logger.Error("IsOk: False | get payload failed")
		return
	}

	var productDto *dto.ProductDto
	if productDto, err = h.repository.GetProductById(payload.ProductId); err != nil {
		h.logger.Error("IsOk: False | get product failed")
		event.ResolvePromise(shared_constant.BYTE_EMPTY)
		return
	}

	if productDto == nil {
		h.logger.Error("IsOk: False | product not found")
		event.ResolvePromise(shared_constant.BYTE_EMPTY)
		return
	}

	promise.Id = productDto.Id
	promise.IsActive = productDto.IsActive

	var productFound []byte
	productFound, err = json.Marshal(promise)
	if err != nil {
		h.logger.Error("IsOk: False | marshal product failed")
		event.ResolvePromise(shared_constant.BYTE_EMPTY)
		return
	}

	event.ResolvePromise(productFound)
}
