package bootstrap

import (
	"fmt"
	"os"

	"github.com/Blockchain-Framework/controller/cmd/godsight_controller"
	"github.com/Blockchain-Framework/controller/internal/godsight-aggregator/config"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "aggregator",
	Short: "AppDirector Aggregator Service",
	Run: func(cmd *cobra.Command, args []string) {

		log := zerolog.New(os.Stderr).With().Timestamp().Logger()

		// load .env
		err := godotenv.Load()
		if err != nil {
			log.Warn().Err(err).Msg("No .env file found. [only required for local development]")
		}

		// create configuration using the environment variables
		err = config.LoadFromEnvironment()
		if err != nil {
			log.Panic().Err(err).Msg("Failed to configure service")
		}

		log.Info().Msgf("loaded configuration. %+v", config.Conf)

		// setup readiness liveness probe file
		// liveFile, err := util.Create()

		// if err != nil {
		// 	log.Panic().Err(err).Msg("Unable to create readiness liveness probe file")
		// }

		// log.Info().Msgf("Created liveness probe file, %s", liveFile.Name())

		// start service
		godsight_controller.Start(config.Conf)

		// clean up readiness liveness probe file
		// if err := util.Remove(); err != nil {
		// 	log.Panic().Err(err).Msg("Unable to remove readiness liveness probe file")
		// }
	},
}

func Execute() {

	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
