package handlers

import (
	"net/http"

	"kiosk-xephang/dto"
	"kiosk-xephang/model"
	"kiosk-xephang/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func CreateProvince(c *gin.Context) {
	var province model.Province
	if err := c.ShouldBindJSON(&province); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreateProvince(province); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Tạo province thành công",
		"data":    province,
	})
}

func GetAllProvinces(c *gin.Context) {
	res, err := services.GetAllProvinces()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}

func GetProvinceByID(c *gin.Context) {
	id := c.Param("id")
	res, err := services.GetProvinceByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}

func GetProvincesByCodeProvince(c *gin.Context) {
	code := c.Param("code")
	res, err := services.GetProvincesByCodeProvince(code)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": res})
}

func UpdateProvince(c *gin.Context) {
	idStr := c.Param("id")
	objectID, err := bson.ObjectIDFromHex(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid province ID format"})
		return
	}

	var req dto.UpdateProvinceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.UpdateProvince(objectID, req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Province updated successfully"})
}

func DeleteProvince(c *gin.Context) {
	id := c.Param("id")
	if err := services.DeleteProvince(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Province deleted successfully"})
}
