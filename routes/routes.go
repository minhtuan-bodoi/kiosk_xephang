package routes

import (
	"net/http"

	"queue-kiosk/handlers"
)

func SetupRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/auth/login", handlers.LoginHandler)
	mux.HandleFunc("/tickets", handlers.TicketHandler)
	mux.HandleFunc("/queues", handlers.QueueHandler)
	mux.HandleFunc("/counters", handlers.CounterHandler)
	mux.HandleFunc("/services", handlers.ServiceHandler)

	return mux
}
