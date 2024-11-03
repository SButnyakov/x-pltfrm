package v1

import (
	"net/http"
	"x-pltfrm/music/upload/config"
	"x-pltfrm/music/upload/internal/http/v1/handlers"

	"github.com/go-chi/chi/v5"
)

func Router(cfg config.Routes) http.Handler {
	r := chi.NewRouter()
	r.Route(cfg.HTTP.V1.Root, func(r chi.Router) {
		r.Get(cfg.HTTP.V1.Hello, handlers.Hello())
	})
	return r
}
