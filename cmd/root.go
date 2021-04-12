package cmd

import (
	"fmt"
	"os"

	"car_scraper/server"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "start:server",
	Short: "Start the webserver",
	Long:  "Start the webserver using the provided configuration",
	Run: func(cmd *cobra.Command, args []string) {
		server.StartServer()
	},
}

func Init() {
	rootCmd.AddCommand(StartScraperCommand)
}

func Execute() {
	Init()

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
