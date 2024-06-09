package controllers

import (
	"byfood-test/config"
	"byfood-test/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	rows, err := config.DB.Query("SELECT * FROM books")
	if err != nil {
		config.Log.WithError(err).Error("Error fetching books")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching books"})
		return
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year); err != nil {
			config.Log.WithError(err).Error("Error scanning book")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning book"})
			return
		}
		books = append(books, book)
	}
	c.JSON(http.StatusOK, gin.H{"message": "Books fetched successfully", "data": books})
}
