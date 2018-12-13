package main

import (
	"github.com/andreyberezin/gin-site/models"
	"github.com/andreyberezin/gin-site/system"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	system.LoadDbConfig()
	models.SetDB(system.GetConnectionString())
	models.AutoMigrate()

	router = gin.Default()
	store := sessions.NewCookieStore([]byte("secret"))
	router.Use(sessions.Sessions("gosession", store))
	router.LoadHTMLGlob("templates/*")

	initializeRoutes()

	router.Run()
}
