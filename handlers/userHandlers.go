package handlers

import (
	"github.com/andreyberezin/gin-site/models"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func ShowLoginPage(context *gin.Context) {
	context.HTML(
		http.StatusOK,
		"login.html",
		gin.H{
			"title": "Login",
		},
	)
}

func PerformLogin(context *gin.Context) {
	session := sessions.Default(context)
	password := context.PostForm("password")
	email := context.PostForm("email")
	email = strings.ToLower(email)

	if models.IsUserValid(email, password) {
		session.Set("user-id", email)
		err := session.Save()
		if err != nil {
			context.HTML(http.StatusBadRequest, "login.html", gin.H{
				"ErrorTitle":   "Login Failed",
				"ErrorMessage": "Failed to generate session token"})
		} else {
			context.HTML(http.StatusOK, "login-successful.html",
				gin.H{
					"title": "Successful Login",
				},
			)
		}
	} else {
		context.HTML(http.StatusBadRequest, "login.html", gin.H{
			"ErrorTitle":   "Login Failed",
			"ErrorMessage": "Invalid credentials provided"})
	}
}

func Logout(context *gin.Context) {
	// Clear the cookie
	session := sessions.Default(context)
	session.Delete("user-id")
	err := session.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
		return
	}
	// Redirect to the home page
	context.Redirect(http.StatusTemporaryRedirect, "/")
}

func Register(context *gin.Context) {
	// Obtain the POSTed username and password values
	username := context.PostForm("username")
	password := context.PostForm("password")
	email := context.PostForm("email")
	email = strings.ToLower(email)

	if _, err := models.RegisterNewUser(email, username, password); err == nil {
		// If the user is created, set the token in a cookie and log the user in
		session := sessions.Default(context)
		session.Set("user-id", email)
		err := session.Save()
		if err != nil {
			context.HTML(http.StatusBadRequest, "login.html",
				gin.H{
					"ErrorTitle":   "Login Failed",
					"ErrorMessage": "Failed to generate session token"})
		} else {
			context.HTML(http.StatusOK, "login-successful.html",
				gin.H{
					"title": "Successful registration & Login",
				},
			)
		}
	} else {
		// If the username/password combination is invalid,
		// show the error message on the login page
		context.HTML(http.StatusBadRequest, "login.html",
			gin.H{
				"ErrorTitle":   "Registration Failed",
				"ErrorMessage": err.Error()})
	}
}
