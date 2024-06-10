package main

import (
	"byfood-test/config"
	"byfood-test/controllers"

	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVariables()
	config.InitLogger()
	config.ConnectToDB()
	config.MigrateDatabase(config.DB)
}

func main() {
	router := gin.Default()

	router.GET("/books", controllers.GetBooks)
	router.POST("/books", controllers.AddBook)
	router.GET("/books/:id", controllers.GetBookByID)
	router.PUT("/books/:id", controllers.UpdateBookByID)
	router.Run()
}
