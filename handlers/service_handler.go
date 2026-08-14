package handlers

import (
	"net/http"
	"path/filepath"

	"kiosk-xephang/dto"
	"kiosk-xephang/model"
	"kiosk-xephang/services"
	"kiosk-xephang/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

// CreateService godoc
// @Summary Create a new service directly using Service model
// @Tags Services
// @Accept multipart/form-data
// @Produce json
// @Param name_service formData string true "Service Name"
// @Param code_service formData string false "Service Code"
// @Param ticket_format_code formData string false "Ticket Format Code"
// @Param icon formData file false "Service Icon Image"
// @Success 201 {object} model.Service
// @Router /services [post]
func CreateService(c *gin.Context) {
	var service model.Service
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file, err := c.FormFile("icon")
	if err != nil {
		file, _ = c.FormFile("icon_service")
	}

	if file != nil {
		if !utils.IsAllowedImage(file.Filename) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid file type, allowed: jpg, jpeg, png, gif, svg, webp"})
			return
		}

		filename, err := utils.SaveUploadedFile(file, services.UploadDir)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error saving icon file: " + err.Error()})
			return
		}
		service.IconService = filename
	}

	if err := services.CreateService(service); err != nil {
		if service.IconService != "" {
			_ = utils.DeleteFile(filepath.Join(services.UploadDir, service.IconService))
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Service created successfully",
		"data":    service,
	})
}

// GetAllServices godoc
// @Summary Get list of all services
// @Tags Services
// @Produce json
// @Success 200 {array} model.Service
// @Router /services [get]
func GetAllServices(c *gin.Context) {
	res, err := services.GetAllServices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

// GetServiceByID godoc
// @Summary Get a service by ID
// @Tags Services
// @Produce json
// @Param id path string true "Service ID"
// @Success 200 {object} model.Service
// @Router /services/{id} [get]
func GetServiceByID(c *gin.Context) {
	id := c.Param("id")

	res, err := services.GetServiceByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}

// UpdateService godoc
// @Summary Update service details and/or icon image using UpdateServiceRequest
// @Tags Services
// @Accept multipart/form-data
// @Produce json
// @Param id path string true "Service ID"
// @Param name_service formData string false "Service Name"
// @Param code_service formData string false "Service Code"
// @Param ticket_format_code formData string false "Ticket Format Code"
// @Param icon formData file false "Service Icon Image"
// @Success 200 {object} model.Service
// @Router /services/{id} [put]
func UpdateService(c *gin.Context) {
	idStr := c.Param("id")
	objectID, err := bson.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid service ID format"})
		return
	}

	var req dto.UpdateServiceRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.IconService == nil {
		if file, err := c.FormFile("icon_service"); err == nil {
			req.IconService = file
		} else if file, err := c.FormFile("icon"); err == nil {
			req.IconService = file
		}
	}

	if err := services.UpdateService(objectID, req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Service updated successfully",
	})
}

// DeleteService godoc
// @Summary Delete service and remove icon image file
// @Tags Services
// @Param id path string true "Service ID"
// @Success 200 {object} map[string]string
// @Router /services/{id} [delete]
func DeleteService(c *gin.Context) {
	id := c.Param("id")

	if err := services.DeleteService(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Service and icon deleted successfully",
	})
}
