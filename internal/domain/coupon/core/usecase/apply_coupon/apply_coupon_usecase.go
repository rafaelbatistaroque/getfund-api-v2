package apply_coupon

import "getfund-api-v2/internal/shared/result_app"

type UseCase = applyCoupon

type applyCoupon interface {
	Execute(input *Input) (*Output, *result_app.ApplicationError)
}
