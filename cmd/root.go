package cmd

import (
	"car_scraper/database"
	"car_scraper/models"
	"car_scraper/server"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "start:server",
	Short: "Start the webserver",
	Long:  "Start the webserver using the provided configuration",
	Run: func(cmd *cobra.Command, args []string) {
		gin.SetMode(os.Getenv("GIN_MODE"))
		server.StartServer()
	},
}

func Init() {
	database.ConnectToDatabase()
	models.InitiateModels()


	rootCmd.AddCommand(StartScraperCommand)
	rootCmd.AddCommand(RecreateUserFiltersCommand)
}

func Execute() {
	Init()

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
