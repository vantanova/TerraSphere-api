package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
	router.GET("/", getHelloWorld)

    router.Run("localhost:8080")
}

func getHelloWorld(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
		"message": "Hello World, this is a message",
	})
}
