package cmd

import (
	"car_scraper/models"
	"car_scraper/scraper"
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"github.com/gammazero/workerpool"
	"github.com/spf13/cobra"
	"strconv"
)

var RecreateUserFiltersCommand = &cobra.Command{
	Use:   "filter:recreate [userId]",
	Short: "Recreate a user's filters",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			panic("No user id has been provided")
		}
		intVal, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			panic("Invalid username provided")
		}
		userId := uint8(intVal)
		recreateUserFilters(&userId)
	},
}

func recreateUserFilters(userId *uint8) {
	userRepo := models.UserRepository{}

	user := userRepo.GetUserById(*userId)

	wp := workerpool.New(20)
	bar := pb.Full.Start(len(user.Filters))
	for _, filter := range user.Filters {
		filter := filter
		wp.Submit(func() {
			filterRepo := models.FilterRepository{}
			s, err := scraper.GetScraper(filter.Type)
			if err != nil {
				panic(fmt.Sprintf("Failed to get scraper for filter %v %v", filter.ID, err))
			}

			newFilter := &models.Filter{
				UserID:    filter.UserID,
				User:      filter.User,
				Type:      filter.Type,
				Search:    filter.Search,
				Cars:      nil,
			}
			err = s.InitiateFilter(newFilter)
			if err != nil {
				panic(fmt.Sprintf("Failed to recrete filter %v %v: JSON %v", filter.ID, err, filter.Search))
			}

			err = filterRepo.SaveFilter(newFilter)
			if err != nil {
				panic(fmt.Sprintf("Failed to save recreated filter %v %v", filter.ID, err))
			}

			err = filterRepo.DeleteFilter(&filter)
			if err != nil {
				panic(fmt.Sprintf("Failed to delete filter %v %v", filter.ID, err))
			}

			bar.Increment()
		})
	}
	wp.StopWait()
	bar.Finish()
}