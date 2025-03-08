package validate_prizedraw_coupon_application

import (
	"errors"
	prizedraw_contract "getfund-api-v2/internal/domain/prizedraw/core/contract"
	"getfund-api-v2/internal/domain/prizedraw/core/dto/prizedraw_dto"
	"getfund-api-v2/internal/domain/prizedraw/core/dto/prizedraw_payload"
	"getfund-api-v2/internal/domain/prizedraw/core/entity"
	"getfund-api-v2/internal/domain/prizedraw/core/usecase/validate_prizedraw_coupon"
	vo "getfund-api-v2/internal/domain/prizedraw/core/value_object"
	"getfund-api-v2/internal/settings"
	"getfund-api-v2/internal/shared/result_app"
	"getfund-api-v2/pkg/bus"
	"time"
)

const (
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
	_COUPON_NOT_APPLICABLE_EMAIL  = "coupon not applicable to this email"
	_COUPON_ALREADY_APPLIED       = "coupon already applied"
	_COUPON_LIMIT_REACHED         = "coupon application limit reached"
	_FOUND_NULL                   = "coupon null"
	_COUPON_VALIDATE              = "[coupon validate] "

	_NO_PRIZEDRAW_LINKED = 0
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
	coupon := v.getCouponEntityFilled(couponDtoFound)
	if err := v.isCouponValid(coupon, input); err != nil {
		return nil, err
	}

	//has prizeDraw linked
	if coupon.GetPrizeDrawId() == _NO_PRIZEDRAW_LINKED {
		coupon.SetPrazeDrawId(input.SelectedPrizeDrawId)
	}

	//Handle PrizeDraw validation
	var prizeDrawDtoFound *prizedraw_dto.PrizeDrawDto
	if prizeDrawDtoFound, applicationError = v.getPrizeDrawFromDb(coupon.GetPrizeDrawId()); applicationError != nil {
		return nil, applicationError
	}
	prizeDraw := entity.PrizeDrawFill(prizeDrawDtoFound.Id, prizeDrawDtoFound.WinnerEntranceId)
	if err := v.isPrizeDrawValid(prizeDraw, input.SelectedPrizeDrawId); err != nil {
		return nil, err
	}

	//Handle Product validation
	var validationData = &prizedraw_dto.ValidationData{}
	if validationData, applicationError = v.getValidationCouponFromResponse(coupon.GetProductId(), validationData); applicationError != nil {
		return nil, applicationError
	}

	product := entity.ProductFill(validationData.Product.Id, validationData.Product.IsActive)
	if applicationError := v.isProductValid(product, input.SelectedProductId); applicationError != nil {
		return nil, applicationError
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
	if err != nil {
		return nil, result_app.New(result_app.NOT_FOUND_CODE, errors.New(_PRIZE_DRAW_NOT_FOUND))
	}

	if prizeDrawFound == nil {
		return nil, result_app.New(result_app.UNAVAILABLE_CODE, errors.New(_INVALID_PRIZE_DRAW))
	}

	return prizeDrawFound, nil
}

func (v *validatePrizeDrawCouponApplication) getCouponFromDb(couponCode string) (*prizedraw_dto.CouponDto, *result_app.ApplicationError) {
	couponDtoFound, err := v.repository.GetCouponByCode(couponCode)
	if err != nil {
		return nil, result_app.New(result_app.NOT_FOUND_CODE, err)
	}

	if couponDtoFound == nil {
		return nil, result_app.New(result_app.SERVER_ERROR_CODE, errors.New(_FOUND_NULL))
	}

	return couponDtoFound, nil
}

func (v *validatePrizeDrawCouponApplication) getCouponEntityFilled(couponDtoFound *prizedraw_dto.CouponDto) *entity.Coupon {
	var endAt *time.Time
	if couponDtoFound.CouponTypeApplicability.EndAt != nil {
		endAtTime := time.Unix(*couponDtoFound.CouponTypeApplicability.EndAt, 0)
		endAt = &endAtTime
	}

	userCoupnApplies := make([]vo.CouponUserApply, len(couponDtoFound.UserCouponApplies))
	for i, userCouponApply := range couponDtoFound.UserCouponApplies {
		userCoupnApplies[i] = vo.NewUserCouponApply(userCouponApply.UserId)
	}

	couponApplicability := vo.NewCouponTypeApplicability(
		couponDtoFound.CouponTypeApplicability.Id,
		couponDtoFound.CouponTypeApplicability.CouponTypeCode,
		couponDtoFound.CouponTypeApplicability.LinkedEmail,
		time.Unix(couponDtoFound.CouponTypeApplicability.StartAt, 0),
		endAt,
		couponDtoFound.CouponTypeApplicability.LimitApplication,
	)

	return entity.CouponFill(
		couponDtoFound.Id,
		couponDtoFound.Code,
		couponDtoFound.PrizeDrawId,
		couponDtoFound.ProductId,
		userCoupnApplies,
		couponApplicability,
	)
}

func (*validatePrizeDrawCouponApplication) isCouponValid(coupon *entity.Coupon, input *validate_prizedraw_coupon.Input) *result_app.ApplicationError {
	if coupon.NotStartYet() {
		return result_app.New(result_app.UNAVAILABLE_CODE, errors.New(_HAS_NOT_START))
	}

	switch coupon.GetCouponTypeCode() {
	case entity.UNIQUE_APPLICATION_BY_EMAIL_TYPE:
		if coupon.IsNotSameLinkedEmail(input.Email) {
			return result_app.New(result_app.UNAVAILABLE_CODE, errors.New(_COUPON_NOT_APPLICABLE_EMAIL))
		}
	case entity.UNIQUE_APPLICATION_TYPE:
		if coupon.CountApplies() >= 1 {
			return result_app.New(result_app.UNAVAILABLE_CODE, errors.New(_COUPON_ALREADY_APPLIED))
		}
	case entity.LIMIT_APPLICATION_TYPE:
		if coupon.ReachedApplicationLimit() {
			return result_app.New(result_app.UNAVAILABLE_CODE, errors.New(_COUPON_LIMIT_REACHED))
		}
	case entity.EXPIRATION_TYPE:
		if coupon.IsExpired() {
			return result_app.New(result_app.UNAVAILABLE_CODE, errors.New(_EXPIRED_COUPON))
		}
	}

	//validate if user already applied coupon
	if coupon.CouponAlreadyAppliedByUser(input.UserId) {
		return result_app.New(result_app.UNAVAILABLE_CODE, errors.New(_COUPON_APPLIED_BY_USER))
	}

	return nil
}

func (v *validatePrizeDrawCouponApplication) getValidationCouponFromResponse(productId int, validationData *prizedraw_dto.ValidationData) (*prizedraw_dto.ValidationData, *result_app.ApplicationError) {
	payload := &prizedraw_payload.ValidatePrizeDrawCouponStartedPayload{
		ProductId: productId,
	}
	promise := v.bus.EmitAndWaitPromise(&validate_prizedraw_coupon.ValidatePrizeDrawCouponStartedEvent{}, payload, &validationData)
	if !promise.IsErrorNil() {
		return nil, result_app.New(result_app.UNAVAILABLE_CODE, errors.New(_COUPON_VALIDATE+promise.GetError().Error()))
	}

	if validationData.Product == nil {
		return nil, result_app.New(result_app.UNAVAILABLE_CODE, errors.New(_COUPON_VALIDATE+_INVALID_PRODUCT))
	}

	return validationData, nil
}

func (*validatePrizeDrawCouponApplication) isProductValid(product *entity.Product, selectedProductId int) *result_app.ApplicationError {
	if !product.IsActive() {
		return result_app.New(result_app.UNAVAILABLE_CODE, errors.New(_INACTIVE_PRODUCT))
	}

	if product.GetId() != selectedProductId {
		return result_app.New(result_app.UNAVAILABLE_CODE, errors.New(_COUPON_INVALID_FOR_PRODUCT))
	}

	return nil
}

func (*validatePrizeDrawCouponApplication) isPrizeDrawValid(prizeDraw *entity.PrizeDraw, selectedPrizeDrawId int) *result_app.ApplicationError {
	if prizeDraw.HasWinner() {
		return result_app.New(result_app.UNAVAILABLE_CODE, errors.New(_PRIZE_DRAW_HAS_WINNER))
	}

	if prizeDraw.GetId() != selectedPrizeDrawId {
		return result_app.New(result_app.UNAVAILABLE_CODE, errors.New(_COUPON_INVALID_FOR_PRIZEDRAW))
	}

	return nil
}
