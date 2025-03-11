package prizedraw_composer

import (
	"getfund-api-v2/internal/domain/prizedraw/adapter/event_handler/activate_user_with_coupon_confirmed_event_handler"
	prizedraw_repository "getfund-api-v2/internal/domain/prizedraw/adapter/repository"
	apply_prizedraw_coupon_application "getfund-api-v2/internal/domain/prizedraw/core/usecase/apply_prizedraw_coupon/application"
	validate_prizedraw_coupon_application "getfund-api-v2/internal/domain/prizedraw/core/usecase/validate_prizedraw_coupon/application"
	"getfund-api-v2/internal/settings"
	"getfund-api-v2/internal/shared/security"
	"getfund-api-v2/internal/shared/service/cache_service"
	"getfund-api-v2/pkg/bus"

	"gorm.io/gorm"
)

func Compose(settings settings.ApplicationSettings, eventBus bus.EventBus, cacheService cache_service.Cache, db *gorm.DB) {

	//Services
	prizedrawRepository := prizedraw_repository.New(db)
	hasher := security.New()

	//Applications
	validatePrizedrawCouponApplication := validate_prizedraw_coupon_application.New(prizedrawRepository, eventBus, settings)
	applyPrizedrawCouponApplication := apply_prizedraw_coupon_application.New(prizedrawRepository, eventBus, hasher)

	//Event Handler
	handlers := map[string]bus.Handler{
		"ActivateUserWithCouponConfirmedEvent": activate_user_with_coupon_confirmed_event_handler.New(prizedrawRepository, validatePrizedrawCouponApplication, applyPrizedrawCouponApplication, cacheService),
	}

	for eventName, handler := range handlers {
		eventBus.Subscribe(eventName, handler)
	}
}
