package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// DeleteFile handles the request to delete a file
func DeleteFile(c *gin.Context) {
	// Get the relative path from the request parameters
	relativePath := c.Query("path")

	// Replace all ',' with '/' to ensure proper path formatting
	relativePath = strings.ReplaceAll(relativePath, ",", "/")

	// Log the relative path for debugging
	fmt.Println("Relative path:", relativePath)

	// Get the file name from the request body formData
	fileName := c.PostForm("fileName")
	if fileName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "fileName parameter is required",
		})
		return
	}

	// Construct the full file path
	filePath := "./static/" + relativePath + "/" + fileName
	fmt.Println("Relative path file:", filePath)

	// Remove the file
	err := os.Remove(filePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to delete file",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "File deleted successfully",
	})

}
