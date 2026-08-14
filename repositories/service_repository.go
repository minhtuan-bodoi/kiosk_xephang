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

func CreateService(service model.Service) error {
	service.CreatedAt = time.Now()
	service.UpdatedAt = time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := database.ServiceCollection.InsertOne(ctx, service)
	return err
}

func GetServices() ([]model.Service, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := database.ServiceCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	services := make([]model.Service, 0)
	if err := cursor.All(ctx, &services); err != nil {
		return nil, err
	}

	return services, nil
}

func GetServiceByID(id string) (model.Service, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return GetServicesByCodeService(id)
	}

	var service model.Service
	err = database.ServiceCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&service)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return GetServicesByCodeService(id)
		}
		return model.Service{}, err
	}

	return service, nil
}

func GetServicesByCodeService(codeService string) (model.Service, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var service model.Service
	err := database.ServiceCollection.FindOne(ctx, bson.M{"code_service": codeService}).Decode(&service)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return model.Service{}, errors.New("service not found")
		}
		return model.Service{}, err
	}

	return service, nil
}

func UpdateService(id bson.ObjectID, updates bson.M) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	updates["updated_at"] = time.Now()

	result, err := database.ServiceCollection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": updates},
	)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("service not found")
	}

	return nil
}

func DeleteService(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid ID format")
	}

	result, err := database.ServiceCollection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("service not found")
	}

	return nil
}

func IsServiceCodeExists(code string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	count, err := database.ServiceCollection.CountDocuments(ctx, bson.M{
		"code_service": code,
	})
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func IsTicketFormatCodeExists(code string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	count, err := database.ServiceCollection.CountDocuments(ctx, bson.M{
		"ticket_format_code": code,
	})
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
