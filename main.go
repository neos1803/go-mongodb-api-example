package main

import (
	"go-mongodb-api-example/configs"
	"go-mongodb-api-example/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	configs.DBConfig()

	router.GET("/", func(g *gin.Context) {
		g.JSON(200, gin.H{
			"data": "Test Gin API",
		})
	})

	routes.UserRoute(router)

	router.Run(os.Getenv("SERVER_ADDRESS"))
}
