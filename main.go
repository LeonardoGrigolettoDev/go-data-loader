package main

import (
	"github.com/LeonardoGrigolettoDev/go-data-loader.git/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Configurar rotas
	routes.SetupRoutes(router)

	// Iniciar o servidor
	router.Run(":8080")
}
