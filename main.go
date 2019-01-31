package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/andreyberezin/gin-site/models"
	"github.com/andreyberezin/gin-site/system"
	"github.com/andreyberezin/gin-site/validator"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	initLogger()
	system.LoadDbConfig()
	models.SetDB(system.GetConnectionString())
	models.AutoMigrate()
	validator.SetValidator("validate")
	system.InitModelsForSession()

	router = gin.Default()
	router.Use(static.Serve("/public", static.LocalFile("./public", true)))
	store := sessions.NewCookieStore([]byte("secret"))
	router.Use(sessions.Sessions("gosession", store))
	router.LoadHTMLGlob("view/*/*.html")

	initializeRoutes()

	router.Run()
}

func initLogger() {
	logrus.SetFormatter(&logrus.TextFormatter{})
	//logrus.SetOutput(os.Stderr)
	if gin.Mode() == gin.DebugMode {
		logrus.SetLevel(logrus.DebugLevel)
	}
}
