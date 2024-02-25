package bootstrap

import (
	"os"

	"github.com/Blockchain-Framework/controller/pkg/util"

	"github.com/spf13/cobra"
)

var liveCmd = &cobra.Command{
	Use:   "live",
	Short: "Check if application is live",
	Run: func(cmd *cobra.Command, args []string) {

		if util.Exists() {
			os.Exit(0)
			return
		}

		os.Exit(1)
	},
}

func init() {
	rootCmd.AddCommand(liveCmd)
}
