package validator

import (
	"github.com/andreyberezin/gin-site/models"
	"github.com/gin-gonic/contrib/sessions"
	"gopkg.in/go-playground/validator.v8"
)

const (
	nameSizeError     = "Имя должен быть от 4 до 20 символов."
	passwordSizeError = "Пароль должен быть от 4 до 8 символов."
	emailSyntaxError  = "Email введен неверно."
)

var Validate *validator.Validate

func SetValidator(tagName string) {
	config := &validator.Config{TagName: tagName}
	Validate = validator.New(config)
}

// Validate Register{} struct field and add Flashes in Session, return true if valid
func RegisterValidate(model *models.Register, s sessions.Session) bool {
	if errs := Validate.Struct(model); errs != nil {
		er := errs.(validator.ValidationErrors)
		if mail := er["Register.Email"]; mail != nil {
			s.AddFlash(emailSyntaxError)
		}
		if name := er["Register.Name"]; name != nil {
			s.AddFlash(nameSizeError)
		}
		if pass := er["Register.Password"]; pass != nil {
			s.AddFlash(passwordSizeError)
		}
		s.Save()
		return false
	}
	return true
}

// Validate Login{} struct field and add Flashes in Session, return true if valid

func LoginValidate(model *models.Login, s sessions.Session) bool {
	if errs := Validate.Struct(model); errs != nil {
		er := errs.(validator.ValidationErrors)
		if mail := er["Login.Email"]; mail != nil {
			s.AddFlash(emailSyntaxError)
		}
		if pass := er["Login.Password"]; pass != nil {
			s.AddFlash(passwordSizeError)
		}
		s.Save()
		return false
	}
	return true
}
