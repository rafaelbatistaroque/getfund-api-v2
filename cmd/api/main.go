package main

import (
	"getfund-api-v2/internal/pkg/settings"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	settings := settings.Load()

	r := chi.NewRouter()

	r.Route("/api/v2", func(api chi.Router) {
		api.Get("/", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("API - GetFund v2.0 - ON"))
		})
	})

	http.ListenAndServe(settings.GetPort(), r)
}
