package handlers

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/aminelch/go-user-generator/models"
	"net/http"
	"time"
)

func HealthHandler(c *gin.Context) {
	sqliteStatus, dbSize := models.GetSQLiteStatus(DbPath)
	uptime := GetUptime()

	globalStatus := "OK"
	switch sqliteStatus {
	case "error":
		globalStatus = "DEGRADED"
	case "unavailable":
		globalStatus = "UNAVAILABLE"
	}

	response := models.HealthResponse{
		Hostname:  GetHostname(),
		Uptime:    uptime,
		Status:    globalStatus,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Version:   AppVersion,
		Dependencies: models.Dependencies{
			SQLite: models.SQLiteInfo{
				Status: sqliteStatus,
				File:   DbPath,
				SizeMB: dbSize,
			},
		},
	}

	c.JSON(http.StatusOK, response)
}
