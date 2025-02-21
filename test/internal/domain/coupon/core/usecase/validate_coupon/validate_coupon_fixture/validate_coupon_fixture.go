package validate_coupon_fixture

import (
	"getfund-api-v2/internal/domain/coupon/core/usecase/validate_coupon"
	validate_coupon_application "getfund-api-v2/internal/domain/coupon/core/usecase/validate_coupon/application"
	"getfund-api-v2/test/helper/repository_spy/coupon_repository_spy"
)

type ValidateCouponFixture struct {
	RepoSpy *coupon_repository_spy.CouponRepositorySpy
}

func NewSut() (validate_coupon.UseCase, *ValidateCouponFixture) {
	repoSpy := coupon_repository_spy.New()

	return validate_coupon_application.New(repoSpy),
		&ValidateCouponFixture{
			RepoSpy: repoSpy,
		}

}

type Option func(*validate_coupon.Input)

func GetInput(options ...Option) *validate_coupon.Input {
	input := &validate_coupon.Input{
		CouponCode: "FAKE_CPN",
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
