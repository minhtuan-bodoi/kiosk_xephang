package database

import (
	"context"
	"log"
	"time"

	"kiosk-xephang/configs"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var DB *mongo.Database
var ServiceCollection *mongo.Collection
var TicketCollection *mongo.Collection
var UserCollection *mongo.Collection
var AppointmentCollection *mongo.Collection
var ProvinceCollection *mongo.Collection
var WardCollection *mongo.Collection

func ConnectMongoDB(config *configs.Config) *mongo.Client {

	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	client, err := mongo.Connect(
		options.Client().ApplyURI(config.MongoURI),
	)

	if err != nil {
		log.Fatal("MongoDB connection error:", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("MongoDB ping error:", err)
	}

	DB = client.Database(config.MongoDatabase)
	ServiceCollection = DB.Collection("services")
	TicketCollection = DB.Collection("tickets")
	UserCollection = DB.Collection("users")
	AppointmentCollection = DB.Collection("appointment")
	ProvinceCollection = DB.Collection("provinces")
	WardCollection = DB.Collection("wards")

	log.Println("MongoDB connected successfully")

	return client
}
