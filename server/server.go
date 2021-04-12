package server

import (
	"car_scraper/database"
	authActions "car_scraper/server/authentication/routes"
	filterActions "car_scraper/server/filter/routes"
	userActions "car_scraper/server/user/routes"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	database.ConnectToDatabase()
	database.InitiateModels()

	app := gin.Default()

	setupRoutes(app)

	app.Run()
}

func setupRoutes(app *gin.Engine) {
	filterRoutes := app.Group("/api/filters")
	{
		filterRoutes.GET("", filterActions.GetCarFilters)
		filterRoutes.POST("", filterActions.CreateCarFilrer)
		filterRoutes.DELETE("", filterActions.DeleteCarFilter)
	}

	autheRoutes := app.Group("/api/auth")
	{
		autheRoutes.POST("login", authActions.Login)
		autheRoutes.GET("logout", authActions.Logout)
	}

	userRoutes := app.Group("/api/user")
	{
		userRoutes.POST("", userActions.RegisterUser)
	}
}
