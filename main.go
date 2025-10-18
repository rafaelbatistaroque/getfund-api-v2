package main

import (
	"context"
	config_gorm_postgres "getfund-api-v2/internal/config/db/gorm_postgres"
	"getfund-api-v2/internal/config/env"
	config_redis_cache "getfund-api-v2/internal/config/redis_cache"
	"getfund-api-v2/internal/domain/auth/main/auth_composer"
	"getfund-api-v2/internal/domain/notification/main/notification_composer"
	"getfund-api-v2/internal/domain/prizedraw/main/prizedraw_composer"
	shared_bus "getfund-api-v2/internal/shared/bus"

	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	//general dependences
	ctx := context.Background()
	env_loaded := env.Load()
	bus := shared_bus.New(env_loaded.GetTimeoutResponseEvent())
	get_fund_db := config_gorm_postgres.New(env_loaded)
	cache := config_redis_cache.New(ctx, env_loaded)

	defer cache.Close()
	currentDb, _ := get_fund_db.DB.DB()
	defer currentDb.Close()

	//Composer
	notification_composer.Compose(env_loaded, bus, cache)
	prizedraw_composer.Compose(env_loaded, bus, cache, get_fund_db)
	authComposer := auth_composer.Compose(env_loaded, cache, get_fund_db, bus)

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
		route.Post("/auth/user", authComposer.Signup)
		route.Get("/auth/user/activate/{activation_code}", authComposer.ActivateUser)
	})

	http.ListenAndServe(env_loaded.GetPort(), r)
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API - GetFund v2.0 - ON"))
}
