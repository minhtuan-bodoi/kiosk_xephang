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

func CreateWard(ward model.Ward) error {
	ward.CreatedAt = time.Now()
	ward.UpdatedAt = time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := database.WardCollection.InsertOne(ctx, ward)
	return err
}

func GetWards() ([]model.Ward, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := database.WardCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	wards := make([]model.Ward, 0)
	if err := cursor.All(ctx, &wards); err != nil {
		return nil, err
	}

	return wards, nil
}

func GetWardByID(id string) (model.Ward, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return GetWardsByCodeWard(id)
	}

	var ward model.Ward
	err = database.WardCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&ward)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return GetWardsByCodeWard(id)
		}
		return model.Ward{}, err
	}

	return ward, nil
}

func GetWardsByCodeWard(code string) (model.Ward, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var ward model.Ward
	err := database.WardCollection.FindOne(ctx, bson.M{"code": code}).Decode(&ward)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return model.Ward{}, errors.New("ward not found")
		}
		return model.Ward{}, err
	}

	return ward, nil
}

func UpdateWard(id bson.ObjectID, updates bson.M) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	updates["updated_at"] = time.Now()
	result, err := database.WardCollection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": updates},
	)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("ward not found")
	}

	return nil
}

func DeleteWard(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid ID format")
	}

	result, err := database.WardCollection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("ward not found")
	}

	return nil
}

func IsWardCodeExists(code string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	count, err := database.WardCollection.CountDocuments(ctx, bson.M{"code": code})
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
