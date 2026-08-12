package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"

	"appointment-kiosk/configs"
	"appointment-kiosk/database"
)

func main() {

	config := configs.LoadConfig()

	mongoClient := database.ConnectMongoDB(config)

	defer func() {
		if err := mongoClient.Disconnect(context.Background()); err != nil {
			log.Println("MongoDB disconnect error:", err)
		}
	}()

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "OK",
			"message": "Appointment Server is running",
		})
	})

	log.Printf(
		"Appointment Server running on :%s",
		config.ServerPort,
	)

	if err := router.Run(":" + config.ServerPort); err != nil {
		log.Fatal(err)
	}
}
