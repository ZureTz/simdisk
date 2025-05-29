package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/ZureTz/simdisk/utils"
	"github.com/gin-gonic/gin"
)

func DownloadFile(c *gin.Context) {
	// Get the relative path from the request parameters
	relativePath := c.Query("path")

	// Replace all ',' with '/' to ensure proper path formatting
	relativePath = strings.ReplaceAll(relativePath, ",", "/")

	// Get fileName from the params
	fileName := c.Query("fileName")
	if fileName == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "File name is required",
		})
		return
	}

	// Log the relative path for debugging
	fmt.Println("Relative path:", relativePath)

	workingDir := utils.Config.WorkingDirectory.Path
	// Construct the full file path
	filePathFull := workingDir + relativePath + "/" + fileName
	fmt.Println("File path:", filePathFull)

	// Check if the file exists
	var stat os.FileInfo
	var err error
	if stat, err = os.Stat(filePathFull); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "File not found",
		})
		return
	}

	// Set the response headers to indicate file download
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", fileName))
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Length", fmt.Sprintf("%d", stat.Size()))
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Expires", "0")
	c.Header("Cache-Control", "must-revalidate")
	c.Header("Pragma", "public")

	// Send the file to the client
	c.File(filePathFull)
}
