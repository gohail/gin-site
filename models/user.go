package models

import (
	"errors"
	"strings"
)

type User struct {
	Model

	UserName string `form:"name"`
	Password string `form:"password"`
}

func IsUserValid(name, password string) bool {
	db := GetBD()
	u := User{}
	db.Where("user_name = ? AND password = ?", name, password).First(&u)
	if u.ID == 0 {
		return false
	}
	return true
}

func RegisterNewUser(username, password string) (*User, error) {
	if strings.TrimSpace(password) == "" {
		return nil, errors.New("The password can't be empty")
	}
	if !isUserNameAvailable(username) {
		return nil, errors.New("The username isn't available")
	}
	db := GetBD()
	u := User{}
	u.UserName = username
	u.Password = password
	if err := db.Create(&u).Error; err != nil {
		return nil, errors.New("Error whilst registering user")
	}
	return &u, nil
}

func isUserNameAvailable(username string) bool {
	db := GetBD()
	user := User{}
	db.Where("user_name = ?", username).First(&user)
	if user.ID != 0 {
		return false
	}
	return true
}
