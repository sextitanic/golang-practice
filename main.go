package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/api/hello-world", responseHelloWorld)

	server.Run(":8080")
}

func responseHelloWorld(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Hello World"})
}
