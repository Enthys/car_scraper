package server

import (
	"car_scraper/server/authentication"
	"car_scraper/server/filter"
	"car_scraper/server/middleware"
	"car_scraper/server/user"
	"log"

	"car_scraper/auth"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	auth.InitAuth()

	app := gin.Default()

	setupRoutes(app)

	err := app.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func setupRoutes(app *gin.Engine) {
	filterRoutes := app.Group("/filters")
	{
		filterRoutes.GET("", middleware.AuthenticateUser(), filter.GetCarFilters)
		filterRoutes.POST("", middleware.AuthenticateUser(), filter.CreateCarFilter)
		filterRoutes.DELETE(":filterId", middleware.AuthenticateUser(), filter.DeleteCarFilter)
	}

	authRoutes := app.Group("/auth")
	{
		authRoutes.POST("login", authentication.Login)
		authRoutes.GET("logout", authentication.Logout)
	}

	userRoutes := app.Group("/users")
	{
		userRoutes.POST("", user.RegisterUser)
	}

	carsBGRoutes := app.Group("/carsbg")
	{
		carsBGRoutes.GET(":type/brands/:brandId/models", middleware.AuthenticateUser(), filter.GetCarsBGBrandModels)
	}
}
