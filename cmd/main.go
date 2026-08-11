package main

import (
	"log"

	"queue-kiosk/routes"
)

func main() {
	router := routes.SetupRouter()
	log.Println("Queue kiosk server starting on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
