package main

import (
	"context"
	authComposer "getfund-api-v2/internal/domain/auth/main/composer"
	notificationComposer "getfund-api-v2/internal/domain/notification/main/notificationcomposer"
	authMiddleware "getfund-api-v2/internal/middleware/authmiddleware"
	"getfund-api-v2/internal/pkg/db"
	"getfund-api-v2/internal/pkg/eventbus"
	"getfund-api-v2/internal/settings"
	"getfund-api-v2/internal/shared/security"
	"getfund-api-v2/internal/shared/service/cacheservice"
	"getfund-api-v2/internal/shared/service/sessionservice"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	appSettings := settings.Load()
	eventBus := eventbus.New()
	cache := cacheservice.New(context.Background(), appSettings)
	db := db.New()
	session := sessionservice.New(cache, security.New(), appSettings)

	defer cache.Close()

	authMiddleware.New(session)

	r := chi.NewRouter()

	notificationComposer.SubscribeEventHandlers(appSettings, eventBus)

	r.Route("/api/v2", func(api chi.Router) {
		api.Get("/", HelloWorld)

		authHandlers := authComposer.GetHandlers(appSettings, cache, session, db)
		api.Post("/sign-in", authHandlers.Signin)
	})

	http.ListenAndServe(appSettings.GetPort(), r)
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API - GetFund v2.0 - ON"))
}
