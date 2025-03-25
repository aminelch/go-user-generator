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

	r.GET("/", HomepageHandler)
	r.GET("/health", HealthHandler)
	r.GET("/generate", GenerateHandler)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	Logger.Info("Server running on port " + port)

	err := r.Run(":" + port)
	if err != nil {
		Logger.Fatalf("Failed to start server")
		os.Exit(2)
	}
}
