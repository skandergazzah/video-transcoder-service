package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/skandergazzah/transcode/controller"
)

func main() {
	router := gin.Default()

	if err := os.MkdirAll("/app/uploads", os.ModePerm); err != nil {
		panic("Failed to create uploads directory: " + err.Error())
	}

	router.POST("/transcode", controller.Transcode)
	if err := router.Run(":9000"); err != nil {
		panic("Failed to start server: " + err.Error())
	}
}
