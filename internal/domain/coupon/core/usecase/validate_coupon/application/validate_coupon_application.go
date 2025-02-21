package validate_coupon_application

import (
	"errors"
	coupon_contract "getfund-api-v2/internal/domain/coupon/core/contract"
	"getfund-api-v2/internal/domain/coupon/core/usecase/validate_coupon"
	"getfund-api-v2/internal/shared/result_app"
	"time"
)

var (
	_NOW = uint64(time.Now().Unix())
)

type validateCouponApplication struct {
	repository coupon_contract.Repository
}

func New(repository coupon_contract.Repository) validate_coupon.UseCase {
	return &validateCouponApplication{
		repository: repository,
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

	//"coupon validity has not start yet"
	return nil, nil
}
