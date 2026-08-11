package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func QueueHandler(c *gin.Context) {
	c.String(http.StatusOK, "queue handler placeholder")
}
