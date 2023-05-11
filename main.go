package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	// creates a Gin router with default middleware
	router := gin.Default()

	// when a HTTP GET request is made to the root ("/") URL path, function is executed
	router.GET("/", func(c *gin.Context) {

		// sends an HTTP response back to client using JSON method of gin.Context object
		c.JSON(200, gin.H{
			"message": "Welcome to ShortURL!",
		}) // two arguments: HTTP status code, JSON data

		// gin.H{} creates a map with string key ("message") and a corresponding value 
	})

	// starts HTTP server and listens on port 8080
	err := router.Run(":8080")

	// stops execution of goroutine and return error message
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}