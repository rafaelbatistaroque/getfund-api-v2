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
	_EMPTY_RESPONSE               = "empty response from coupon validation"
	_INVALID_RESPONSE             = "invalid response from coupon validation"
	_INVALID_PRODUCT              = "invalid product data"
	_COUPON_INVALID_FOR_PRODUCT   = "coupon is not valid for this product"
	_INACTIVE_PRODUCT             = "inactive product"
	_INVALID_PRIZE_DRAW           = "invalid prizedraw data"
	_PRIZE_DRAW_HAS_WINNER        = "prizedraw has winner"
	_COUPON_INVALID_FOR_PRIZEDRAW = "prizedraw is not valid for this coupon"
	_TIME_OUT                     = "timeout waiting for coupon validation"
	_EXPIRED_COUPON               = "coupon expired"
	_HAS_NOT_START                = "coupon validity has not start yet"
	_FOUND_NULL                   = "coupon null"
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

	if err := v.isCouponValid(couponFound); err != nil {
		return nil, err
	}

	responseChannel := make(chan []byte, 2)
	v.emitValidateCouponStartedEvent(couponFound, responseChannel)

	couponData, errResponse := v.getCouponValidationFromResponse(responseChannel)
	if errResponse != nil {
		return nil, errResponse
	}

	if err := v.isProductValid(couponData.Product, input.SelectedProductId); err != nil {
		return nil, err
	}

	if err := v.isPrizeDrawValid(couponData.PrizeDraw, input.SelectedPrizeDrawId); err != nil {
		return nil, err
	}

	return &validate_coupon.Output{
		Message: "coupon is valid",
	}, nil
}

func (*validateCouponApplication) isCouponValid(couponFound *coupon_dto.CouponDto) *result_app.ApplicationError {

	if couponFound == nil {
		return result_app.New(result_app.SERVER_ERROR_CODE, errors.New(_FOUND_NULL))
	}

	if couponFound.StartAt > _NOW {
		return result_app.New(result_app.UNAVAILABLE_CODE, errors.New(_HAS_NOT_START))
	}

	if couponFound.EndAt != nil && *couponFound.EndAt < _NOW {
		return result_app.New(result_app.UNAVAILABLE_CODE, errors.New(_EXPIRED_COUPON))
	}

	return nil
}

func (v *validateCouponApplication) emitValidateCouponStartedEvent(couponFound *coupon_dto.CouponDto, responseChannel chan []byte) {
	payload := &coupon_payload.ValidateCouponStartedPayload{
		ProductId:   couponFound.ProductId,
		PrizeDrawId: couponFound.PrizeDrawId,
	}

	v.bus.EmitWithPayloadAndResponse(&validate_coupon.ValidateCouponStartedEvent{}, payload, responseChannel)
}

func (v *validateCouponApplication) getCouponValidationFromResponse(responseChannel chan []byte) (*coupon_dto.CouponValidationData, *result_app.ApplicationError) {
	var couponData = &coupon_dto.CouponValidationData{}

	received := 0
	for received < 2 {
		select {
		case response := <-responseChannel:
			if bytes.Equal(response, app_constant.EMPTYB) {
				return nil, result_app.New(result_app.UNAVAILABLE_CODE, errors.New(_EMPTY_RESPONSE))
			}

			if err := json.Unmarshal(response, &couponData); err != nil {
				return nil, result_app.New(result_app.UNAVAILABLE_CODE, errors.New(_INVALID_RESPONSE))
			}

			received++
			continue

		case <-time.After(time.Duration(v.settings.GetTimeoutResponseEvent()) * time.Second):
			return nil, result_app.New(result_app.UNAVAILABLE_CODE, errors.New(_TIME_OUT))
		}
	}

	return couponData, nil
}

func (*validateCouponApplication) isProductValid(productData *coupon_dto.ProductData, selectedProductId int) *result_app.ApplicationError {
	if productData == nil {
		return result_app.New(result_app.UNAVAILABLE_CODE, errors.New(_INVALID_PRODUCT))
	}

	if !productData.IsActive {
		return result_app.New(result_app.UNAVAILABLE_CODE, errors.New(_INACTIVE_PRODUCT))
	}

	if productData.Id != selectedProductId {
		return result_app.New(result_app.UNAVAILABLE_CODE, errors.New(_COUPON_INVALID_FOR_PRODUCT))
	}

	return nil
}

func (*validateCouponApplication) isPrizeDrawValid(prizeDrawData *coupon_dto.PrizeDrawData, selectedPrizeDrawId int) *result_app.ApplicationError {
	if prizeDrawData == nil {
		return result_app.New(result_app.UNAVAILABLE_CODE, errors.New(_INVALID_PRIZE_DRAW))
	}

	if prizeDrawData.WinnerEntranceId != nil {
		return result_app.New(result_app.UNAVAILABLE_CODE, errors.New(_PRIZE_DRAW_HAS_WINNER))
	}

	if prizeDrawData.Id != selectedPrizeDrawId {
		return result_app.New(result_app.UNAVAILABLE_CODE, errors.New(_COUPON_INVALID_FOR_PRIZEDRAW))
	}

	return nil
}
