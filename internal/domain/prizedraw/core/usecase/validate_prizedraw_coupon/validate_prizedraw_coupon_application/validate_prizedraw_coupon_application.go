package validate_prizedraw_coupon_application

import (
	"bytes"
	"encoding/json"
	"errors"
	prizedraw_contract "getfund-api-v2/internal/domain/prizedraw/core/contract"
	"getfund-api-v2/internal/domain/prizedraw/core/dto/prizedraw_dto"
	"getfund-api-v2/internal/domain/prizedraw/core/dto/prizedraw_payload"
	"getfund-api-v2/internal/domain/prizedraw/core/entity"
	"getfund-api-v2/internal/domain/prizedraw/core/usecase/validate_prizedraw_coupon"
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
	_PRIZE_DRAW_NOT_FOUND         = "prizedraw not found"
	_INVALID_PRIZE_DRAW           = "invalid prizedraw data"
	_PRIZE_DRAW_HAS_WINNER        = "prizedraw has winner"
	_COUPON_INVALID_FOR_PRIZEDRAW = "prizedraw is not valid for this coupon"
	_TIME_OUT                     = "timeout waiting for coupon validation"
	_EXPIRED_COUPON               = "coupon expired"
	_HAS_NOT_START                = "coupon validity has not start yet"
	_COUPON_APPLIED_BY_USER       = "coupon already applied by user"
	_COUPON_ALREADY_APPLIED       = "coupon already applied"
	_FOUND_NULL                   = "coupon null"

	UNIQUE_APPLICATION_TYPE = 1
	LIMIT_APPLICATION_TYPE  = 2
	EXPIRATION_TYPE         = 3
)

type validateCouponApplication struct {
	repository prizedraw_contract.Repository
	bus        bus.EventBus
	settings   settings.ApplicationSettings
}

func New(repository prizedraw_contract.Repository, bus bus.EventBus, settings settings.ApplicationSettings) validate_prizedraw_coupon.UseCase {
	return &validateCouponApplication{
		repository: repository,
		bus:        bus,
		settings:   settings,
	}
}

func (v *validateCouponApplication) Execute(input *validate_prizedraw_coupon.Input) (*validate_prizedraw_coupon.Output, *result_app.ApplicationError) {
	validatable := input.Validate()
	if validatable.IsInvalid() {
		return nil, result_app.New(result_app.UNPROCESSABLE_CONTENT_CODE, validatable.GetErrors())
	}

	var applicationError = &result_app.ApplicationError{}

	//Handle Coupon validation
	var couponDtoFound = &prizedraw_dto.CouponDto{}
	if couponDtoFound, applicationError = v.getCouponFromDb(input.CouponCode); applicationError != nil {
		return nil, applicationError
	}
	coupon := v.getCouponEntityFilled(couponDtoFound)
	if err := v.isCouponValid(coupon, couponDtoFound.UserCouponApplies, input.UserId); err != nil {
		return nil, err
	}

	//Handle PrizeDraw validation
	var prizeDrawDtoFound *prizedraw_dto.PrizeDrawDto
	if prizeDrawDtoFound, applicationError = v.getPrizeDraw(couponDtoFound.PrizeDrawId); applicationError != nil {
		return nil, applicationError
	}
	prizeDraw := entity.PrizeDrawFill(prizeDrawDtoFound.Id, prizeDrawDtoFound.WinnerEntranceId)
	if err := v.isPrizeDrawValid(prizeDraw, input.SelectedPrizeDrawId); err != nil {
		return nil, err
	}

	//Handle Product validation
	responseChannel := make(chan []byte, 1)
	v.emitValidateCouponStartedEvent(couponDtoFound, responseChannel)

	var validationCouponData *prizedraw_dto.ValidationCouponData
	if validationCouponData, applicationError = v.getValidationCouponFromResponse(responseChannel); applicationError != nil {
		return nil, applicationError
	}
	if applicationError := v.isProductValid(validationCouponData.Product, input.SelectedProductId); applicationError != nil {
		return nil, applicationError
	}

	return &validate_prizedraw_coupon.Output{
		Message: "coupon is valid",
	}, nil
}

func (v *validateCouponApplication) getPrizeDraw(prizeDrawId int) (*prizedraw_dto.PrizeDrawDto, *result_app.ApplicationError) {
	prizeDrawFound, err := v.repository.GetPrizeDrawById(prizeDrawId)
	if err != nil {
		return nil, result_app.New(result_app.NOT_FOUND_CODE, errors.New(_PRIZE_DRAW_NOT_FOUND))
	}

	if prizeDrawFound == nil {
		return nil, result_app.New(result_app.UNAVAILABLE_CODE, errors.New(_INVALID_PRIZE_DRAW))
	}

	return prizeDrawFound, nil
}

