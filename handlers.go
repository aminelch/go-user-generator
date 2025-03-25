package main

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/aminelch/go-user-generator/models"
	"math"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var (
	appName          = getEnv("APP_NAME", "user-generator-api")
	appDocumentation = getEnv("APP_DOC", "/docs")
	appVersion       = getEnv("APP_VERSION", "1.0.0")
	startTime        = time.Now()
	dbPath           = getEnv("SQLITE_DB_PATH", "data.db")
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
	sqliteStatus, dbSize := models.GetSQLiteStatus(dbPath)
	uptime := math.Round(time.Since(startTime).Minutes()) // Uptime en minutes arrondi

	globalStatus := "OK"
	if sqliteStatus == "error" {
		globalStatus = "DEGRADED"
	} else if sqliteStatus == "unavailable" {
		globalStatus = "UNAVAILABLE"
	}

	response := models.HealthResponse{
		Hostname:  getHostname(),
		Uptime:    uptime,
		Status:    globalStatus,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Version:   appVersion,
		Dependencies: models.Dependencies{
			SQLite: models.SQLiteInfo{
				Status: sqliteStatus,
				File:   dbPath,
				SizeMB: dbSize,
			},
		},
	}

	c.JSON(http.StatusOK, response)
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

func HomepageHandler(c *gin.Context) {
	Logger.Info("Received request on /")
	c.JSON(http.StatusOK, gin.H{
		"api":           appName,
		"version":       appVersion,
		"status":        "OK",
		"documentation": appDocumentation,
	})
}
