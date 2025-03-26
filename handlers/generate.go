package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gitlab.com/aminelch/go-user-generator/models"
	"math/rand"
	"net/http"
)

func GenerateHandler(c *gin.Context) {
	hc.Logger.Info("Received request on /generate")

	var users []models.User
	if result := hc.DB.Unscoped().Find(&users); result.Error != nil {
		hc.Logger.Error("Database error:", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	if result := hc.DB.Find(&users); result.Error != nil || len(users) == 0 {
		hc.Logger.Warn("No users found in the database")
		c.JSON(http.StatusNotFound, gin.H{"error": "No user found"})
		return
	}

	randomUser := users[rand.Intn(len(users))]
	hc.Logger.WithFields(logrus.Fields{
		"uuid":  randomUser.Uuid,
		"name":  randomUser.Name,
		"email": randomUser.Email,
	}).Info("Selected user")

	c.JSON(http.StatusOK, gin.H{
		"uuid":  randomUser.Uuid,
		"name":  randomUser.Name,
		"email": randomUser.Email,
	})
}
