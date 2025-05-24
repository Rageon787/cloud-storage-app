package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

const storageDir = "./uploads"

func main() {
	// Ensure uploads folder exists
	if err := os.MkdirAll(storageDir, os.ModePerm); err != nil {
		log.Fatalf("Could not create upload directory: %v", err)
	}

	router := gin.Default()

	// Authenticated group
	api := router.Group("/api", AuthMiddleware())
	{
		api.POST("/upload", uploadHandler)
		api.GET("/files/:filename", downloadHandler)
		api.DELETE("/files/:filename", deleteHandler)
	}

	router.Run(":8080")
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		c.Next()
	}
}

// Upload file
func uploadHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File not provided"})
		return
	}

	dst := filepath.Join(storageDir, filepath.Base(file.Filename))
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"file_url":     fmt.Sprintf("http://localhost:8080/api/files/%s", file.Filename),
		"delete_token": "dummy-token", // Placeholder; you can generate a secure token
	})
}

// Download file
func downloadHandler(c *gin.Context) {
	filename := c.Param("filename")
	path := filepath.Join(storageDir, filepath.Base(filename))

	if _, err := os.Stat(path); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	c.FileAttachment(path, filename)
}

// Delete file
func deleteHandler(c *gin.Context) {
	// Validate token if needed
	filename := c.Param("filename")
	path := filepath.Join(storageDir, filepath.Base(filename))

	if err := os.Remove(path); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not delete file"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "File deleted"})
}
