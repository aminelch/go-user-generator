package handlers

import (
	"github.com/gin-gonic/gin"
	"gitlab.com/aminelch/go-user-generator/models"
	"net/http"
)

func HomepageHandler(c *gin.Context) {
	hc.Logger.Info("Received request on /")

	response := models.HomePageResponse{
		Api:           AppName,
		Version:       AppVersion,
		Status:        "OK",
		Documentation: AppDocumentation,
		Author: models.Author{
			Name:    "Amine LOUHICHI",
			Website: "https://aminelch.cloud",
		},
	}

	c.JSON(http.StatusOK, response)
}
