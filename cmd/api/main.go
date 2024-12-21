package main

import (
	"context"
	auth_composer "getfund-api-v2/internal/domain/auth/main/composer"
	"getfund-api-v2/internal/domain/notification/main/notification_composer"
	"getfund-api-v2/internal/middleware/auth_middleware"
	"getfund-api-v2/internal/settings"
	"getfund-api-v2/internal/shared/security"
	"getfund-api-v2/internal/shared/service/cache_service"
	"getfund-api-v2/internal/shared/service/code_service"
	"getfund-api-v2/internal/shared/service/session_service"
	"getfund-api-v2/pkg/eventbus"
	redisconfig "getfund-api-v2/pkg/redis"
	sqlitedb "getfund-api-v2/pkg/sqlite"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	//general dependences
	db := sqlitedb.New()
	eventBus := eventbus.New()
	ctx := context.Background()
	appSettings := settings.Load()
	redis := redisconfig.New(ctx, appSettings)

	//Services
	cacheService := cache_service.New(redis, ctx)
	sessionService := session_service.New(cacheService, security.New(), appSettings)
	codeService := code_service.New(security.New(), appSettings)

	defer cacheService.Close()

	//Subscriber
	notification_composer.SubscribeEventHandlers(appSettings, eventBus)

	//Composer
	authHandlers := auth_composer.GetHandlers(appSettings, cacheService, sessionService, db, eventBus, codeService)

	//Routes
	r := chi.NewRouter()
	r.Route("/api/v2", func(api chi.Router) {
		api.Get("/", HelloWorld)

		//Auth
		authMiddleware := auth_middleware.New(sessionService)
		api.Post("/sign-in", authHandlers.Signin)
		api.With(authMiddleware.Authenticate).Get("/sign-out", authHandlers.Signout)
		api.Post("/recover-password", authHandlers.RecoverPassword)
	})

	http.ListenAndServe(appSettings.GetPort(), r)
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API - GetFund v2.0 - ON"))
}
