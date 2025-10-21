package main

import (
	"context"
	config_gorm_postgres "getfund-api-v2/internal/config/db/gorm_postgres"
	"getfund-api-v2/internal/config/env"
	config_redis_cache "getfund-api-v2/internal/config/redis_cache"
	auth_composer "getfund-api-v2/internal/domain/auth/main/composer"
	notification_composer "getfund-api-v2/internal/domain/notification/main/composer"
	prizedraw_composer "getfund-api-v2/internal/domain/prizedraw/main/composer"
	shared_bus "getfund-api-v2/internal/shared/bus"

	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	//general dependences
	env_loaded := env.Load()
	bus := shared_bus.New(env_loaded.GetTimeoutResponseEvent())
	get_fund_db, opened_db := config_gorm_postgres.New(env_loaded)
	cache := config_redis_cache.New(context.Background(), env_loaded)

	defer cache.Close()
	defer opened_db.Close()

	//Composers
	notification_composer.Compose(env_loaded, bus, cache)
	prizedraw_composer.Compose(env_loaded, bus, cache, get_fund_db)
	authEndpoint := auth_composer.Compose(env_loaded, cache, get_fund_db, bus)

	//Routes
	r := chi.NewRouter()
	r.Route("/api/v2", func(route chi.Router) {
		route.Get("/", HelloWorld)

		//Auth
		route.
			With(authEndpoint.MiddlewareAutenticate).
			Get("/auth/sign-out", authEndpoint.Signout)
		route.Post("/auth/sign-in", authEndpoint.Signin)
		route.Post("/auth/recover-password", authEndpoint.RecoverPassword)
		route.Post("/auth/reset-password", authEndpoint.ResetPassword)
		route.Post("/auth/user", authEndpoint.Signup)
		route.Get("/auth/user/activate/{activation_code}", authEndpoint.ActivateUser)
	})

	http.ListenAndServe(env_loaded.GetPort(), r)
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API - GetFund v2.0 - ON"))
}
