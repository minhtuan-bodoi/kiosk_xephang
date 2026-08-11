package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ServiceHandler(c *gin.Context) {
	c.String(http.StatusOK, "service handler placeholder")
}
