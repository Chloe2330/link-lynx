package store

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testStoreService = &StorageService{}

// special function automatically executed before main()
func init() {
	testStoreService = InitializeStore()
}

// unit test ensures initialization of store service 
func TestStoreInit(t *testing.T) {
	assert.True(t, testStoreService.redisClient != nil)
}

// unit test for verifying the insertion and retrieval of data from the store
func TestInsertionAndRetrieval(t *testing.T) {
	initialLink := "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html"
	userUUId := "e0dba740-fc4b-4977-872c-d360239e6b1a"
	shortURL := "Jsz4k57oAX"

	// maps short URL to initial URL and user ID
	SaveUrlMapping(shortURL, initialLink, userUUId)

	// retrieve initial URL
	retrievedUrl := RetrieveInitialUrl(shortURL)

	// ensures that initial URL and retrieved URL are the same
	assert.Equal(t, initialLink, retrievedUrl)
}