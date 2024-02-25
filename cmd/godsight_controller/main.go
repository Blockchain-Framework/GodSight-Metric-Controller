package godsight_controller

import (
	"context"
	"fmt"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Blockchain-Framework/controller/internal/godsight-aggregator/config"
	"github.com/Blockchain-Framework/controller/internal/godsight-aggregator/modules"
	"github.com/Blockchain-Framework/controller/pkg/middleware"
	"github.com/Blockchain-Framework/controller/pkg/util"
	"github.com/Blockchain-Framework/controller/pkg/version"

	log "github.com/Blockchain-Framework/controller/pkg/logger"
	"github.com/go-chi/chi/v5"
)

func printEndpoints(ctx context.Context, router *chi.Mux) {

	walkFunc := func(method string, route string, handler http.Handler, middleware ...func(http.Handler) http.Handler) error {
		log.Debug(ctx).Msgf(">>> %s %s", method, route)
		return nil
	}

	log.Debug(ctx).Msg("> Available routes <")

	if err := chi.Walk(router, walkFunc); err != nil {
		log.Debug(ctx).Err(err).Msg("Unable to print routes.")
	}

	log.Debug(ctx).Msg("< Available routes >")
}

func Start(conf *config.Config) {

	ctx := context.Background()
	log.New(conf.Debug)

	_, globalCancel := context.WithCancel(context.Background())
	defer globalCancel()

	// setup app deployer
	if val, err := util.NewValidator(); err != nil {
		log.Fatal(ctx).Err(err).Msgf("cannot create validator for %s", conf.Server.AppName)
		return
	} else {
		conf.Validator = val
	}

	router := chi.NewRouter()

	// custom middleware
	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://admin.dev.cmlinsight.com", "http://localhost:3000"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		//ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
		Debug:            false,
	}))
	router.Use(middleware.TraceId)
	router.Use(middleware.MandatoryHeaders)
	router.Use(middleware.RequestLogger)
	router.Use(middleware.Recoverer)

	// add system and api routes
	router.Mount(config.Conf.Server.ContextPath, modules.SetupRoutes())

	//Swagger Setup
	// @title App Deployer
	// @version 1.0
	// @termsOfService http://swagger.io/terms/

	// @contact.name CmlInsight API Support
	// @contact.url https://cmlinsight.com
	// @contact.email contact@cmlinsight.com
	// @BasePath /godsight-aggregator/api/v1
	router.Get("/swagger/*", httpSwagger.WrapHandler)

	log.Info(ctx).Msgf(
		"Build = release: %s | commit: %s | build time: %s",
		version.Release, version.Commit, version.BuildTime,
	)

	log.Info(ctx).Msgf("Starting %s on port %d", config.Conf.Server.AppName, config.Conf.Server.ServicePort)

	// start server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Conf.Server.ServicePort),
		Handler: router,
	}

	// handle graceful shutdown
	idleConnectionsClosed := make(chan struct{})

	go func() {

		sigint := make(chan os.Signal, 1)

		signal.Notify(sigint, os.Interrupt)
		signal.Notify(sigint, syscall.SIGTERM)

		<-sigint

		log.Info(ctx).Msgf("%s shutdown initiated", config.Conf.Server.AppName)

		if err := server.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Error(ctx).Err(err).Msgf("unable to gracefully shutdown https server for %s", config.Conf.Server.AppName)
		}

		log.Info(ctx).Msgf("%s shutdown complete", config.Conf.Server.AppName)

		close(idleConnectionsClosed)
	}()

	// print available endpoints
	printEndpoints(ctx, router)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(ctx).Err(err).Msgf("%s startup failed.", config.Conf.Server.AppName)
	}

	<-idleConnectionsClosed
}
