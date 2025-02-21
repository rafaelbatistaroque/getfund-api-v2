package validate_coupon_fixture

import (
	"getfund-api-v2/internal/domain/coupon/core/usecase/validate_coupon"
	validate_coupon_application "getfund-api-v2/internal/domain/coupon/core/usecase/validate_coupon/application"
)

type ValidateCouponFixture struct {
}

func NewSut() (validate_coupon.UseCase, *ValidateCouponFixture) {

	return validate_coupon_application.New(),
		&ValidateCouponFixture{}

}

type Option func(*validate_coupon.Input)

func GetInput(options ...Option) *validate_coupon.Input {
	input := &validate_coupon.Input{
		CouponCode:  "fake-coupon-code",
		ProductId:   1,
		PrizeDrawId: 1,
	}

	for _, opt := range options {
		opt(input)
	}

	return input
}

func WithEmptyCouponCode() Option {
	return func(params *validate_coupon.Input) {
		params.CouponCode = ""
	}
}

func WithInvalidCouponCode() Option {
	return func(params *validate_coupon.Input) {
		params.CouponCode = "fake" //less than 8 characters
	}
}
