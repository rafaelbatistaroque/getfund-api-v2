package validate_coupon_application

import (
	"bytes"
	"encoding/json"
	"errors"
	coupon_contract "getfund-api-v2/internal/domain/coupon/core/contract"
	"getfund-api-v2/internal/domain/coupon/core/dto/coupon_dto"
	"getfund-api-v2/internal/domain/coupon/core/dto/coupon_payload"
	"getfund-api-v2/internal/domain/coupon/core/usecase/validate_coupon"
	"getfund-api-v2/internal/settings"
	"getfund-api-v2/internal/shared/app_constant"
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
	settings   settings.ApplicationSettings
}

func New(repository coupon_contract.Repository, bus bus.EventBus, settings settings.ApplicationSettings) validate_coupon.UseCase {
	return &validateCouponApplication{
		repository: repository,
		bus:        bus,
		settings:   settings,
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

	if couponFound == nil {
		return nil, result_app.New(result_app.SERVER_ERROR_CODE, errors.New("error on get coupon data [found null]"))
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
	case response := <-responseChannel:
		if bytes.Equal(response, app_constant.EMPTYB) {
			return nil, result_app.New(result_app.UNAVAILABLE_CODE, errors.New("error on get coupon data [response empty]"))
		}

		var couponData = &coupon_dto.CouponValidationData{}
		if err := json.Unmarshal(response, &couponData); err != nil {
			return nil, result_app.New(result_app.UNAVAILABLE_CODE, errors.New("error on get coupon data [response invalid]"))
		}

		if couponData.Product == nil {
			return nil, result_app.New(result_app.UNAVAILABLE_CODE, errors.New("error on get coupon data [product invalid]"))
		}

		return nil, nil
	case <-time.After(time.Duration(v.settings.GetTimeoutResponseEvent()) * time.Second):
		return nil, result_app.New(result_app.UNAVAILABLE_CODE, errors.New("error on get coupon data [timeout]"))
	}
}
