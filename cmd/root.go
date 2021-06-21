package cmd

import (
	"car_scraper/database"
	"car_scraper/models"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"path/filepath"
	"runtime"

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
	_, b, _, _ := runtime.Caller(0)
	err := godotenv.Overload(filepath.Join(filepath.Dir(b), "../.env"))
	if err != nil {
		log.Fatalln(err)
	}

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
