package controllers

import (
	"byfood-test-backend/config"
	"byfood-test-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type URLRequest struct {
	URL       string `json:"url" binding:"required"`
	Operation string `json:"operation" binding:"required"`
}

// ProcessURL godoc
// @Summary Process a URL
// @Description Process a URL for canonicalization or redirection
// @Tags URL Cleanup
// @Accept json
// @Produce json
// @Param url body URLRequest true "URL and Operation"
// @Success 200 {object} services.SuccessProcessURL
// @Failure 400 {object} services.ErrorResponse
// @Router /process_url [post]
func ProcessURL(c *gin.Context) {
	var request URLRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		config.Log.WithError(err).Error("Invalid input")
		c.JSON(http.StatusBadRequest, services.ErrorResponse{Error: "Invalid input"})
		return
	}

	processedURL, err := services.ProcessURL(request.URL, request.Operation)
	if err != nil {
		config.Log.WithError(err).Error("Error processing URL")
		c.JSON(http.StatusInternalServerError, services.ErrorResponse{Error: "Error processing URL"})
		return
	}

	c.JSON(http.StatusOK, services.SuccessProcessURL{ProcessedUrl: processedURL})
}
