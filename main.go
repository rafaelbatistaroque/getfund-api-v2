package main

import (
	"context"
	"getfund-api-v2/internal/domain/auth/main/auth_entry_point_composer"
	"getfund-api-v2/internal/domain/notification/main/notification_composer"
	"getfund-api-v2/internal/domain/user/main/user_entry_point_composer"
	"getfund-api-v2/internal/settings"
	"getfund-api-v2/internal/shared/service/cache_service"
	"getfund-api-v2/pkg/bus"
	postgresdb "getfund-api-v2/pkg/db/postgres"
	redisconfig "getfund-api-v2/pkg/redis"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	//general dependences
	eventBus := bus.New()
	ctx := context.Background()
	appSettings := settings.Load()
	db := postgresdb.New(appSettings)
	redis := redisconfig.New(ctx, appSettings)

	//Services
	cacheService := cache_service.New(redis, ctx)

	defer cacheService.Close()

	//Subscriber
	notification_composer.SubscribeEventHandlers(appSettings, eventBus, cacheService)

	//Entry Points
	authEntryPoints := auth_entry_point_composer.Get(appSettings, cacheService, db, eventBus)
	userEntryPoints := user_entry_point_composer.Get(appSettings, cacheService, db, eventBus)

	//Routes
	r := chi.NewRouter()
	r.Route("/api/v2", func(api chi.Router) {
		api.Get("/", HelloWorld)

		//Auth
		api.Post("/sign-in", authEntryPoints.Signin)
		api.With(authEntryPoints.MiddlewareAutenticate).Get("/sign-out", authEntryPoints.Signout)
		api.Post("/recover-password", authEntryPoints.RecoverPassword)
		api.Post("/reset-password", authEntryPoints.ResetPassword)

		//User
		api.Post("/user", userEntryPoints.CreateUser)
		api.Get("/user/activate/{activation_code}", userEntryPoints.ActivateUser)
	})

	http.ListenAndServe(appSettings.GetPort(), r)
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API - GetFund v2.0 - ON"))
}
