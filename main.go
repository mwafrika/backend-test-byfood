package main

import (
	"byfood-test/config"
	"byfood-test/controllers"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariables()
	config.InitLogger()
	config.ConnectDatabase()
}

func main() {
	router := gin.Default()

	router.GET("/", controllers.Welcome)

	router.Run()
}
