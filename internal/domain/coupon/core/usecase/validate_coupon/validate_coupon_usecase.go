package validate_coupon

import "getfund-api-v2/internal/shared/result_app"

type UseCase = validateCoupon

type validateCoupon interface {
	Execute(input *Input) (*Output, *result_app.ApplicationError)
}
