package main

import (
	"context"
	authComposer "getfund-api-v2/internal/domain/auth/main/composer"
	notificationComposer "getfund-api-v2/internal/domain/notification/main/notificationcomposer"
	applog "getfund-api-v2/internal/log"
	authMiddleware "getfund-api-v2/internal/middleware/authmiddleware"
	"getfund-api-v2/internal/pkg/cache"
	"getfund-api-v2/internal/pkg/db"
	"getfund-api-v2/internal/pkg/eventbus"
	"getfund-api-v2/internal/settings"
	"getfund-api-v2/internal/shared/security"
	"getfund-api-v2/internal/shared/service/sessionservice"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	//general dependences
	applog.Load()
	db := db.New()
	appSettings := settings.Load()
	eventBus := eventbus.New()
	cache := cache.New(context.Background(), appSettings)
	sessionService := sessionservice.New(cache, security.New(), appSettings)

	defer cache.Close()

	//Subscriber
	notificationComposer.SubscribeEventHandlers(appSettings, eventBus)

	//Routes
	r := chi.NewRouter()
	r.Route("/api/v2", func(api chi.Router) {
		api.Get("/", HelloWorld)

		//Auth
		authMiddleware := authMiddleware.New(sessionService)
		authHandlers := authComposer.GetHandlers(appSettings, cache, sessionService, db, eventBus)
		api.Post("/sign-in", authHandlers.Signin)
		api.With(authMiddleware.Authenticate).Get("/sign-out", authHandlers.Signout)
	})

	http.ListenAndServe(appSettings.GetPort(), r)
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API - GetFund v2.0 - ON"))
}
