// models/sqlite.go
package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"math"
	"os"
)

func GetSQLiteStatus(dbPath string) (string, float64) {
	fileInfo, err := os.Stat(dbPath)
	if err != nil {
		return "unavailable", 0
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return "error", roundSize(float64(fileInfo.Size()))
	}

	sqlDB, err := db.DB()
	if err != nil || sqlDB.Ping() != nil {
		return "error", roundSize(float64(fileInfo.Size()))
	}

	return "connected", roundSize(float64(fileInfo.Size()))
}

func roundSize(size float64) float64 {
	return math.Round((size/(1024*1024))*100) / 100 // Converti en MB et arrondi
}
