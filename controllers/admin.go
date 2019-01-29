package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AdminTestGet(c *gin.Context) {
	h := DefaultH(c)
	h["Title"] = "Admin page"
	c.HTML(http.StatusOK, "admin_test.html", h)
}
