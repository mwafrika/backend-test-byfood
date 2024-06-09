package controllers

import (
	"byfood-test/config"

	"github.com/gin-gonic/gin"
)

func Welcome(c *gin.Context) {
	config.Log.Info("Welcome to the server")
	c.JSON(200, gin.H{
		"message": "Welcome to ByFood API!",
	})
}
