package main

import (
	"github.com/gin-gonic/gin"
)

func getHelloWorld(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, World!",
	})
}

func main() {
	r := gin.Default()
	r.GET("/hello", getHelloWorld)
	r.Run(":8081") // Run on port 8081
}
