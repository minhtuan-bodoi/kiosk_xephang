package routes

import (
	"kiosk-xephang/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Serve uploaded files statically
	router.Static("/uploads", "./uploads")

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	api := router.Group("/api")
	{
		services := api.Group("/services")
		{
			services.POST("", handlers.CreateService)
			services.GET("", handlers.GetAllServices)
			services.GET("/code/:code", handlers.GetServicesByCodeService)
			services.GET("/:id", handlers.GetServiceByID)
			services.PUT("/:id", handlers.UpdateService)
			services.DELETE("/:id", handlers.DeleteService)
		}

		tickets := api.Group("/tickets")
		{
			tickets.POST("", handlers.CreateTicket)
			tickets.GET("", handlers.GetAllTickets)
			tickets.GET("/code/:code", handlers.GetTicketByTicketCode)
			tickets.GET("/:id", handlers.GetTicketByID)
			tickets.PUT("/:id", handlers.UpdateTicket)
			tickets.DELETE("/:id", handlers.DeleteTicket)
		}

		appointments := api.Group("/appointments")
		{
			appointments.POST("", handlers.CreateAppointment)
			appointments.GET("", handlers.GetAllAppointments)
			appointments.GET("/cccd/:cccd", handlers.GetAppointmentByCCCD)
			appointments.PUT("/:id", handlers.UpdateAppointment)
			appointments.DELETE("/:id", handlers.DeleteAppointment)
		provinces := api.Group("/provinces")
		{
			provinces.POST("", handlers.CreateProvince)
			provinces.GET("", handlers.GetAllProvinces)
			provinces.GET("/code/:code", handlers.GetProvincesByCodeProvince)
			provinces.GET("/:id", handlers.GetProvinceByID)
			provinces.PUT("/:id", handlers.UpdateProvince)
			provinces.DELETE("/:id", handlers.DeleteProvince)
		}

		wards := api.Group("/wards")
		{
			wards.POST("", handlers.CreateWard)
			wards.GET("", handlers.GetAllWards)
			wards.GET("/code/:code", handlers.GetWardsByCodeWard)
			wards.GET("/:id", handlers.GetWardByID)
			wards.PUT("/:id", handlers.UpdateWard)
			wards.DELETE("/:id", handlers.DeleteWard)
		}
	}

	return router
}

}