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

func CreateProvince(province model.Province) error {
	province.CreatedAt = time.Now()
	province.UpdatedAt = time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := database.ProvinceCollection.InsertOne(ctx, province)
	return err
}

func GetProvinces() ([]model.Province, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := database.ProvinceCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	provinces := make([]model.Province, 0)
	if err := cursor.All(ctx, &provinces); err != nil {
		return nil, err
	}

	return provinces, nil
}

func GetProvinceByID(id string) (model.Province, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return GetProvincesByCodeProvince(id)
	}

	var province model.Province
	err = database.ProvinceCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&province)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return GetProvincesByCodeProvince(id)
		}
		return model.Province{}, err
	}

	return province, nil
}

func GetProvincesByCodeProvince(code string) (model.Province, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var province model.Province
	err := database.ProvinceCollection.FindOne(ctx, bson.M{"code": code}).Decode(&province)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return model.Province{}, errors.New("province not found")
		}
		return model.Province{}, err
	}

	return province, nil
}

func UpdateProvince(id bson.ObjectID, updates bson.M) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	updates["updated_at"] = time.Now()
	result, err := database.ProvinceCollection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": updates},
	)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("province not found")
	}

	return nil
}

func DeleteProvince(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid ID format")
	}

	result, err := database.ProvinceCollection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("province not found")
	}

	return nil
}

func IsProvinceCodeExists(code string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	count, err := database.ProvinceCollection.CountDocuments(ctx, bson.M{"code": code})
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func GetProvince() ([]model.Province, error) {
	return GetProvinces()
}
