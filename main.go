package main

import (
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	InitLogger()
	Logger.Info("Starting API...")

	InitDatabase()

	r := gin.Default()

	r.GET("/health", HealthHandler)
	r.GET("/generate", GenerateHandler)

	Logger.Info("Server running on port 8080")
	err := r.Run(":8080")
	if err != nil {
		os.Exit(2)
	}
}
