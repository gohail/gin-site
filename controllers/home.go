package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//GetHome handles GET / route
func GetHome(c *gin.Context) {
	h := DefaultH(c)
	h["Title"] = "Welcome to Desk service"
	c.HTML(http.StatusOK, "home.html", h)
}
