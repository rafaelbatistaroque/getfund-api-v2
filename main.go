package main

import (
	"context"
	"getfund-api-v2/internal/domain/auth/main/auth_composer"
	"getfund-api-v2/internal/domain/notification/main/notification_composer"
	"getfund-api-v2/internal/domain/prizedraw/main/prizedraw_composer"
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
	ctx := context.Background()
	appSettings := settings.Load()
	eventBus := bus.New(appSettings.GetTimeoutResponseEvent())
	db := postgresdb.New(appSettings)
	redis := redisconfig.New(ctx, appSettings)

	//Services
	cacheService := cache_service.New(redis, ctx)

	defer cacheService.Close()
	currentDb, _ := db.DB()
	defer currentDb.Close()

	//Composer
	notification_composer.Compose(appSettings, eventBus, cacheService)
	prizedraw_composer.Compose(appSettings, eventBus, cacheService, db)
	authComposer := auth_composer.Compose(appSettings, cacheService, db, eventBus)

	//Routes
	r := chi.NewRouter()
	r.Route("/api/v2", func(route chi.Router) {
		route.Get("/", HelloWorld)

		//Auth
		route.
			With(authComposer.MiddlewareAutenticate).
			Get("/auth/sign-out", authComposer.Signout)
		route.Post("/auth/sign-in", authComposer.Signin)
		route.Post("/auth/recover-password", authComposer.RecoverPassword)
		route.Post("/auth/reset-password", authComposer.ResetPassword)
		route.Post("/auth/user", authComposer.CreateUser)
		route.Get("/auth/user/activate/{activation_code}", authComposer.ActivateUser)
	})

	http.ListenAndServe(appSettings.GetPort(), r)
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API - GetFund v2.0 - ON"))
}
