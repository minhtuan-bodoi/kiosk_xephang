package repositories

import (
	// "os/user"
	"context"
	"queue-kiosk/database"
	"queue-kiosk/models"
	"time"
)

func CreateUser(user models.User) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    _,err := database.UserCollection.InsertOne(ctx,user)

    return err
}