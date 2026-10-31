package main

import (
	"commitment-stake/config"
	"commitment-stake/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	config.ConnectDB()
	r := gin.Default()
	r.POST("/api/generate-qr", controllers.GenerateQR)
	r.POST("/api/events", controllers.CreateEvent)
	r.GET("/api/events", controllers.GetEvents)

	r.Run(":8080")
}