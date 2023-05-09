package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {

	// Creates a Gin router with default middleware
	router := gin.Default()

	// When a HTTP GET request is made to the root ("/") URL path, function is executed
	router.GET("/", func(c *gin.Context) {

		// Sends an HTTP response back to client using JSON method of gin.Context object
		c.JSON(200, gin.H{
			"message": "Hey Go URL Shortener!",
		}) // two arguments: HTTP status code, JSON data

		// gin.H{} creates a map with string key ("message") and corresponding values 
	})

	// Starts HTTP server and listens on port 8080
	err := router.Run(":8080")

	// Stops execution of goroutine and return error message
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}