package routes

import (
	"github.com/LeonardoGrigolettoDev/go-data-loader.git/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/zip", handlers.ZIPHandler)
}
