package repositories

import (
	"context"
	"errors"
	"time"

	"kiosk-xephang/database"
	"kiosk-xephang/model"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func CreateAppointment(appointment model.Appointmenter) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := database.AppointmentCollection.InsertOne(ctx, appointment)
	return err
}

func GetAppointments() ([]model.Appointmenter, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := database.AppointmentCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	appointments := make([]model.Appointmenter, 0)
	if err := cursor.All(ctx, &appointments); err != nil {
		return nil, err
	}

	return appointments, nil
}

func GetAppointmentByCCCD(cccd string) (model.Appointmenter, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var appointment model.Appointmenter
	err := database.AppointmentCollection.FindOne(ctx, bson.M{"CCCD": cccd}).Decode(&appointment)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return model.Appointmenter{}, errors.New("appointment not found")
		}
		return model.Appointmenter{}, err
	}

	return appointment, nil
}

func UpdateAppointment(id bson.ObjectID, updates bson.M) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := database.AppointmentCollection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": updates},
	)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("appointment not found")
	}

	return nil
}

func DeleteAppointment(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid ID format")
	}

	result, err := database.AppointmentCollection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("appointment not found")
	}

	return nil
}
