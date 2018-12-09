package middlewere

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

// This middleware ensures that a request will be aborted with an error
// if the user is not logged in
func EnsureLoggedIn() gin.HandlerFunc {
	return func(context *gin.Context) {
		// If there's an error or if the token is empty
		// the user is not logged in
		session := sessions.Default(context)
		user := session.Get("user-id")
		if user == nil {
			//context.JSON(http.StatusBadRequest, gin.H{"error": "Already logged"})
			context.AbortWithStatus(http.StatusUnauthorized)
		}
		//c.AbortWithStatus(http.StatusUnauthorized)
	}
}

// This middleware ensures that a request will be aborted with an error
// if the user is already logged in
func EnsureNotLoggedIn() gin.HandlerFunc {
	return func(context *gin.Context) {
		session := sessions.Default(context)
		user := session.Get("user-id")
		if user != nil {
			//context.JSON(http.StatusBadRequest, gin.H{"error": "Already logged"})
			context.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
