package server

import (
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

	app.Run()
}

func setupRoutes(app *gin.Engine) {
	filterRoutes := app.Group("/api/filters")
	{
		filterRoutes.GET("", auth.AuthenticateUser(), filterActions.GetCarFilters)
		filterRoutes.POST("", auth.AuthenticateUser(), filterActions.CreateCarFilrer)
		filterRoutes.DELETE("", auth.AuthenticateUser(), filterActions.DeleteCarFilter)
	}

	autheRoutes := app.Group("/api/auth")
	{
		autheRoutes.POST("login", authActions.Login)
		autheRoutes.GET("logout", authActions.Logout)
	}

	userRoutes := app.Group("/api/users")
	{
		userRoutes.POST("", userActions.RegisterUser)
	}
}
