package handlers

import (
	"net/http"

	"kiosk-xephang/dto"
	"kiosk-xephang/model"
	"kiosk-xephang/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func CreateWard(c *gin.Context) {
	var ward model.Ward
	if err := c.ShouldBindJSON(&ward); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreateWard(ward); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Tạo ward thành công",
		"data":    ward,
	})
}

func GetAllWards(c *gin.Context) {
	res, err := services.GetAllWards()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}

func GetWardByID(c *gin.Context) {
	id := c.Param("id")
	res, err := services.GetWardByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}

func GetWardsByCodeWard(c *gin.Context) {
	code := c.Param("code")
	res, err := services.GetWardsByCodeWard(code)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}

func UpdateWard(c *gin.Context) {
	idStr := c.Param("id")
	objectID, err := bson.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ward ID format"})
		return
	}

	var req dto.UpdateWardRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.UpdateWard(objectID, req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ward updated successfully"})
}

func DeleteWard(c *gin.Context) {
	id := c.Param("id")
	if err := services.DeleteWard(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Ward deleted successfully"})
}
