package models

import (
	"errors"
	"github.com/Sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

//User type contains user basic info
type User struct {
	Model
	UserName  string
	Password  string
	UserEmail string `gorm:"unique_index"`

	ExtraUserInfo ExtraUserInfo `gorm:"ForeignKey:InfoRefer"` // Has-One relationship
	InfoRefer     uint

	Adverts []Advert `gorm:"ForeignKey:UserID"` // One-To-Many relationship
}

type ExtraUserInfo struct {
	Model
	ContactEmail string `form:"email"`
	PhoneNumber  string `form:"phone"`
	AboutMe      string `form:"about"`
	Photo        string `form:"photo"`
}

//Login view model
type Login struct {
	Email    string `form:"email" binding:"required" validate:"required,email"`
	Password string `form:"password" binding:"required" validate:"required,max=8,min=4"`
}

//Register view model
type Register struct {
	Name     string `form:"name" binding:"required" validate:"required,max=20,min=4"`
	Email    string `form:"email" binding:"required" validate:"required,email"`
	Password string `form:"password" binding:"required" validate:"required,max=10,min=4"`
}

func (r *Register) RemovePass() {
	r.Password = ""
}

func (l *Login) RemovePass() {
	l.Password = ""
}

func IsUserValid(email, password string) (uint64, bool) {
	db := GetBD()
	u := User{}
	db.Where("user_email = lower(?)", email).First(&u)
	if u.ID == 0 || bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) != nil {
		return 0, false
	}
	return u.ID, true
}

func RegisterNewUser(r *Register) (*User, error) {
	if strings.TrimSpace(r.Password) == "" {
		return nil, errors.New("пароль не может быть пустым")
	}
	if !isUserEmailAvailable(r.Email) {
		return nil, errors.New("этот email уже занят")
	}
	db := GetBD()
	u := User{}
	u.UserName = r.Name
	u.Password = r.Password
	u.UserEmail = r.Email
	if err := db.Create(&u).Error; err != nil {
		logrus.Error(err)
		return nil, errors.New("ошибка регистрации пользователя")
	}
	return &u, nil
}

func isUserEmailAvailable(userEmail string) bool {
	db := GetBD()
	user := User{}
	db.Where("user_email = ?", userEmail).First(&user)
	if user.ID != 0 {
		return false
	}
	return true
}

//BeforeSave gorm hook
func (u *User) BeforeSave() (err error) {
	var hash []byte
	hash, err = bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	u.Password = string(hash)
	return
}
