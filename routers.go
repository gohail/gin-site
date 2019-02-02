package main

import (
	"github.com/andreyberezin/gin-site/controllers"
	"github.com/andreyberezin/gin-site/middlewere"
)

func initializeRoutes() {

	// Handle the GET requests at /
	router.GET("/", controllers.GetHome)

	// Handle the GET requests at /search
	router.GET("/search", controllers.GetSearch)

	authRoutes := router.Group("/auth")
	{
		// Handle the GET requests at /u/login
		// Show the login page
		// Ensure that the user is not logged in by using the middleware
		authRoutes.GET("/login", controllers.ShowLoginPage)

		// Handle POST requests at /u/login
		// Ensure that the user is not logged in by using the middleware
		authRoutes.POST("/login", controllers.PostLogin)

		// Handle GET requests at /u/logout
		// Ensure that the user is logged in by using the middleware
		authRoutes.GET("/logout", controllers.Logout)

		// Handle POST requests at /u/register
		// Ensure that the user is not logged in by using the middleware
		authRoutes.POST("/register", middlewere.EnsureNotLoggedIn(), controllers.PostRegister)
	}

	advertRoutes := router.Group("/baraholka")
	{
		// Handle GET requests at /baraholka
		advertRoutes.GET("/", controllers.BaraholkaMain)

		// Handle GET requests at /baraholka/all
		advertRoutes.GET("/all", controllers.GetAllAdverts)

		// Handle GET requests at /baraholka/:id
		//dvertRoutes.GET("/show/:id", controllers.AdvertGet)

		// Handle GET requests /adverts/new_adv
		//advertRoutes.GET("/new_adv")
	}

	userRouters := router.Group("/user")
	userRouters.Use(middlewere.EnsureLoggedIn())
	{
		// Handle GET request /users/myinfo
		userRouters.GET("/myinfo", controllers.GetSelfInfo)

		// Handle POST request /users/myinfo receive ExtraUserInfo{}
		userRouters.POST("/myinfo", controllers.PostSelfInfo)
	}

	adminRouters := router.Group("/admin")
	{
		adminRouters.GET("/", middlewere.EnsureLoggedIn(), controllers.AdminTestGet)
	}
}
