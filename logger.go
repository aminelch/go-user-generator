package main

import (
    "os"
    "github.com/sirupsen/logrus"
)

var Logger = logrus.New()

func InitLogger() {
    file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        logrus.Fatalf("Failed to open log file: %v", err)
    }

    Logger.SetOutput(file)
    Logger.SetFormatter(&logrus.JSONFormatter{})
    Logger.SetLevel(logrus.InfoLevel)
}
