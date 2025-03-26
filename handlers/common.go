package handlers

import (
	"math"
	"os"
	"time"
)

var (
	AppName          = GetEnv("APP_NAME", "user-generator-api")
	AppDocumentation = GetEnv("APP_DOC", "/docs")
	AppVersion       = GetEnv("APP_VERSION", "1.0.0")
	StartTime        = time.Now()
	DbPath           = GetEnv("SQLITE_DB_PATH", "data.db")
)

func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func GetHostname() string {
	if name, err := os.Hostname(); err == nil {
		return name
	}
	return GetEnv("HOSTNAME", "unknown-server")
}

func GetUptime() float64 {
	return math.Round(time.Since(StartTime).Minutes())
}
