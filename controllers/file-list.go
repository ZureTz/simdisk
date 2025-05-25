package controllers

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

// ListFiles handles the request to list files in the static directory
func ListFiles(c *gin.Context) {
	// Get relative path from the request parameters
	relativePath := c.Query("path")

	// Replace all ',' with '/' to ensure proper path formatting
	relativePath = strings.ReplaceAll(relativePath, ",", "/") 

	// Get the list of files in the static directory
	files, err := os.ReadDir("./static/" + relativePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "No such file or directory",
		})
		return
	}

	// List both files and directories
	var fileList []string
	for _, file := range files {
		// If the file is a directory, append a slash to its name
		if file.IsDir() {
			fileList = append(fileList, file.Name()+"/")
			continue
		}
		// If the file is a regular file, just append its name
		fileList = append(fileList, file.Name())
	}

	c.JSON(http.StatusOK, gin.H{
		"files": fileList,
	})
}
