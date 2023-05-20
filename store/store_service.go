package store

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

// Defines the struct wrapper around raw Redis client
type StorageService struct {
	redisClient *redis.Client
}

// Creates a new instance of the struct and assigns its address to storeService 
var (
	storeService = &StorageService{}
)

// Note: the cache duration shouldn't have an expiration time, an LRU 
// policy config should be set where values that are retrieved less 
// often are purged automatically from the cache and stored back in 
// RDBMS (PostgreSQL, MySQL etc.) whenever the cache is full (future)

// Cached data will be stored in Redis for six hours
const CacheDuration = 6 * time.Hour

// Initializes the store service and returns a pointer to StorageService object 
func InitializeStore() *StorageService {

	// Initializes a new Redis client with configuration object
	redisClient := redis.NewClient(&redis.Options{
		// default address of Redis server*
		Addr: "redis:6379",
		// empty password
		Password: "",
		// default Redis database (DB 0)
		DB: 0,
	}) 
	// *SUPER IMPORTANT COMMENT: port 6379 of the 'redis' container is binded to
	// to the computer, so "localhost:6379" allows only the computer to connect
	// to the server... each docker compose container can access other containers
	// by using the service name as the hostname, so the 'main' container can 
	// access the redis service with the hostname 'redis' (or 'db' or whatever), 
	// so "redis:6379" should be used for connection in this situation

	// sends 'PING' command to Redis server, checks for response (pong) and error
	pong, err := redisClient.Ping().Result()

	// Redis server did not start
	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}

	// Redis server successfully started
	fmt.Printf("\nRedis started successfully: pong message = {%s}", pong)

	// assigns the initialized Redis client to the storeService object
	storeService.redisClient = redisClient

	// returns pointer to an instance of the struct
	return storeService
}

// Saves URL mapping in the cache 
func SaveUrlMapping(shortUrl string, originalUrl string, userId string) {

	// sets key-value pair (shortUrl - userId: originalUrl) and key expiration time
	err := storeService.redisClient.Set(shortUrl, userId+": "+originalUrl, CacheDuration).Err()

	// Note: redisClient.Set() takes three arguments, string key, interface{}
	// value (strings, numbers, JSON data, etc.), and time.Duration

	// interface{} in Go can represent the value of any type and is the 
	// equivalent to the object type in Python and the Object class in Java
	
	// interface{} allows for the storage of values of various types in Redis, 
	// and the library automatically serializes and deserializes values when 
	// interacting with the Redis server

	// could not save key
	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl))
	}
	
}

// retrieves original URL from the cache 
func RetrieveInitialUrl(shortUrl string) string {

	// retrieves value associated with the shortUrl key from the Redis server
	result, err := storeService.redisClient.Get(shortUrl).Result()

	// could not retrieve value 
	if err != nil {
		panic(fmt.Sprintf("Failed RetrieveInitialUrl url | Error: %v - shortUrl: %s\n", err, shortUrl))
	}
	
	// returns value if retrieval is successful
	return result
}