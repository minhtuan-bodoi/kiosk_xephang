package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TicketHandler(c *gin.Context) {
	c.String(http.StatusOK, "ticket handler placeholder")
}
