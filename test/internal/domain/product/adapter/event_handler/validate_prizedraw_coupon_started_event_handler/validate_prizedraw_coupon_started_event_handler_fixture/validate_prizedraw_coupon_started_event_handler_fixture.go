package validate_prizedraw_coupon_started_event_handler_fixture

import (
	"encoding/json"
	"getfund-api-v2/internal/domain/prizedraw/core/usecase/validate_prizedraw_coupon"
	"getfund-api-v2/internal/domain/product/adapter/event_handler/validate_prizedraw_coupon_started_event_handler"
	"getfund-api-v2/internal/domain/product/core/dto/product_dto"
	"getfund-api-v2/pkg/bus"
	"getfund-api-v2/test/helper/repository_spy/product_repository_spy"
)

type ValidatePrizeDrawCouponStartedEventHandlerFixture struct {
	RepoSpy *product_repository_spy.ProductRepositorySpy
}

func NewSut() (bus.Handler, *ValidatePrizeDrawCouponStartedEventHandlerFixture) {
	repoSpy := product_repository_spy.New()

	return validate_prizedraw_coupon_started_event_handler.New(repoSpy),
		&ValidatePrizeDrawCouponStartedEventHandlerFixture{
			RepoSpy: repoSpy,
		}
}

func GetInvalidValidatePrizeDrawCouponStartedEvent() *validate_prizedraw_coupon.ValidatePrizeDrawCouponStartedEvent {
	return &validate_prizedraw_coupon.ValidatePrizeDrawCouponStartedEvent{}
}

func GetValidValidatePrizeDrawCouponStartedEvent(productId int) *validate_prizedraw_coupon.ValidatePrizeDrawCouponStartedEvent {
	payload, _ := json.Marshal(map[string]any{
		"product_id": productId,
	})
	channel := make(chan []byte, 1)

	event := &validate_prizedraw_coupon.ValidatePrizeDrawCouponStartedEvent{}
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
