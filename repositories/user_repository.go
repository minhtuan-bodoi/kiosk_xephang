package repositories

import (
	// "os/user"
	"context"
	"queue-kiosk/database"
	"queue-kiosk/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func CreateUser(user models.User) error {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    _,err := database.UserCollection.InsertOne(ctx,user)

    return err
}

func GetUser()([]models.User, error) {
	ctx , cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := database.UserCollection.Find(ctx,bson.D{})
	if  err != nil {
		return nil, err
	}

	users := make([]models.User, 0)

	err = cursor.All(ctx, &users)
	if  err != nil {
		return nil, err
	}

	return users, nil
}