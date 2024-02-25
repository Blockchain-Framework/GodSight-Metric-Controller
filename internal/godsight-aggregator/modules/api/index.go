package api

import (
	v1 "github.com/Blockchain-Framework/controller/internal/godsight-aggregator/modules/api/v1"
	"github.com/go-chi/chi/v5"
)

func Routes() *chi.Mux {

	router := chi.NewRouter()

	router.Group(func(r chi.Router) {
		r.Mount("/v1", v1.Routes())
	})

	return router
}
