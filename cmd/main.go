package main

import (
	"log"

	"queue-kiosk/configs"
	"queue-kiosk/database"
	"queue-kiosk/routes"
)

func main() {
	cfg := configs.LoadConfig()

	client, err := database.Connect(cfg.Dsn)
	if err != nil {
		log.Fatal("MongoDB connection failed: ", err)
	}
	defer client.Disconnect(nil)

	router := routes.SetupRouter()
	log.Println("Queue kiosk server starting on :" + cfg.Port)
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatal(err)
	}
}
