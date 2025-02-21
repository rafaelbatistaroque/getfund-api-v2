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

const (
	_RESPONSE_EMPTY     = "error on get coupon data [response empty]"
	_RESPONSE_INVALID   = "error on get coupon data [response invalid]"
	_PRODUCT_INVALID    = "error on get coupon data [product invalid]"
	_PRIZE_DRAW_INVALID = "error on get coupon data [prizedraw invalid]"
	_TIME_OUT           = "error on get coupon data [timeout]"
	_COUPON_EXPIRED     = "coupon expired"
	_HAS_NOT_START      = "coupon validity has not start yet"
	_FOUND_NULL         = "error on get coupon data [found null]"
)

var _NOW = time.Now().Unix()

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
		return nil, result_app.New(result_app.SERVER_ERROR_CODE, errors.New(_FOUND_NULL))
	}

	if couponFound.StartAt > _NOW {
		return nil, result_app.New(result_app.UNAVAILABLE_CODE, errors.New(_HAS_NOT_START))
	}

	if couponFound.EndAt != nil && *couponFound.EndAt < _NOW {
		return nil, result_app.New(result_app.UNAVAILABLE_CODE, errors.New(_COUPON_EXPIRED))
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
			return nil, result_app.New(result_app.UNAVAILABLE_CODE, errors.New(_RESPONSE_EMPTY))
		}

		var couponData = &coupon_dto.CouponValidationData{}
		if err := json.Unmarshal(response, &couponData); err != nil {
			return nil, result_app.New(result_app.UNAVAILABLE_CODE, errors.New(_RESPONSE_INVALID))
		}

		if couponData.Product == nil {
			return nil, result_app.New(result_app.UNAVAILABLE_CODE, errors.New(_PRODUCT_INVALID))
		}

		if couponData.PrizeDraw == nil {
			return nil, result_app.New(result_app.UNAVAILABLE_CODE, errors.New(_PRIZE_DRAW_INVALID))
		}

		return nil, nil
	case <-time.After(time.Duration(v.settings.GetTimeoutResponseEvent()) * time.Second):
		return nil, result_app.New(result_app.UNAVAILABLE_CODE, errors.New(_TIME_OUT))
	}
}
