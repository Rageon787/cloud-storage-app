package main

import (
	"github.com/gin-gonic/gin"
)

func main() {

	// Setup routes
	router := gin.Default()

	// User operations
	router.POST("/user", createUser)            // creates an account for the user
	router.GET("/user:username", getByUsername) // GET a user by username

	// Drive operations
	router.GET("/drive")     // Get all drives of a user
	router.GET("/drive/:id") //

}
