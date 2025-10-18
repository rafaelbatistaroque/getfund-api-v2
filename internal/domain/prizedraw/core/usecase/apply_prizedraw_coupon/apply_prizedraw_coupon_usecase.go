package apply_prizedraw_coupon

import (
	shared_error "getfund-api-v2/internal/shared/error"
)

type UseCase = applyPrizeDrawCoupon

type applyPrizeDrawCoupon interface {
	Execute(input *Input) (*Output, *shared_error.Error)
}
