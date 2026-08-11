package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	c.String(http.StatusOK, "auth login placeholder")
}
