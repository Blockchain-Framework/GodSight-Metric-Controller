package system

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Blockchain-Framework/controller/internal/godsight-aggregator/config"
	log "github.com/Blockchain-Framework/controller/pkg/logger"
	"github.com/Blockchain-Framework/controller/pkg/version"
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Routes() *chi.Mux {

	router := chi.NewRouter()

	// service info
	router.Get("/info", func(w http.ResponseWriter, r *http.Request) {

		buildInfo := version.NewBuildInfo()

		info, err := json.Marshal(buildInfo)

		if err != nil {
			log.Error(r.Context()).Err(err).Msgf("Unable to generate version information for %s", config.Conf.Server.AppName)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		if _, err := w.Write(info); err != nil {
			log.Error(r.Context()).Err(err).Msg("Unable to respond with service info")
		}
	})

	// readiness info
	router.Get("/readyz", func(w http.ResponseWriter, r *http.Request) {

		if _, err := fmt.Fprintf(w, "healthy"); err != nil {
			log.Error(r.Context()).Err(err).Msg("Unable to respond with ready status")
		}
	})

	// health info
	router.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusOK)
		log.Trace(r.Context()).Msg("Responded OK for /healthz")
	})

	router.Method("GET", "/metrics", promhttp.Handler())

	return router
}
