package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/ZureTz/simdisk/utils"
	"github.com/gin-gonic/gin"
)

// UploadFiles handles the file upload request
func UploadFile(c *gin.Context) {
	// Get count of files uploaded from the params
	count := c.Query("fileCount")
	if count == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "fileCount parameter is required",
		})
		return
	}
	fmt.Println("File count:", count)

	// Convert count to an integer
	var err error
	var fileCount int
	if fileCount, err = strconv.Atoi(count); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid fileCount parameter",
		})
		return
	}

	// Get relative path from the request parameters
	relativePath := c.Query("path")

	// Replace all ',' with '/' to ensure proper path formatting
	relativePath = strings.ReplaceAll(relativePath, ",", "/")
	// Log the relative path for debugging
	fmt.Println("Relative path:", relativePath)

	// Get working directory from the config
	workingDir := utils.Config.WorkingDirectory.Path

	// Save file in a for loop
	for i := range fileCount {
		// Generate the file name based on the index
		fileName := "file" + strconv.Itoa(i)

		file, err := c.FormFile(fileName)
		if err != nil {
			// Log the error if needed
			fmt.Println("Error uploading file:", err.Error())

			c.JSON(http.StatusBadRequest, gin.H{
				"error": "File upload failed",
			})
			return
		}

		// Save the file to the server
		if err := c.SaveUploadedFile(file, workingDir+relativePath+"/"+file.Filename); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "File save failed",
			})
			return
		}

	}

	c.JSON(http.StatusOK, gin.H{
		"message": "File uploaded successfully",
	})
}
