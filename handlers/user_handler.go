package handlers

import (
	"net/http"
	"kiosk-xephang/model"
	"kiosk-xephang/services"

	"github.com/gin-gonic/gin"
)

func CreateUser(ctx *gin.Context) {

var  user model.User

	if err := ctx.ShouldBindJSON(&user); err!=nil {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}

	err := services.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest,gin.H{
			"error":err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Tạo user thành công",
		"user":    user,
	})
}


func GetUser (ctx *gin.Context){
	users, err := services.GetUser()

	if  err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error":err.Error(),})
		return 
	}

	ctx.JSON(http.StatusOK, gin.H{"users":users,})
}