package main

import (
	"context"
	"getfund-api-v2/internal/domain/auth/main/auth_entry_point"
	"getfund-api-v2/internal/domain/notification/main/notification_composer"
	"getfund-api-v2/internal/middleware/auth_middleware"
	"getfund-api-v2/internal/settings"
	"getfund-api-v2/internal/shared/security"
	"getfund-api-v2/internal/shared/service/cache_service"
	"getfund-api-v2/internal/shared/service/session_service"
	"getfund-api-v2/pkg/bus"
	redisconfig "getfund-api-v2/pkg/redis"
	sqlitedb "getfund-api-v2/pkg/sqlite"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	//general dependences
	db := sqlitedb.New()
	eventBus := bus.New()
	ctx := context.Background()
	appSettings := settings.Load()
	redis := redisconfig.New(ctx, appSettings)

	//Services
	cacheService := cache_service.New(redis, ctx)
	sessionService := session_service.New(cacheService, security.New(), appSettings)

	defer cacheService.Close()

	//Subscriber
	notification_composer.SubscribeEventHandlers(appSettings, eventBus, cacheService)

	//Entry Points
	authEntryPoints := auth_entry_point.Get(appSettings, cacheService, sessionService, db, eventBus)

	//Routes
	r := chi.NewRouter()
	r.Route("/api/v2", func(api chi.Router) {
		api.Get("/", HelloWorld)

		//Auth
		authMiddleware := auth_middleware.New(sessionService)
		api.Post("/sign-in", authEntryPoints.Signin)
		api.With(authMiddleware.Authenticate).Get("/sign-out", authEntryPoints.Signout)
		api.Post("/recover-password", authEntryPoints.RecoverPassword)
		api.Post("/reset-password", authEntryPoints.ResetPassword)
	})

	http.ListenAndServe(appSettings.GetPort(), r)
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API - GetFund v2.0 - ON"))
}
