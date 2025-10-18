package validate_prizedraw_coupon_started_event_handler_fixture

import (
	"encoding/json"
	event_prizedraw "getfund-api-v2/internal/domain/prizedraw/core/usecase/validate_prizedraw_coupon/event"
	"getfund-api-v2/internal/domain/product/adapter/event_handler/validate_prizedraw_coupon_started_event_handler"
	"getfund-api-v2/internal/domain/product/core/dto/product_dto"
	shared_bus "getfund-api-v2/internal/shared/bus"
	"getfund-api-v2/test/helper/repository_spy/product_repository_spy"
)

type ValidatePrizeDrawCouponStartedEventHandlerFixture struct {
	RepoSpy *product_repository_spy.ProductRepositorySpy
}

func NewSut() (shared_bus.Handler, *ValidatePrizeDrawCouponStartedEventHandlerFixture) {
	repoSpy := product_repository_spy.New()

	return validate_prizedraw_coupon_started_event_handler.New(repoSpy),
		&ValidatePrizeDrawCouponStartedEventHandlerFixture{
			RepoSpy: repoSpy,
		}
}

func GetInvalidValidatePrizeDrawCouponStartedEvent() *event_prizedraw.ValidatePrizeDrawCouponStartedEvent {
	return &event_prizedraw.ValidatePrizeDrawCouponStartedEvent{}
}

func GetValidValidatePrizeDrawCouponStartedEvent(productId int) *event_prizedraw.ValidatePrizeDrawCouponStartedEvent {
	payload, _ := json.Marshal(map[string]any{
		"product_id": productId,
	})
	channel := make(chan []byte, 1)

	event := &event_prizedraw.ValidatePrizeDrawCouponStartedEvent{}
	event.SetPayload(payload)
	event.SetChannel(channel)

	return event
}

func GetValidProduct() *product_dto.ProductDto {
	return &product_dto.ProductDto{
		Id:       1,
		IsActive: true,
	}
}
