package controllers

import (
	"byfood-test/config"
	"byfood-test/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type URLRequest struct {
	URL       string `json:"url" binding:"required"`
	Operation string `json:"operation" binding:"required"`
}

func ProcessURL(c *gin.Context) {
	var request URLRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		config.Log.WithError(err).Error("Invalid input")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	processedURL, err := services.ProcessURL(request.URL, request.Operation)
	if err != nil {
		config.Log.WithError(err).Error("Error processing URL")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error processing URL"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"processed_url": processedURL})
}
