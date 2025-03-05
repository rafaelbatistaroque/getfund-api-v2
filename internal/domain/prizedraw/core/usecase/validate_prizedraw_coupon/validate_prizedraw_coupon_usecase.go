package validate_prizedraw_coupon

import "getfund-api-v2/internal/shared/result_app"

type UseCase = validatePrizeDrawCoupon

type validatePrizeDrawCoupon interface {
	Execute(input *Input) (*Output, *result_app.ApplicationError)
}
