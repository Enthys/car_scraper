package cmd

import (
	"car_scraper/models"
	"github.com/spf13/cobra"
)

var ClearOldCarsCommand = &cobra.Command{
	Use:   "scraper:clear_old_cars",
	Short: "Delete old cars from the database",
	Run: func(cmd *cobra.Command, args []string) {
		userRepo := models.UserRepository{}
		users := userRepo.GetUsers()

		carRepo := models.CarRepository{}
		filterRepo := models.FilterRepository{}
		for _, user := range users {
			for _, filter := range user.Filters {
				filterRepo.DeleteOldCars(&carRepo, &filter)
			}
		}
	},
}
