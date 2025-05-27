package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/ZureTz/simdisk/utils"
	"github.com/gin-gonic/gin"
)

// CreateFolder handles the request to create a new folder

func CreateFolder(c *gin.Context) {
	// Get the relative path from the request parameters
	relativePath := c.Query("path")

	// Replace all ',' with '/' to ensure proper path formatting
	relativePath = strings.ReplaceAll(relativePath, ",", "/")

	// Log the relative path for debugging
	fmt.Println("Relative path:", relativePath)

	// Get the folder name from the request body formData
	folderName := c.PostForm("folderName")
	if folderName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "folderName parameter is required",
		})
		return
	}
	// Ensure the folder name does not contain any invalid characters
	invalidChars := []string{"/", "\\", ":", "*", "?", "\"", "<", ">", "|"}
	for _, char := range invalidChars {
		if strings.Contains(folderName, char) {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("Invalid folder name: '%s' contains invalid characters", folderName),
			})
			return
		}
	}

	// Get the working directory from the config
	workingDir := utils.Config.WorkingDirectory.Path

	// Create the new folder in the working directory
	err := os.Mkdir(workingDir+relativePath+"/"+folderName, 0755)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create folder",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Folder created successfully",
	})
}
