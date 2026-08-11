package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CounterHandler(c *gin.Context) {
	c.String(http.StatusOK, "counter handler placeholder")
}
