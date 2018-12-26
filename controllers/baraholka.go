package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func BaraholkaMain(c *gin.Context) {
	h := DefaultH(c)
	c.HTML(http.StatusOK, "baraholka.html", h)
}
