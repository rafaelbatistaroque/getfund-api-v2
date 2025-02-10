package validate_coupon_code

import "getfund-api-v2/internal/shared/result_app"

type UseCase = validateCouponCode

type validateCouponCode interface {
	Execute(input *Input) (*Output, *result_app.ApplicationError)
}
