package user_entry_point_composer

import (
	"getfund-api-v2/internal/domain/user/adapter/gateway/activate_user_gateway"
	activate_user_application "getfund-api-v2/internal/domain/user/core/usecase/activate_user/application"
	"getfund-api-v2/internal/proxy/response_proxy"
	"getfund-api-v2/internal/settings"
	"getfund-api-v2/internal/shared/service/cache_service"
	"getfund-api-v2/pkg/bus"
	"net/http"

	"gorm.io/gorm"
)

type UserEntryPointComposer struct {
	ActivateUser http.HandlerFunc
}

func Get(
	settings settings.ApplicationSettings,
	cache cache_service.Cache,
	db *gorm.DB,
	eventBus bus.EventBus) UserEntryPointComposer {

	//dependencies

	//applications
	activate_user_application := activate_user_application.New(nil, nil, nil, eventBus, settings)

	//gateways
	activate_user_gateway := activate_user_gateway.New(activate_user_application)

	return UserEntryPointComposer{
		ActivateUser: response_proxy.New(activate_user_gateway.ActivateUser),
	}
}
