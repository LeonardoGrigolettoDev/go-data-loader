package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ZIPHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "not implemented",
	})
}
