package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// Set up a simple GET endpoint
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Handle static files
	r.Static("/static", "./static")

	// Handle file uploads
	r.POST("/api/upload", func(c *gin.Context) {
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
			if err := c.SaveUploadedFile(file, "./static/"+file.Filename); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "File save failed",
				})
				return
			}

		}

		c.JSON(http.StatusOK, gin.H{
			"message": "File uploaded successfully",
		})

	})
	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
