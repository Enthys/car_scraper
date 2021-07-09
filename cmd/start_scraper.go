package cmd

import (
	"bytes"
	"car_scraper/models"
	"car_scraper/scraper"
	"fmt"
	"github.com/gammazero/workerpool"
	"github.com/spf13/cobra"
	"html/template"
	"log"
	"net/smtp"
	"os"
	"strconv"
	"sync"
)

var StartScraperCommand = &cobra.Command{
	Use:   "start:scraper",
	Short: "Start scraping through filters",
	Run: func(cmd *cobra.Command, args []string) {
		users := models.UserRepository{}.GetUsers()
		for _, user := range users {
			runUserFilters(&user)
		}
	},
}

func runUserFilters(user *models.User) {
	workerCount, err := strconv.ParseInt(os.Getenv("WORKER_COUNT"), 10, 16)
	if err != nil {
		workerCount = 100
	}
	wp := workerpool.New(int(workerCount))

	newCarHolder := sync.Map{}
	for _, filter := range user.Filters {
		filterID := filter.ID

		wp.Submit(func() {
			cars, _ := runFilter(&filterID)
			for _, car := range cars {
				newCarHolder.Store(car.Link, car)
			}
		})
	}

	wp.StopWait()
	newCars := make([]models.CarDTO, 0)
	newCarHolder.Range(func(key, value interface{}) bool {
		newCars = append(newCars, value.(models.CarDTO))

		return true
	})

	if len(newCars) != 0 {
		sendEmailsForNewCars(newCars, user.Email)
	}
}

func runFilter(filterID *uint32) ([]models.CarDTO, error) {
	filterRepo := models.FilterRepository{}
	carRepo := models.CarRepository{}

	filter := filterRepo.GetFilterByID(*filterID)
	filterScraper, err := scraper.GetScraper(filter.Type)
	if err != nil {
		log.Printf("Could not get scraper for filter '%v' type '%s'", filter.ID, filter.Type)
		return nil, err
	}

	result := make([]models.CarDTO, 0)
	newCars := filterScraper.GetNewCars(&filter)
	for _, key := range newCars.Keys() {
		var carDTO models.CarDTO
		carVal, _ := newCars.Get(key)
		carDTO = carVal.(models.CarDTO)

		err = carRepo.SaveCar(&models.Car{
			FilterID: filter.ID,
			Filter:   filter,
			Link:     carDTO.Link,
		})

		if err != nil {
			log.Println("Failed to save car: ", carDTO.Link)
		}
	}

	for _, key := range newCars.Keys() {
		val, _ := newCars.Get(key)
		result = append(result, val.(models.CarDTO))
	}

	return result, nil
}

func sendEmailsForNewCars(cars []models.CarDTO, receiverEmail string) {
	from := os.Getenv("MAIL_SENDER_EMAIL")
	pass := os.Getenv("MAIL_SENDER_PASSWORD")
	mailAddr := os.Getenv("MAIL_SENDER_ADDRESS")
	mailHost := os.Getenv("MAIL_SENDER_HOST")

	println("Sending email with ", from, pass, mailHost, mailAddr)

	gmailAuth := smtp.PlainAuth("", from, pass, mailHost)

	tmpl, err := template.ParseFiles("views/new-cars.html")
	if err != nil {
		panic(err)
	}

	var body bytes.Buffer
	headers := "Content-Type: text/html; charset=\"ISO-8859-1\""
	body.Write([]byte(fmt.Sprintf("Subject: New Cars\n%s\n\n", headers)))

	tmpl.Execute(&body, struct {
		Cars []models.CarDTO
	}{
		Cars: cars,
	})

	err = smtp.SendMail(mailAddr, gmailAuth, from, []string{receiverEmail}, body.Bytes())
	if err != nil {
		log.Fatal(err)
	}
}
