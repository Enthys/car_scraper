package cmd

import "github.com/spf13/cobra"

var StartScraperCommand = &cobra.Command{
	Use:   "start:scraper",
	Short: "Start scraping through filters",
	Run: func(cmd *cobra.Command, args []string) {
		println("scraping cars . . .")
	},
}
