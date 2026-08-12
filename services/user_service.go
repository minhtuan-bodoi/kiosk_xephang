package services

import (
	"errors"
	"queue-kiosk/models"
	"queue-kiosk/repositories"
)

func CreateUser(user models.User) error {
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