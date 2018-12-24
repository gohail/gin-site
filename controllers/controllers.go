package controllers

import "github.com/gin-gonic/gin"

func DefaultH(c *gin.Context) gin.H {
	return gin.H{
		"Title":   "",
		"Context": c,
	}
}
