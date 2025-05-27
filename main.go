package main

import (
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/ZureTz/simdisk/controllers"
	"github.com/ZureTz/simdisk/utils"
)

func main() {
	// Initialize the configuration
	utils.InitConfig()

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

	// Get port from config
	portNumber := strconv.Itoa(utils.Config.Server.Port)

	// listen and serve on given port from config
	r.Run(":" + portNumber)
}
