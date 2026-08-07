package main

import (
	"log"
	"net/http"

	"queue-kiosk/routes"
)

func main() {
	router := routes.SetupRouter()
	log.Println("Queue kiosk server starting on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
