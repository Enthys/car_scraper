package server

import (
	"car_scraper/server/middleware"
	"log"

	"car_scraper/auth"
	"car_scraper/database"
	"car_scraper/models"
	authActions "car_scraper/server/authentication/routes"
	filterActions "car_scraper/server/filter/routes"
	userActions "car_scraper/server/user/routes"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	database.ConnectToDatabase()
	models.InitiateModels()
	auth.InitAuth()

	app := gin.Default()

	setupRoutes(app)

	err := app.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func setupRoutes(app *gin.Engine) {
	filterRoutes := app.Group("/api/filters")
	{
		filterRoutes.GET("", middleware.AuthenticateUser(), filterActions.GetCarFilters)
		filterRoutes.POST("", middleware.AuthenticateUser(), filterActions.CreateCarFilrer)
		filterRoutes.DELETE("", middleware.AuthenticateUser(), filterActions.DeleteCarFilter)
	}

	authRoutes := app.Group("/api/auth")
	{
		authRoutes.POST("login", authActions.Login)
		authRoutes.GET("logout", authActions.Logout)
	}

	userRoutes := app.Group("/api/users")
	{
		userRoutes.POST("", userActions.RegisterUser)
	}
}
