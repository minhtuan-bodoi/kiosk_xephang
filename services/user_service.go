package services

import (
	"errors"
	"kiosk-xephang/model"
	"kiosk-xephang/repositories"
)

func CreateUser(user model.User) error {
	if user.Username == "" {
		return errors.New("Tên user không được để trống")
	}

	if user.Password == "" {
		return errors.New("Password không được để trống")
	}

	if user.Email == "" {
		return errors.New("Email không được để trống")
	}

	return repositories.CreateUser(user)
}

func GetUser()([]model.User, error){
	return repositories.GetUser()
}