package validate_coupon_application

import (
	coupon_contract "getfund-api-v2/internal/domain/coupon/core/contract"
	"getfund-api-v2/internal/domain/coupon/core/usecase/validate_coupon"
	"getfund-api-v2/internal/shared/result_app"
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

	_, err := v.repository.GetCouponByCode(input.CouponCode)
	if err != nil {
		return nil, result_app.New(result_app.NOT_FOUND_CODE, err)
	}
	return nil, nil
}
