package prizedraw_composer

import (
	"getfund-api-v2/internal/config/env"
	"getfund-api-v2/internal/domain/prizedraw/adapter/event_handler/activate_user_with_coupon_confirmed_event_handler"
	prizedraw_repository "getfund-api-v2/internal/domain/prizedraw/adapter/repository"
	apply_prizedraw_coupon_application "getfund-api-v2/internal/domain/prizedraw/core/usecase/apply_prizedraw_coupon/application"
	validate_prizedraw_coupon_application "getfund-api-v2/internal/domain/prizedraw/core/usecase/validate_prizedraw_coupon/application"
	"getfund-api-v2/internal/infra/db"
	shared_bus "getfund-api-v2/internal/shared/bus"
	"getfund-api-v2/internal/shared/cache"
	"getfund-api-v2/internal/shared/security"
)

func Compose(env env.Variable, eventBus shared_bus.EventBus, cacheService cache.Service, db *db.GetFund) {

	//Services
	prizedrawRepository := prizedraw_repository.New(db)
	hasher := security.NewHasher()

	//Applications
	validatePrizedrawCouponApplication := validate_prizedraw_coupon_application.New(prizedrawRepository, eventBus, env)
	applyPrizedrawCouponApplication := apply_prizedraw_coupon_application.New(prizedrawRepository, eventBus, hasher)

	//Event Handler
	handlers := map[string]shared_bus.Handler{
		"activate.user.with.coupon.confirmed": activate_user_with_coupon_confirmed_event_handler.New(prizedrawRepository, validatePrizedrawCouponApplication, applyPrizedrawCouponApplication, cacheService),
	}

	for eventName, handler := range handlers {
		eventBus.Subscribe(eventName, handler)
	}
}
