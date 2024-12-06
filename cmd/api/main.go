package main

import (
	"context"
	authComposer "getfund-api-v2/internal/domain/auth/main/composer"
	notificationComposer "getfund-api-v2/internal/domain/notification/main/notificationcomposer"
	authMiddleware "getfund-api-v2/internal/middleware/authmiddleware"
	"getfund-api-v2/internal/settings"
	"getfund-api-v2/internal/shared/security"
	"getfund-api-v2/internal/shared/service/cacheservice"
	"getfund-api-v2/internal/shared/service/sessionservice"
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
	cacheService := cacheservice.New(redis, ctx)
	sessionService := sessionservice.New(cacheService, security.New(), appSettings)

	defer cacheService.Close()

	//Subscriber
	notificationComposer.SubscribeEventHandlers(appSettings, eventBus)

	//Routes
	r := chi.NewRouter()
	r.Route("/api/v2", func(api chi.Router) {
		api.Get("/", HelloWorld)

		//Auth
		authMiddleware := authMiddleware.New(sessionService)
		authHandlers := authComposer.GetHandlers(appSettings, cacheService, sessionService, db, eventBus)
		api.Post("/sign-in", authHandlers.Signin)
		api.With(authMiddleware.Authenticate).Get("/sign-out", authHandlers.Signout)
	})

	http.ListenAndServe(appSettings.GetPort(), r)
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API - GetFund v2.0 - ON"))
}
