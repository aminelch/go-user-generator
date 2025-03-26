package main

import (
	"gitlab.com/aminelch/go-user-generator/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	DB = DB.Debug()

	if err != nil {
		Logger.Fatal("Failed to connect to the database:", err)
	}

	if err := DB.AutoMigrate(&models.User{}); err != nil {
		Logger.Fatal("Migration failed:", err)
	}

	Logger.Info("Database migrated successfully!")
	SeedUsers(DB)
}
