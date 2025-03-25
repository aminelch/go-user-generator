package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var (
	startTime  = time.Now()
	appVersion = getEnv("APP_VERSION", "1.0.0")
)

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getHostname() string {
	if name, err := os.Hostname(); err == nil {
		return name
	}
	return getEnv("HOSTNAME", "unknown-server")
}

func HealthHandler(c *gin.Context) {
	Logger.Info("Received request on /health")
	c.JSON(http.StatusOK, gin.H{
		"hostname":  getHostname(),
		"uptime":    startTime.UTC().Format(time.RFC3339),
		"status":    "OK",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
		"version":   appVersion, // Version dynamique
	})
}

func GenerateHandler(c *gin.Context) {
	Logger.Info("Received request on /generate")

	var users []User
	result := DB.Find(&users)
	if result.Error != nil || len(users) == 0 {
		Logger.Warn("No users found in the database")
		c.JSON(http.StatusNotFound, gin.H{"error": "No user found"})
		return
	}

	randomUser := users[rand.Intn(len(users))]

	Logger.Infof("Selected user: %s %s (%s)", randomUser.Uuid, randomUser.Name, randomUser.Email)
	c.JSON(http.StatusOK, randomUser)
}
