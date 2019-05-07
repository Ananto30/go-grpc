package cmd

import (
	"github.com/Ananto30/go-grpc/config"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print server version",
	Long:  `Print the version of promo backend server`,
	Run:   version,
}

func init() {
	RootCmd.AddCommand(versionCmd)
}

func version(cmd *cobra.Command, args []string) {
	config.LoadApp()
	app := config.App()
	println("Promo Service: " + app.Version + " Env: " + app.Env)
}
