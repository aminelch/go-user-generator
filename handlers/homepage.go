package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HomepageHandler(c *gin.Context) {
	hc.Logger.Info("Received request on /")
	c.JSON(http.StatusOK, gin.H{
		"api":           AppName,
		"version":       AppVersion,
		"status":        "OK",
		"documentation": AppDocumentation,
	})
}
