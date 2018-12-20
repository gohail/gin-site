package main

import (
	"github.com/andreyberezin/gin-site/models"
	"github.com/andreyberezin/gin-site/system"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	system.LoadDbConfig()
	models.SetDB(system.GetConnectionString())
	models.AutoMigrate()

	router = gin.Default()
	router.Use(static.Serve("/public", static.LocalFile("./public", true)))
	store := sessions.NewCookieStore([]byte("secret"))
	router.Use(sessions.Sessions("gosession", store))
	router.LoadHTMLGlob("view/*/*.html")

	initializeRoutes()

	router.Run()
}
