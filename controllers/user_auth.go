package controllers

import (
	"github.com/andreyberezin/gin-site/models"
	"github.com/andreyberezin/gin-site/validator"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowLoginPage(c *gin.Context) {
	h := DefaultH(c)
	h["Title"] = "Basic GIN web-site signin form"
	session := sessions.Default(c)
	h["Flash"] = session.Flashes()
	login := session.Get("login")
	if login != nil {
		h["Login"] = login
		session.Delete("login")
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
		c.Redirect(http.StatusFound, "/auth/login")
		return
	}

	if isValid := validator.LoginValidate(&login, session); !isValid {
		login.RemovePass()
		session.Set("login", login)
		session.Save()
		c.Redirect(http.StatusFound, "/auth/login")
		return
	}

	ID, ok := models.IsUserValid(login.Email, login.Password)
	if !ok {
		session.AddFlash("Ошибка авторизации. E-MAIL АДРЕС или ПАРОЛЬ введен не верно.")
		session.Set("login", login)
		session.Save()
		c.Redirect(http.StatusFound, "/auth/login")
		return
	}
	session.Set(USER_ID, ID)
	session.Save()
	c.Redirect(http.StatusFound, "/baraholka")
}

func Logout(c *gin.Context) {
	// Clear the cookie
	session := sessions.Default(c)
	session.Delete(USER_ID)
	session.Save()
	c.Redirect(http.StatusSeeOther, "/")
}

func PostRegister(c *gin.Context) {
	session := sessions.Default(c)
	register := models.Register{}
	if err := c.ShouldBind(&register); err != nil {
		session.AddFlash("Пожалуйста, заполните пустые поля!")
		register.RemovePass()
		session.Set("register", register)
		session.Save()
		c.Redirect(http.StatusFound, "/auth/login")
		return
	}
	if isValid := validator.RegisterValidate(&register, session); !isValid {
		register.RemovePass()
		session.Set("register", register)
		session.Save()
		c.Redirect(http.StatusFound, "/auth/login")
		return
	}

	u, err := models.RegisterNewUser(&register)
	if err != nil {
		session.AddFlash(err.Error())
		register.RemovePass()
		session.Set("register", register)
		session.Save()
		c.Redirect(http.StatusFound, "/auth/login")
		return
	}
	session.Set(USER_ID, u.ID)
	session.Save()
	c.Redirect(http.StatusFound, "/baraholka")
}
