package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a new Gin router
	router := gin.Default()

	// Define the route and handler for the GET request
	router.GET("/", func(c *gin.Context) {
		// Set the response content type as JSON
		c.Header("Content-Type", "application/json")

		// Generate a random number
		rand.Seed(time.Now().UnixNano())
		randomNumber := rand.Intn(100)

		// Return the random number in the response
		c.JSON(http.StatusOK, gin.H{
			"random_number": randomNumber,
		})
	})

	// Run the HTTP server
	router.Run("0.0.0.0:8080")
}

