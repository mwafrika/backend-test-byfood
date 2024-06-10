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
	config.MigrateDatabase()
}

func main() {
	router := gin.Default()

	router.POST("/books", controllers.AddBook)
	router.GET("/books", controllers.GetBooks)
	router.GET("/books/:id", controllers.GetBookByID)
	router.PUT("/books/:id", controllers.UpdateBookByID)
	router.DELETE("/books/:id", controllers.DeleteBookByID)
	router.POST("/process_url", controllers.ProcessURL)
	router.Run()
}
