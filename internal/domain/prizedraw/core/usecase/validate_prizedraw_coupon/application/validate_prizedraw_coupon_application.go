package validate_prizedraw_coupon_application

import (
	"errors"
	prizedraw_contract "getfund-api-v2/internal/domain/prizedraw/core/contract"
	"getfund-api-v2/internal/domain/prizedraw/core/dto/prizedraw_dto"
	"getfund-api-v2/internal/domain/prizedraw/core/dto/prizedraw_payload"
	"getfund-api-v2/internal/domain/prizedraw/core/usecase/validate_prizedraw_coupon"
	"getfund-api-v2/internal/settings"
	"getfund-api-v2/internal/shared/result_app"
	"getfund-api-v2/pkg/bus"
)

const (
	_INVALID_PRODUCT = "invalid product data"
	_COUPON_VALIDATE = "[coupon validate] "
)

type validatePrizeDrawCouponApplication struct {
	repository prizedraw_contract.Repository
	bus        bus.EventBus
	settings   settings.ApplicationSettings
}

func New(repository prizedraw_contract.Repository, bus bus.EventBus, settings settings.ApplicationSettings) validate_prizedraw_coupon.UseCase {
	return &validatePrizeDrawCouponApplication{
		repository: repository,
		bus:        bus,
		settings:   settings,
	}
}

func (v *validatePrizeDrawCouponApplication) Execute(input *validate_prizedraw_coupon.Input) (*validate_prizedraw_coupon.Output, *result_app.ApplicationError) {
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

	// coupon := v.getCouponEntityFilled(couponDtoFound)
	coupon := couponDtoFound.ToEntity()
	if err := coupon.Validate(input.Email, input.UserId); err != nil {
		return nil, result_app.New(result_app.UNAVAILABLE_CODE, err)
	}

	coupon.LinkPrizeDrawIfThereIsNo(input.SelectedPrizeDrawId)

	//Handle PrizeDraw validation
	var prizeDrawDtoFound *prizedraw_dto.PrizeDrawDto
	if prizeDrawDtoFound, applicationError = v.getPrizeDrawFromDb(coupon.GetPrizeDrawId()); applicationError != nil {
		return nil, applicationError
	}

	prizeDraw := prizeDrawDtoFound.ToEntity()
	if err := prizeDraw.Validate(input.SelectedPrizeDrawId); err != nil {
		return nil, result_app.New(result_app.UNAVAILABLE_CODE, err)
	}

	//Handle Product validation
	var productDto *prizedraw_dto.ProductDto
	if productDto, applicationError = v.getValidationCouponFromResponse(coupon.GetProductId()); applicationError != nil {
		return nil, applicationError
	}

	product := productDto.ToEntity()
	if err := product.Validate(input.SelectedProductId); err != nil {
		return nil, result_app.New(result_app.UNAVAILABLE_CODE, err)
	}

	return &validate_prizedraw_coupon.Output{
		Message:     "coupon is valid",
		CouponId:    coupon.GetId(),
		PrizeDrawId: prizeDraw.GetId(),
		ProductId:   product.GetId(),
	}, nil
}

func (v *validatePrizeDrawCouponApplication) getPrizeDrawFromDb(prizeDrawId int) (*prizedraw_dto.PrizeDrawDto, *result_app.ApplicationError) {
	prizeDrawFound, err := v.repository.GetPrizeDrawById(prizeDrawId)
	if err != nil || prizeDrawFound == nil {
		return nil, result_app.New(result_app.NOT_FOUND_CODE, errors.New("prizedraw not found"))
	}

	return prizeDrawFound, nil
}

func (v *validatePrizeDrawCouponApplication) getCouponFromDb(couponCode string) (*prizedraw_dto.CouponDto, *result_app.ApplicationError) {
	couponDtoFound, err := v.repository.GetCouponByCode(couponCode)
	if err != nil || couponDtoFound == nil {
		return nil, result_app.New(result_app.NOT_FOUND_CODE, errors.New("coupon not found"))
	}

	return couponDtoFound, nil
}

func (v *validatePrizeDrawCouponApplication) getValidationCouponFromResponse(productId int) (*prizedraw_dto.ProductDto, *result_app.ApplicationError) {
	payload := &prizedraw_payload.ValidatePrizeDrawCouponStartedPayload{
		ProductId: productId,
	}

	var product *prizedraw_dto.ProductDto
	promise := v.bus.EmitAndWaitPromise(&validate_prizedraw_coupon.ValidatePrizeDrawCouponStartedEvent{}, payload, &product)
	if promise.HasError() {
		return nil, result_app.New(result_app.UNAVAILABLE_CODE, errors.New(_COUPON_VALIDATE+promise.GetError().Error()))
	}

	if product == nil {
		return nil, result_app.New(result_app.UNAVAILABLE_CODE, errors.New(_COUPON_VALIDATE+_INVALID_PRODUCT))
	}

	return product, nil
}
