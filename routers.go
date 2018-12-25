package main

import (
	"github.com/andreyberezin/gin-site/controllers"
	"github.com/andreyberezin/gin-site/handlers"
	"github.com/andreyberezin/gin-site/middlewere"
)

func initializeRoutes() {

	// Handle the GET requests at /
	router.GET("/", controllers.GetHome)

	userRoutes := router.Group("/u")
	{
		// Handle the GET requests at /u/login
		// Show the login page
		// Ensure that the user is not logged in by using the middleware
		userRoutes.GET("/login", middlewere.EnsureNotLoggedIn(), handlers.ShowLoginPage)

		// Handle POST requests at /u/login
		// Ensure that the user is not logged in by using the middleware
		userRoutes.POST("/login", middlewere.EnsureNotLoggedIn(), handlers.PerformLogin)

		// Handle GET requests at /u/logout
		// Ensure that the user is logged in by using the middleware
		userRoutes.GET("/logout", middlewere.EnsureLoggedIn(), handlers.Logout)

		// Handle the GET requests at /u/register
		// Show the registration page
		// Ensure that the user is not logged in by using the middleware
		userRoutes.GET("/register", middlewere.EnsureNotLoggedIn(), handlers.ShowRegistrationPage)

		// Handle POST requests at /u/register
		// Ensure that the user is not logged in by using the middleware
		userRoutes.POST("/register", middlewere.EnsureNotLoggedIn(), handlers.Register)
	}

	advertRoutes := router.Group("/baraholka")
	{
		// Handle GET requests at /adverts/all
		advertRoutes.GET("/all", controllers.GetAllAdverts)

		// Handle GET requests at /adverts/:id
		advertRoutes.GET("/show/:id", controllers.AdvertGet)

		// Handle GET requests /adverts/new_adv
		//advertRoutes.GET("/new_adv")

		// Handle POST requests /adverts/new_adv
		//advertRoutes.POST("/new_adv")
	}
}
