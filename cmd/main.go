package main

import (
	"context"
	"log"

	"kiosk-xephang/configs"
	"kiosk-xephang/database"
	"kiosk-xephang/routes"
)

func main() {
	config := configs.LoadConfig()

	mongoClient := database.ConnectMongoDB(config)
	defer func() {
		if err := mongoClient.Disconnect(context.Background()); err != nil {
			log.Println("MongoDB disconnect error:", err)
		}
	}()

	// Initialize Gin router
	router := routes.SetupRouter()

	log.Printf("Appointment Server running on :%s", config.ServerPort)

	if err := router.Run(":" + config.ServerPort); err != nil {
		log.Fatal(err)
	}
}
