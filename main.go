package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/aminelch/go-user-generator/handlers"
	"os"
)

func main() {
	InitLogger()
	InitDatabase()

	handlers.InitHandlers(Logger, DB)
	r := gin.Default()

	r.GET("/", handlers.HomepageHandler)
	r.GET("/health", handlers.HealthHandler)
	r.GET("/generate", handlers.GenerateHandler)

	port := handlers.GetEnv("APP_PORT", "8080")
	Logger.Info("Server running on port " + port)

	if err := r.Run(":" + port); err != nil {
		Logger.Fatal("Failed to start server: ", err)
		os.Exit(2)
	}
}
