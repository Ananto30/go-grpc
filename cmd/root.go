package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	// RootCmd is the root command of promo service
	RootCmd = &cobra.Command{
		Use:   "fortress",
		Short: "fortress is a grpc service",
		Long:  `An gRPC API backend for campaign management`,
	}
)

// Execute executes the root command
func Execute() {

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
