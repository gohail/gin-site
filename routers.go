package main

import (
	"github.com/andreyberezin/gin-site/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func initializeRoutes() {
	router.GET("/", showIndexPage)
	router.GET("/article/view/:article_id", getArticleByID)
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
