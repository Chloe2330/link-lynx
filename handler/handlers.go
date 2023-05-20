package handler

import (
	"github.com/Chloe2330/link-lynx/shortener"
	"github.com/Chloe2330/link-lynx/store"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// Represents request body for creating a short URL 
type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId string `json:"user_id" binding:"required"`
}
// JSON struct tag defines mapping between struct field and corresponding JSON key
// Binding (gin) struct tag states that both fields must be present in the request
// body when binding JSON data

// Handler function for short URL creation 
func CreateShortUrl(c *gin.Context) {

	// holds parsed request body 
	var creationRequest UrlCreationRequest

	// attempts to bind JSON data from request body to creationRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		// request body is in incorrect format 
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// generates short URL by calling function from shorturl_generator.go
	shortUrl := shortener.GenerateShortLink(creationRequest.LongUrl, creationRequest.UserId)

	// saves URL mapping in cache by calling function from store_service.go 
	store.SaveUrlMapping(shortUrl, creationRequest.LongUrl, creationRequest.UserId)

	// JSON message displayed in browser after successful short URL creation 
	host := "http://localhost:8080/"
	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": host +shortUrl,
	})
}

// Handler function for URL redirection 
func HandleShortUrlRedirect(c *gin.Context) {

	// extracts "shortURL" parameter from route pattern and assigns to new variable
	// for further handling
	shortUrl := c.Param("shortUrl")

	// retrieves value corresponding to shortURL key from cache by calling function 
	// from store_service.go
	initialUrl := store.RetrieveInitialUrl(shortUrl)

	// slices string (format: "userID: initialUrl") by taking substring after first ':'
	initialUrl = initialUrl[strings.Index(initialUrl, ":")+1:]

	// redirects browser to initialURL associated with shortURL
	c.Redirect(302, initialUrl)
}