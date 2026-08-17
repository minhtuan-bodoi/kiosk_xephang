package repositories

import (
	// "os/user"
	"context"
	"kiosk-xephang/database"
	"kiosk-xephang/model"

	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

func CreateUser(user model.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := database.UserCollection.InsertOne(ctx, user)

	return err
}

func GetUser() ([]model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := database.UserCollection.Find(ctx, bson.D{})
	
	if err != nil {
		return nil, err
	}

	users := make([]model.User, 0)

	err = cursor.All(ctx, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}
