package main

import (
	"fmt"
	"github.com/Chloe2330/link-lynx/handler"
	"github.com/Chloe2330/link-lynx/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	// creates a Gin router with default middleware
	router := gin.Default()

	router.LoadHTMLGlob("templates/index.html")

	// GET request is made to the root ("/") path
	router.Static("/static", "./static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"cssFile": "/static/styles.css",
			"logoPath": "/static/lynx-logo.png",
			"jsFile":    "/static/placeholder.js",
		})
	})

	/*
	router.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})
	*/

	// POST request is made to the "/create-short-url" path, handles short URL creation
	router.POST("/create-short-url", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	// GET request is made to the "/:shortUrl" path, handles redirection to initial URL
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
