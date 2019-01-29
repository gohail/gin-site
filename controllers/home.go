package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetHome handles GET / route
func GetHome(c *gin.Context) {
	h := DefaultH(c)
	h["Title"] = "Welcome to Desk service"
	c.HTML(http.StatusOK, "index.html", h)
}

//GetSearch handles GET /search route
func GetSearch(c *gin.Context) {
	h := DefaultH(c)
	c.HTML(http.StatusOK, "search.html", h)
}
