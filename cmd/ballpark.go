package cmd

import (
	"github.com/rms1000watt/degeneres-test/ballpark"
	"github.com/rms1000watt/degeneres-test/server"
	"github.com/spf13/cobra"
)

var (
	ballparkCfg ballpark.Config
	serverCfg   server.Config
)

// ballparkCmd represents the ballpark command
var ballparkCmd = &cobra.Command{
	Use:   "ballpark",
	Short: "Ballpark Service API for stadium information",
	Long:  ``,
	Run:   runBallpark,
}

func init() {
	RootCmd.AddCommand(ballparkCmd)

	ballparkCmd.Flags().StringVar(&serverCfg.Host, "host", "0.0.0.0", "Host address for server")
	ballparkCmd.Flags().IntVar(&serverCfg.Port, "port", 8080, "Port for server")

	ballparkCmd.Flags().StringVar(&serverCfg.CertsPath, "certs-path", "", "Path for certs")
	ballparkCmd.Flags().StringVar(&serverCfg.KeyName, "key-name", "", "Private key name in certs path")
	ballparkCmd.Flags().StringVar(&serverCfg.CertName, "cert-name", "", "Public key name in certs path")

	SetFlagsFromEnv(ballparkCmd)
}

func runBallpark(cmd *cobra.Command, args []string) {
	configureLogging()

	server.Ballpark(serverCfg, ballparkCfg)
}
