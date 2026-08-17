package repositories

import (
	"context"
	"errors"
	"time"

	"kiosk-xephang/database"
	"kiosk-xephang/model"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func CreateTicket(ticket model.Ticket) error {
	ticket.CreatedAt = time.Now()
	ticket.UpdatedAt = time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := database.TicketCollection.InsertOne(ctx, ticket)
	return err
}

func GetTickets() ([]model.Ticket, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})
	cursor, err := database.TicketCollection.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	tickets := make([]model.Ticket, 0)
	if err := cursor.All(ctx, &tickets); err != nil {
		return nil, err
	}

	return tickets, nil
}

func GetTicketByID(id string) (model.Ticket, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return GetTicketByTicketCode(id)
	}

	var ticket model.Ticket
	err = database.TicketCollection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&ticket)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return GetTicketByTicketCode(id)
		}
		return model.Ticket{}, err
	}

	return ticket, nil
}

func GetTicketByTicketCode(ticketCode string) (model.Ticket, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var ticket model.Ticket
	err := database.TicketCollection.FindOne(ctx, bson.M{"ticket_code": ticketCode}).Decode(&ticket)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return model.Ticket{}, errors.New("ticket not found")
		}
		return model.Ticket{}, err
	}

	return ticket, nil
}

func UpdateTicket(id bson.ObjectID, updates bson.M) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	updates["updated_at"] = time.Now()
	result, err := database.TicketCollection.UpdateOne(
		ctx,
		bson.M{"_id": id},
		bson.M{"$set": updates},
	)
	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("ticket not found")
	}

	return nil
}

func DeleteTicket(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objectID, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid ID format")
	}

	result, err := database.TicketCollection.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("ticket not found")
	}

	return nil
}
