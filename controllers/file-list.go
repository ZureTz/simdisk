package controllers

import (
	"fmt"
	"net/http"
	"os"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
)

// Example structure for the file list response
// {
//   id: 0,
//   filename: "example.txt",
//   size: 2048, // Size in bytes
//   isFolder: false,
//   // Relative path is used for frontend table row component
//   // to access the file without using useContext()
//   relativePath: "path/to/example.txt"
// },

// Define a structure to hold file details
type FileDetails struct {
	// Unique identifier for the file for frontend use after sorting
	Id int64 `json:"id"`
	// Information about the file
	Filename string `json:"filename"`
	Size     int64  `json:"size"`
	IsFolder bool   `json:"isFolder"`
	// Relative path to the file, used for frontend table row component
	RelativePath string `json:"relativePath"`
}

// ListFiles handles the request to list files in the static directory
func ListFiles(c *gin.Context) {
	// Get relative path from the request parameters
	relativePath := c.Query("path")

	// Replace all ',' with '/' to ensure proper path formatting
	relativePath = strings.ReplaceAll(relativePath, ",", "/")

	// Log the relative path for debugging
	fmt.Println("Relative path:", relativePath)

	// Get the list of files in the static directory
	files, err := os.ReadDir("./static/" + relativePath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "No such file or directory",
		})
		return
	}

	// List both files and directories
	var fileList []FileDetails

	// If the directory is empty, return an empty list
	if len(files) == 0 {
		fileList = append(fileList, FileDetails{
			Id:           0,
			Filename:     "No files found",
			Size:         0,
			IsFolder:     true,
			RelativePath: relativePath,
		})
		// Return an empty file list
		c.JSON(http.StatusOK, gin.H{
			"files": fileList,
		})
		return
	}

	for _, file := range files {
		fileInfo, err := file.Info()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to get file info",
			})
			return
		}

		fileDetails := FileDetails{
			// Id, placeholder for now, will be modified after sorting
			Id:       0,
			Filename: fileInfo.Name(),
			Size:     fileInfo.Size(),
			IsFolder: fileInfo.IsDir(),
			// Relative path to the file (parent directory without file name)
			RelativePath: relativePath,
		}

		fileList = append(fileList, fileDetails)
	}

	// Sort the file list by filename and isFolder
	// Folders should come first, then files, both sorted alphabetically
	sort.Slice(fileList, func(i, j int) bool {
		if fileList[i].IsFolder && !fileList[j].IsFolder {
			return true // Folders come before files
		}

		if !fileList[i].IsFolder && fileList[j].IsFolder {
			return false // Files come after folders
		}

		// If both are folders or both are files, sort by filename
		return fileList[i].Filename < fileList[j].Filename
	})

	// After sorting, assign IDs to each file based on the sorted order
	for i := range fileList {
		// Assign an ID based on the index for sorting purposes
		fileList[i].Id = int64(i)
	}

	c.JSON(http.StatusOK, gin.H{
		"files": fileList,
	})
}
