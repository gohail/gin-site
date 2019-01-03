package handlers

import (
	"github.com/andreyberezin/gin-site/controllers"
	"github.com/andreyberezin/gin-site/models"
	"github.com/andreyberezin/gin-site/validator"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	v8 "gopkg.in/go-playground/validator.v8"
	"net/http"
	"strings"
)

func ShowLoginPage(c *gin.Context) {
	h := controllers.DefaultH(c)
	h["Title"] = "Basic GIN web-site signin form"
	session := sessions.Default(c)
	h["Flash"] = session.Flashes()
	login := session.Get("login")
	if login != nil {
		h["Login"] = login
	}
	session.Save()
	c.HTML(http.StatusOK, "login.html", h)
}

func PostLogin(c *gin.Context) {
	session := sessions.Default(c)
	login := models.Login{}
	if err := c.ShouldBind(&login); err != nil {
		session.AddFlash("Пожалуйста, Заполните пустые поля.")
		session.Save()
		c.Redirect(http.StatusFound, "/u/login")
		return
	}
	if errs := validator.Validate.Struct(login); errs != nil {
		er := errs.(v8.ValidationErrors)
		if mail := er["Login.Email"]; mail != nil {
			session.AddFlash("Email введен неверно.")
		}
		if pass := er["Login.Password"]; pass != nil {
			session.AddFlash("Пароль должен быть от 4 до 8 символов.")
		}
		session.Set("login", login)
		session.Save()
		c.Redirect(http.StatusFound, "/u/login")
		return
	}

	if ID, ok := models.IsUserValid(login.Email, login.Password); !ok {
		session.AddFlash("Ошибка авторизации. E-MAIL АДРЕС или ПАРОЛЬ введен не верно.")
		session.Set("login", login)
		session.Save()
		c.Redirect(http.StatusFound, "/u/login")
		return
	} else {
		session.Set(controllers.USER_ID, ID)
		session.Save()
		c.Redirect(http.StatusOK, "login-successful.htm")
	}
}

func Logout(context *gin.Context) {
	// Clear the cookie
	session := sessions.Default(context)
	session.Delete(controllers.USER_ID)
	session.Save()
	context.Redirect(http.StatusFound, "/")
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
