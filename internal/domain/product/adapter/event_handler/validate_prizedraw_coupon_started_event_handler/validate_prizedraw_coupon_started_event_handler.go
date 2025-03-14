package validate_prizedraw_coupon_started_event_handler

import (
	"encoding/json"
	product_contract "getfund-api-v2/internal/domain/product/core/contract"
	"getfund-api-v2/internal/domain/product/core/dto/product_dto"
	"getfund-api-v2/internal/shared/app_constant"
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
	ProductId int `json:"product_id"`
}

var promise struct {
	Id       int  `json:"id"`
	IsActive bool `json:"is_active"`
}

func (h *validatePrizeDrawCouponStartedStartedEventHandler) Handle(event bus.Event) {
	var err error
	if err = json.Unmarshal(event.GetPayload(), &payload); err != nil {
		h.logger.Error("IsOk: False | get payload failed")
		return
	}

	var productDto *product_dto.ProductDto
	if productDto, err = h.repository.GetProductById(payload.ProductId); err != nil {
		h.logger.Error("IsOk: False | get product failed")
		event.ResolvePromise(app_constant.EMPTYB)
		return
	}

	if productDto == nil {
		h.logger.Error("IsOk: False | product not found")
		event.ResolvePromise(app_constant.EMPTYB)
		return
	}

	promise.Id = productDto.Id
	promise.IsActive = productDto.IsActive

	var productFound []byte
	productFound, err = json.Marshal(promise)
	if err != nil {
		h.logger.Error("IsOk: False | marshal product failed")
		event.ResolvePromise(app_constant.EMPTYB)
		return
	}

	event.ResolvePromise(productFound)
}
