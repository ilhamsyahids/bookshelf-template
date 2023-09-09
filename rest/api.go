package rest

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/ilhamsyahids/bookshelf-template/storage"
	"github.com/ilhamsyahids/bookshelf-template/utils"
	"gopkg.in/validator.v2"
)

type API struct {
	bookStorage storage.Storage
}

type APIConfig struct {
	BookStorage storage.Storage `validate:"nonnil"`
}

func NewAPI(config APIConfig) (*API, error) {
	err := validator.Validate(config)
	if err != nil {
		return nil, fmt.Errorf("invalid API config: %w", err)
	}

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
	query := r.URL.Query().Get("query")

	pageStr := r.URL.Query().Get("page")
	if pageStr == "" {
		pageStr = "1"
	}
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		render.Render(w, r, utils.NewErrorResp(http.StatusBadRequest, ErrInvalidPage.Error()))
		return
	}

	limitStr := r.URL.Query().Get("limit")
	if limitStr == "" {
		limitStr = "10"
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		render.Render(w, r, utils.NewErrorResp(http.StatusBadRequest, ErrInvalidLimit.Error()))
		return
	}

	books, err := api.bookStorage.GetBooks(query, page, limit)
	if err != nil {
		render.Render(w, r, utils.NewErrorResp(http.StatusInternalServerError, err.Error()))
		return
	}

	// return success response
	render.Render(w, r, utils.NewSuccessResp(books))
}
