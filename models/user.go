package models

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

type User struct {
	Model
	UserName  string `form:"name"`
	Password  string `form:"password" binding:"required"`
	UserEmail string `form:"email" binding:"required" gorm:"unique_index"`

	Adverts []Advert `gorm:"ForeignKey:UserID"` // One-To-Many relationship
}

func IsUserValid(email, password string) bool {
	db := GetBD()
	u := User{}
	db.Where("user_email = lower(?)", email).First(&u)
	if u.ID == 0 || bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) != nil {
		return false
	}
	return true
}

func RegisterNewUser(email, username, password string) (*User, error) {
	if strings.TrimSpace(password) == "" {
		return nil, errors.New("The password can't be empty")
	}
	if !isUserEmailAvailable(email) {
		return nil, errors.New("The username isn't available")
	}
	db := GetBD()
	u := User{}
	u.UserName = username
	u.Password = password
	u.UserEmail = email
	if err := db.Create(&u).Error; err != nil {
		return nil, errors.New("Error whilst registering user")
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
