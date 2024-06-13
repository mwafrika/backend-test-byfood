package main

import (
	"byfood-test-backend/config"
	"byfood-test-backend/controllers"
	"byfood-test-backend/docs"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	config.LoadEnvVariables()
	config.InitLogger()
	config.ConnectToDB()
	config.MigrateDatabase()
}

func main() {
	router := gin.Default()

	docs.SwaggerInfo.Title = "Book Management System API"
	docs.SwaggerInfo.Description = "This is a server for managing books."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	config := cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}

	router.Use(cors.New(config))

	api := router.Group("/api")
	{
		api.POST("/books", controllers.AddBook)
		api.GET("/books", controllers.GetBooks)
		api.GET("/books/:id", controllers.GetBookByID)
		api.PUT("/books/:id", controllers.UpdateBookByID)
		api.DELETE("/books/:id", controllers.DeleteBookByID)
		api.POST("/process_url", controllers.ProcessURL)
		api.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	router.Run()
}
