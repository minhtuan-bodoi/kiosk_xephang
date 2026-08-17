package handlers

import (
	"net/http"

	"kiosk-xephang/dto"
	"kiosk-xephang/model"
	"kiosk-xephang/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// CreateTicket godoc
// @Summary Create a new ticket directly using Ticket model
// @Tags Tickets
// @Accept json,multipart/form-data
// @Produce json
// @Param ticket_code formData string false "Ticket Code"
// @Param ticket_format_code formData string false "Ticket Format Code"
// @Success 201 {object} model.Ticket
// @Router /tickets [post]
func CreateTicket(c *gin.Context) {
	var ticket model.Ticket
	if err := c.ShouldBind(&ticket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreateTicket(ticket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Ticket created successfully",
		"data":    ticket,
	})
}

// GetAllTickets godoc
// @Summary Get list of all tickets
// @Tags Tickets
// @Produce json
// @Success 200 {array} model.Ticket
// @Router /tickets [get]
func GetAllTickets(c *gin.Context) {
	res, err := services.GetAllTickets()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

// GetTicketByID godoc
// @Summary Get a ticket by ID
// @Tags Tickets
// @Produce json
// @Param id path string true "Ticket ID"
// @Success 200 {object} model.Ticket
// @Router /tickets/{id} [get]
func GetTicketByID(c *gin.Context) {
	id := c.Param("id")

	res, err := services.GetTicketByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

// GetTicketByTicketCode godoc
// @Summary Get a ticket by TicketCode
// @Tags Tickets
// @Produce json
// @Param code path string true "Ticket Code"
// @Success 200 {object} model.Ticket
// @Router /tickets/code/{code} [get]
func GetTicketByTicketCode(c *gin.Context) {
	code := c.Param("code")

	res, err := services.GetTicketByTicketCode(code)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

// UpdateTicket godoc
// @Summary Update ticket details using UpdateTicketDTO
// @Tags Tickets
// @Accept json,multipart/form-data
// @Produce json
// @Param id path string true "Ticket ID"
// @Param ticket_code formData string false "Ticket Code"
// @Param ticket_format_code formData string false "Ticket Format Code"
// @Success 200 {object} model.Ticket
// @Router /tickets/{id} [put]
func UpdateTicket(c *gin.Context) {
	idStr := c.Param("id")
	objectID, err := bson.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ticket ID format"})
		return
	}

	var req dto.UpdateTicketRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.UpdateTicket(objectID, req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Ticket updated successfully",
	})
}

// DeleteTicket godoc
// @Summary Delete a ticket
// @Tags Tickets
// @Param id path string true "Ticket ID"
// @Success 200 {object} map[string]string
// @Router /tickets/{id} [delete]
func DeleteTicket(c *gin.Context) {
	id := c.Param("id")

	if err := services.DeleteTicket(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Ticket deleted successfully",
	})
}
