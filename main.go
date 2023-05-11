package main

import (
	"fmt"
	"github.com/Chloe2330/go-url-shortener/handler"
	"github.com/Chloe2330/go-url-shortener/store"
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

	router.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	router.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	// store initialization
	store.InitializeStore()

	// starts HTTP server and listens on port 8080
	err := router.Run(":8080")

	// stops execution of goroutine and return error message
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}

// curl command (WSL Ubuntu)
/*
curl -X POST -H "Content-Type: application/json" -d '{
    "long_url": "https://www.digitalocean.com/community/tutorials/how-to-build-a-ruby-on-rails-application#step-5-adding-validations",
    "user_id": "e0dba740-fc4b-4977-872c-d360239e6b10"
}' http://localhost:8080/create-short-url
*/

// useful Redis commands 
/* 
sudo service redis-server stop 
sudo service redis-server start 
KEYS *
GET <key>
DEL <key>
TIL <key> 
EXPIRE <key> seconds
*/
