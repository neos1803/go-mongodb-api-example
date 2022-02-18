package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	router.GET("/", func(g *gin.Context) {
		g.JSON(200, gin.H{
			"data": "Test Gin API",
		})
	})

	router.Run("localhost:8000")
}
