package routes

import (
	"queue-kiosk/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	router.Any("/auth/login", handlers.LoginHandler)
	router.Any("/tickets", handlers.TicketHandler)
	router.Any("/queues", handlers.QueueHandler)
	router.Any("/counters", handlers.CounterHandler)
	router.Any("/services", handlers.ServiceHandler)

	return router
}
