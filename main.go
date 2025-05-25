package main

import (
	"github.com/ZureTz/simdisk/controllers"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	r := gin.Default()

	// Allow CORS for all origins
	r.Use(cors.Default())

	// Handle file uploads
	r.POST("/api/upload", controllers.UploadFile)

	// Handle the list of files
	r.GET("/api/files", controllers.ListFiles)

	// Create new folder
	r.POST("/api/createFolder", controllers.CreateFolder)

	// Download a file
	r.GET("/api/download", controllers.DownloadFile)

	// Delete a file
	r.DELETE("/api/delete", controllers.DeleteFile)

	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
