package middlewere

import (
	"github.com/andreyberezin/gin-site/controllers"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

// This middleware ensures that a request will be aborted with an error if the user is not logged in
func EnsureLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := sessions.Default(c)
		user := s.Get(controllers.USER_ID)
		if user != nil {
			c.Next()
		} else {
			s.AddFlash("Пожалуйста авторизируйтесь.")
			s.Save()
			c.Redirect(http.StatusFound, "/auth/login")
			c.Abort()
		}
	}
}

// This middleware ensures that a request will be aborted with an error if the user is already logged in
func EnsureNotLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := sessions.Default(c)
		user := s.Get(controllers.USER_ID)
		if user == nil {
			c.Next()
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
