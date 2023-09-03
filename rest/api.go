package rest

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ilhamsyahids/bookshelf-template/storage"
	"github.com/ilhamsyahids/bookshelf-template/utils"
)

type API struct {
	bookStorage storage.Storage
}

type APIConfig struct {
	BookStorage storage.Storage
}

func NewAPI(config APIConfig) (*API, error) {
	return &API{bookStorage: config.BookStorage}, nil
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

func (api *API) serveGetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := api.bookStorage.GetBooks()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	// output success response
	buf := new(bytes.Buffer)
	encoder := json.NewEncoder(buf)
	encoder.Encode(utils.NewSuccessResp(books))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	w.Write(buf.Bytes())
}
