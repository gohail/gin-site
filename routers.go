package main

import (
	"github.com/andreyberezin/gin-site/handlers"
	"github.com/andreyberezin/gin-site/middlewere"
	"github.com/andreyberezin/gin-site/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func initializeRoutes() {
	router.GET("/", showIndexPage)
	router.GET("/article/view/:article_id", getArticleByID)

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
}

func showIndexPage(context *gin.Context) {
	context.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title":   "Home Page",
			"payload": models.GetAllArticles(),
		},
	)
}

func getArticleByID(context *gin.Context) {
	if ID, err := strconv.Atoi(context.Param("article_id")); err == nil {
		if article, err := models.GetArticleByID(ID); err == nil {
			context.HTML(
				http.StatusOK,
				"article.html",
				gin.H{
					"payload": article,
				},
			)
		} else {
			context.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		context.AbortWithStatus(http.StatusNotFound)
	}
}
