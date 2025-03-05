package apply_prizedraw_coupon

import "getfund-api-v2/internal/shared/result_app"

type UseCase = applyPrizeDrawCoupon

type applyPrizeDrawCoupon interface {
	Execute(input *Input) (*Output, *result_app.ApplicationError)
}
