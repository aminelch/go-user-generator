package main

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
    var err error
    DB, err = gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
    if err != nil {
        Logger.Fatal("Failed to connect to the database:", err)
    }

    err = DB.AutoMigrate(&User{})
    if err != nil {
        Logger.Fatal("Database migration failed:", err)
    }

    Logger.Info("Database migrated successfully!")
    SeedUsers(DB)
}
