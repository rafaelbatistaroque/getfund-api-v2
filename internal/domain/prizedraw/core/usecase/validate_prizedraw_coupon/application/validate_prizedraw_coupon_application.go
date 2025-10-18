package validate_prizedraw_coupon_application

import (
	"errors"
	"getfund-api-v2/internal/config/env"
	prizedraw_contract "getfund-api-v2/internal/domain/prizedraw/core/contract"
	"getfund-api-v2/internal/domain/prizedraw/core/dto/prizedraw_dto"
	"getfund-api-v2/internal/domain/prizedraw/core/usecase/validate_prizedraw_coupon"
	"getfund-api-v2/internal/domain/prizedraw/core/usecase/validate_prizedraw_coupon/event"
	shared_bus "getfund-api-v2/internal/shared/bus"
	shared_error "getfund-api-v2/internal/shared/error"
)

const (
	_INVALID_PRODUCT = "invalid product data"
	_COUPON_VALIDATE = "[coupon validate] "
)

type validatePrizeDrawCouponApplication struct {
	repository prizedraw_contract.Repository
	bus        shared_bus.EventBus
	settings   env.Variable
}

func New(repository prizedraw_contract.Repository, bus shared_bus.EventBus, env env.Variable) validate_prizedraw_coupon.UseCase {
	return &validatePrizeDrawCouponApplication{
		repository: repository,
		bus:        bus,
		settings:   env,
	}
}

func (v *validatePrizeDrawCouponApplication) Execute(input *validate_prizedraw_coupon.Input) (*validate_prizedraw_coupon.Output, *shared_error.Error) {
	validatable := input.Validate()
	if validatable.IsInvalid() {
		return nil, shared_error.New(shared_error.UNPROCESSABLE_CONTENT_CODE, validatable.GetErrors())
	}

	var applicationError = &shared_error.Error{}

	//Handle Coupon validation
	var couponDtoFound = &prizedraw_dto.CouponDto{}
	if couponDtoFound, applicationError = v.getCouponFromDb(input.CouponCode); applicationError != nil {
		return nil, applicationError
	}

	// coupon := v.getCouponEntityFilled(couponDtoFound)
	coupon := couponDtoFound.ToEntity()
	if err := coupon.Validate(input.Email, input.UserId); err != nil {
		return nil, shared_error.New(shared_error.UNAVAILABLE_CODE, err)
	}

	coupon.LinkPrizeDrawIfThereIsNo(input.SelectedPrizeDrawId)

	//Handle PrizeDraw validation
	var prizeDrawDtoFound *prizedraw_dto.PrizeDrawDto
	if prizeDrawDtoFound, applicationError = v.getPrizeDrawFromDb(coupon.GetPrizeDrawId()); applicationError != nil {
		return nil, applicationError
	}

	prizeDraw := prizeDrawDtoFound.ToEntity()
	if err := prizeDraw.Validate(input.SelectedPrizeDrawId); err != nil {
		return nil, shared_error.New(shared_error.UNAVAILABLE_CODE, err)
	}

	//Handle Product validation
	var productDto *prizedraw_dto.ProductDto
	if productDto, applicationError = v.getValidationCouponFromResponse(coupon.GetProductId()); applicationError != nil {
		return nil, applicationError
	}

	product := productDto.ToEntity()
	if err := product.Validate(input.SelectedProductId); err != nil {
		return nil, shared_error.New(shared_error.UNAVAILABLE_CODE, err)
	}

	return &validate_prizedraw_coupon.Output{
		Message:     "coupon is valid",
		CouponId:    coupon.GetId(),
		PrizeDrawId: prizeDraw.GetId(),
		ProductId:   product.GetId(),
	}, nil
}

func (v *validatePrizeDrawCouponApplication) getPrizeDrawFromDb(prizeDrawId int) (*prizedraw_dto.PrizeDrawDto, *shared_error.Error) {
	prizeDrawFound, err := v.repository.GetPrizeDrawById(prizeDrawId)
	if err != nil || prizeDrawFound == nil {
		return nil, shared_error.New(shared_error.NOT_FOUND_CODE, errors.New("prizedraw not found"))
	}

	return prizeDrawFound, nil
}

func (v *validatePrizeDrawCouponApplication) getCouponFromDb(couponCode string) (*prizedraw_dto.CouponDto, *shared_error.Error) {
	couponDtoFound, err := v.repository.GetCouponByCode(couponCode)
	if err != nil || couponDtoFound == nil {
		return nil, shared_error.New(shared_error.NOT_FOUND_CODE, errors.New("coupon not found"))
	}

	return couponDtoFound, nil
}

func (v *validatePrizeDrawCouponApplication) getValidationCouponFromResponse(productId int) (*prizedraw_dto.ProductDto, *shared_error.Error) {
	payload := &event.ValidatePrizeDrawCouponStartedPayload{
		ProductId: productId,
	}

	var product *prizedraw_dto.ProductDto
	promise := v.bus.EmitAndWaitPromise(&event.ValidatePrizeDrawCouponStartedEvent{}, payload, &product)
	if promise.HasError() {
		return nil, shared_error.New(shared_error.UNAVAILABLE_CODE, errors.New(_COUPON_VALIDATE+promise.GetError().Error()))
	}

	if product == nil {
		return nil, shared_error.New(shared_error.UNAVAILABLE_CODE, errors.New(_COUPON_VALIDATE+_INVALID_PRODUCT))
	}

	return product, nil
}
