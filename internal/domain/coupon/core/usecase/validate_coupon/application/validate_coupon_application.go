package validate_coupon_application

import (
	"errors"
	coupon_contract "getfund-api-v2/internal/domain/coupon/core/contract"
	"getfund-api-v2/internal/domain/coupon/core/dto/coupon_payload"
	"getfund-api-v2/internal/domain/coupon/core/usecase/validate_coupon"
	"getfund-api-v2/internal/shared/result_app"
	"getfund-api-v2/pkg/bus"
	"time"
)

var (
	_NOW = time.Now().Unix()
)

type validateCouponApplication struct {
	repository coupon_contract.Repository
	bus        bus.EventBus
}

func New(repository coupon_contract.Repository, bus bus.EventBus) validate_coupon.UseCase {
	return &validateCouponApplication{
		repository: repository,
		bus:        bus,
	}
}

func (v *validateCouponApplication) Execute(input *validate_coupon.Input) (*validate_coupon.Output, *result_app.ApplicationError) {
	validatable := input.Validate()
	if validatable.IsInvalid() {
		return nil, result_app.New(result_app.UNPROCESSABLE_CONTENT_CODE, validatable.GetErrors())
	}

	couponFound, err := v.repository.GetCouponByCode(input.CouponCode)
	if err != nil {
		return nil, result_app.New(result_app.NOT_FOUND_CODE, err)
	}

	if couponFound.StartAt > _NOW {
		return nil, result_app.New(result_app.UNAVAILABLE_CODE, errors.New("coupon validity has not start yet"))
	}

	if couponFound.EndAt != nil && *couponFound.EndAt < _NOW {
		return nil, result_app.New(result_app.UNAVAILABLE_CODE, errors.New("coupon expired"))
	}

	responseChannel := make(chan []byte, 2)
	payload := &coupon_payload.ValidateCouponStartedPayload{
		ProductId:   couponFound.ProductId,
		PrizeDrawId: couponFound.PrizeDrawId,
	}

	v.bus.EmitWithPayloadAndResponse(&validate_coupon.ValidateCouponStartedEvent{}, payload, responseChannel)

	select {
	case <-responseChannel:
		return nil, nil
	case <-time.After(time.Second):
		return nil, result_app.New(result_app.UNAVAILABLE_CODE, errors.New("error on get coupon data [timeout]"))
	}
}
