package main

import (
    "github.com/gin-gonic/gin"
    "math/rand"
    "net/http"
)

func HealthHandler(c *gin.Context) {
    Logger.Info("Received request on /health")
    c.JSON(http.StatusOK, gin.H{
        "status":  "OK",
        "version": "1.0.0",
        "maintainer": "Amine LOUHICHI",
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

    Logger.Infof("Selected user: %s (%s)", randomUser.Name, randomUser.Email)
    c.JSON(http.StatusOK, randomUser)
}
