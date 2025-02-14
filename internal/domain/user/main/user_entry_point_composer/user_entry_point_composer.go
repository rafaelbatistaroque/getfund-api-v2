package user_entry_point_composer

import (
	"getfund-api-v2/internal/domain/user/adapter/gateway/activate_user_gateway"
	"getfund-api-v2/internal/domain/user/adapter/gateway/create_user_gateway"
	"getfund-api-v2/internal/domain/user/adapter/proxy/user_repository_proxy"
	user_repository "getfund-api-v2/internal/domain/user/adapter/repository"
	"getfund-api-v2/internal/domain/user/core/domain_service/activate_user_mapper"
	activate_user_application "getfund-api-v2/internal/domain/user/core/usecase/activate_user/application"
	create_user_application "getfund-api-v2/internal/domain/user/core/usecase/create_user/application"
	"getfund-api-v2/internal/proxy/response_proxy"
	"getfund-api-v2/internal/settings"
	"getfund-api-v2/internal/shared/security"
	"getfund-api-v2/internal/shared/service/cache_service"
	"getfund-api-v2/pkg/bus"
	"net/http"

	"gorm.io/gorm"
)

type userEntryPointComposer struct {
	CreateUser   http.HandlerFunc
	ActivateUser http.HandlerFunc
}

func Get(
	settings settings.ApplicationSettings,
	cache cache_service.Cache,
	db *gorm.DB,
	eventBus bus.EventBus) userEntryPointComposer {

	//dependencies
	hasher := security.New()
	mapper := activate_user_mapper.New()
	userRepositoryProxy := user_repository_proxy.New(user_repository.New(db), settings, hasher)

	//applications
	create_user_application := create_user_application.New(userRepositoryProxy, hasher, cache, eventBus, settings)
	activate_user_application := activate_user_application.New(cache, userRepositoryProxy, mapper, eventBus, settings)

	//gateways
	create_user_gateway := create_user_gateway.New(create_user_application)
	activate_user_gateway := activate_user_gateway.New(activate_user_application)

	return userEntryPointComposer{
		ActivateUser: response_proxy.New(activate_user_gateway.ActivateUser),
		CreateUser:   response_proxy.New(create_user_gateway.CreateUser),
	}
}
