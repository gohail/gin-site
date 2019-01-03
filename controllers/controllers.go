package controllers

import "github.com/gin-gonic/gin"

const USER_ID = "user-id"

func DefaultH(c *gin.Context) gin.H {
	return gin.H{
		"Title":   "",
		"Context": c,
	}
}
