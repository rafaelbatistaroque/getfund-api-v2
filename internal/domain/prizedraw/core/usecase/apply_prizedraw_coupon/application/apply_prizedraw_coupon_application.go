package apply_prizedraw_coupon_application

import (
	prizedraw_contract "getfund-api-v2/internal/domain/prizedraw/core/contract"
	"getfund-api-v2/internal/domain/prizedraw/core/usecase/apply_prizedraw_coupon"
	"getfund-api-v2/internal/settings"
	"getfund-api-v2/internal/shared/result_app"
	"getfund-api-v2/pkg/bus"
)

type applyPrizeDrawCouponApplication struct {
	repository prizedraw_contract.Repository
	bus        bus.EventBus
	settings   settings.ApplicationSettings
}

func New(repository prizedraw_contract.Repository, bus bus.EventBus, settings settings.ApplicationSettings) apply_prizedraw_coupon.UseCase {
	return &applyPrizeDrawCouponApplication{
		repository: repository,
		bus:        bus,
		settings:   settings,
	}
}

func (a *applyPrizeDrawCouponApplication) Execute(input *apply_prizedraw_coupon.Input) (*apply_prizedraw_coupon.Output, *result_app.ApplicationError) {
	validatable := input.Validate()
	if validatable.IsInvalid() {
		return nil, result_app.New(result_app.UNPROCESSABLE_CONTENT_CODE, validatable.GetErrors())
	}

	return nil, nil
}