func (v *validateCouponApplication) getCouponFromDb(couponCode string) (*prizedraw_dto.CouponDto, *result_app.ApplicationError) {
	couponDtoFound, err := v.repository.GetCouponByCode(couponCode)
	if err != nil {
		return nil, result_app.New(result_app.NOT_FOUND_CODE, err)
	}

	if couponDtoFound == nil {
		return nil, result_app.New(result_app.SERVER_ERROR_CODE, errors.New(_FOUND_NULL))
	}

	return couponDtoFound, nil
}

func (v *validateCouponApplication) getCouponEntityFilled(couponDtoFound *prizedraw_dto.CouponDto) *entity.Coupon {
	var endAt *time.Time
	if couponDtoFound.EndAt != nil {
		endAtTime := time.Unix(*couponDtoFound.EndAt, 0)
		endAt = &endAtTime
	}

	return entity.CouponFill(
		couponDtoFound.Code,
		couponDtoFound.TypeApplicability,
		couponDtoFound.PrizeDrawId,
		couponDtoFound.ProductId,
		couponDtoFound.Discount,
		couponDtoFound.LimitApplication,
		time.Unix(couponDtoFound.StartAt, 0),
		endAt,
	)
}

func (*validateCouponApplication) isCouponValid(coupon *entity.Coupon, userCouponApplies []prizedraw_dto.UserCouponApply, userId int) *result_app.ApplicationError {
	if coupon.NotStartYet() {
		return result_app.New(result_app.UNAVAILABLE_CODE, errors.New(_HAS_NOT_START))
	}

	switch coupon.GetTypeApplicability() {
	case UNIQUE_APPLICATION_TYPE:
		if len(userCouponApplies) == 1 {
			return result_app.New(result_app.UNAVAILABLE_CODE, errors.New(_COUPON_ALREADY_APPLIED))
		}
	case LIMIT_APPLICATION_TYPE:
		if coupon.ReachedApplicationLimit(len(userCouponApplies)) {
			return result_app.New(result_app.UNAVAILABLE_CODE, errors.New(_COUPON_ALREADY_APPLIED))
		}
	case EXPIRATION_TYPE:
		if coupon.IsExpired() {
			return result_app.New(result_app.UNAVAILABLE_CODE, errors.New(_EXPIRED_COUPON))
		}
	}

	for _, application := range userCouponApplies {
		if application.UserId == userId {
			return result_app.New(result_app.UNAVAILABLE_CODE, errors.New(_COUPON_APPLIED_BY_USER))
		}
	}

	return nil
}

func (v *validateCouponApplication) emitValidateCouponStartedEvent(couponFound *prizedraw_dto.CouponDto, responseChannel chan []byte) {
	payload := &prizedraw_payload.ValidateCouponStartedPayload{
		ProductId: couponFound.ProductId,
	}

	v.bus.EmitWithPayloadAndResponse(&validate_prizedraw_coupon.ValidateCouponStartedEvent{}, payload, responseChannel)
}

func (v *validateCouponApplication) getValidationCouponFromResponse(responseChannel chan []byte) (*prizedraw_dto.ValidationCouponData, *result_app.ApplicationError) {
	var validationCouponData = &prizedraw_dto.ValidationCouponData{}

	select {
	case response := <-responseChannel:
		if bytes.Equal(response, app_constant.EMPTYB) {
			return nil, result_app.New(result_app.UNAVAILABLE_CODE, errors.New(_EMPTY_RESPONSE))
		}

		if err := json.Unmarshal(response, &validationCouponData); err != nil {
			return nil, result_app.New(result_app.UNAVAILABLE_CODE, errors.New(_INVALID_RESPONSE))
		}

	case <-time.After(time.Duration(v.settings.GetTimeoutResponseEvent()) * time.Second):
		return nil, result_app.New(result_app.UNAVAILABLE_CODE, errors.New(_TIME_OUT))
	}

	return validationCouponData, nil
}

// passar para dominio product e receber apenas a entradas caso válido
func (*validateCouponApplication) isProductValid(productData *prizedraw_dto.ProductData, selectedProductId int) *result_app.ApplicationError {
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

func (*validateCouponApplication) isPrizeDrawValid(prizeDraw *entity.PrizeDraw, selectedPrizeDrawId int) *result_app.ApplicationError {
	if prizeDraw.HasWinner() {
		return result_app.New(result_app.UNAVAILABLE_CODE, errors.New(_PRIZE_DRAW_HAS_WINNER))
	}

	if prizeDraw.GetId() != selectedPrizeDrawId {
		return result_app.New(result_app.UNAVAILABLE_CODE, errors.New(_COUPON_INVALID_FOR_PRIZEDRAW))
	}

	return nil
}
