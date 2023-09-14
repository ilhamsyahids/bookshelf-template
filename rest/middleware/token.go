package middleware

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/ilhamsyahids/bookshelf-template/utils"
)

func AccessAPIKeyVerifier() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			apiKey := r.Header.Get("X-API-Key")
			if apiKey == "" {
				render.Render(w, r, utils.NewErrorResp(http.StatusForbidden, ErrMissingAPIKey.Error()))
				return
			}

			if apiKey != "1234567890" {
				render.Render(w, r, utils.NewErrorResp(http.StatusForbidden, ErrInvalidAPIKey.Error()))
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
