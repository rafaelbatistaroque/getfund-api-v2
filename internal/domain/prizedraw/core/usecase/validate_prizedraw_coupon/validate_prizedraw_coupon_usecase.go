package validate_prizedraw_coupon

import (
	shared_error "getfund-api-v2/internal/shared/error"
)

type UseCase = validatePrizeDrawCoupon

type validatePrizeDrawCoupon interface {
	Execute(input *Input) (*Output, *shared_error.Error)
}
