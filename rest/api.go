package rest

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type API struct{}

type APIConfig struct{}

func NewAPI(config APIConfig) (*API, error) {
	return &API{}, nil
}

func (api *API) GetHandler() http.Handler {
	r := chi.NewRouter()

	r.Get("/", api.serveHealthCheck)

	r.Get("/books", api.serveGetBooks)

	return r
}

func (api *API) serveHealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("It's working!"))
}

func (api *API) serveGetBooks(w http.ResponseWriter, r *http.Request) {}
