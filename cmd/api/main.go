package main

import (
	"getfund-api-v2/internal/pkg/helpers/settings"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	settings := settings.Load()

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API - GetFund v2.0 - ON"))
	})

	http.ListenAndServe(settings.Port, r)
}
