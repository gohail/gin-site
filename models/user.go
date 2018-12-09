package models

import (
	"errors"
	"strings"
)

type User struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

var userList = []User{
	{"user1", "pass1"},
	{"user2", "pass2"},
	{"user3", "pass3"},
	{"user4", "pass4"},
	{"admin", "admin"},
}

func IsUserValid(name, password string) bool {
	for _, u := range userList {
		if u.UserName == name && u.Password == password {
			return true
		}
	}
	return false
}

func RegisterNewUser(username, password string) (*User, error) {
	if strings.TrimSpace(password) == "" {
		return nil, errors.New("The password can't be empty")
	}
	if !isUserNameAvailable(username) {
		return nil, errors.New("The username isn't available")
	}

	u := User{username, password}
	userList = append(userList, u)
	return &u, nil
}

func isUserNameAvailable(username string) bool {
	for _, u := range userList {
		if u.UserName == username {
			return false
		}
	}
	return true
}
