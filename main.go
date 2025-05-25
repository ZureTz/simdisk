package main

import (
	"github.com/ZureTz/simdisk/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// Set up a simple GET endpoint
	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	// Handle static files
	r.Static("/static", "./static")

	// Handle file uploads
	r.POST("/api/upload", controllers.UploadFile)

	// Handle the list of files
	r.GET("/api/files", controllers.ListFiles)

	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
