package modules

import (
	"github.com/Blockchain-Framework/controller/internal/godsight-aggregator/modules/api"
	"github.com/Blockchain-Framework/controller/internal/godsight-aggregator/modules/system"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes() chi.Router {

	r := chi.NewRouter()
	r.Mount("/system", system.Routes())
	r.Mount("/api", api.Routes())

	return r
}
