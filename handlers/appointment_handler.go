package handlers

import (
	"net/http"

	"kiosk-xephang/dto"
	"kiosk-xephang/model"
	"kiosk-xephang/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// CreateAppointment godoc
// @Summary Create a new appointment directly using Appointmenter model
// @Tags Appointments
// @Accept json,multipart/form-data
// @Produce json
// @Param appointment body model.Appointmenter true "Appointment Info"
// @Success 201 {object} model.Appointmenter
// @Router /appointments [post]
func CreateAppointment(c *gin.Context) {
	var appointment model.Appointmenter
	if err := c.ShouldBind(&appointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreateAppointment(appointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Appointment created successfully",
		"data":    appointment,
	})
}

// GetAllAppointments godoc
// @Summary Get list of all appointments
// @Tags Appointments
// @Produce json
// @Success 200 {array} model.Appointmenter
// @Router /appointments [get]
func GetAllAppointments(c *gin.Context) {
	res, err := services.GetAllAppointments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

// GetAppointmentByCCCD godoc
// @Summary Get an appointment by CCCD
// @Tags Appointments
// @Produce json
// @Param cccd path string true "CCCD"
// @Success 200 {object} model.Appointmenter
// @Router /appointments/cccd/{cccd} [get]
func GetAppointmentByCCCD(c *gin.Context) {
	cccd := c.Param("cccd")

	res, err := services.GetAppointmentByCCCD(cccd)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

// UpdateAppointment godoc
// @Summary Update appointment details using UpdateAppointmentRequest
// @Tags Appointments
// @Accept json,multipart/form-data
// @Produce json
// @Param id path string true "Appointment ID"
// @Param appointment body dto.UpdateAppointmentRequest true "Update fields"
// @Success 200 {object} map[string]string
// @Router /appointments/{id} [put]
func UpdateAppointment(c *gin.Context) {
	idStr := c.Param("id")
	objectID, err := bson.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid appointment ID format"})
		return
	}

	var req dto.UpdateAppointmentRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.UpdateAppointment(objectID, req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Appointment updated successfully",
	})
}

// DeleteAppointment godoc
// @Summary Delete an appointment
// @Tags Appointments
// @Param id path string true "Appointment ID"
// @Success 200 {object} map[string]string
// @Router /appointments/{id} [delete]
func DeleteAppointment(c *gin.Context) {
	id := c.Param("id")

	if err := services.DeleteAppointment(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Appointment deleted successfully",
	})
}
