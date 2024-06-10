package controllers

import (
	"byfood-test/config"
	"byfood-test/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetBooks(c *gin.Context) {

	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("pageSize", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page parameter. Page must be a positive integer"})
		return
	}
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid pageSize parameter. Page size must be a positive integer"})
		return
	}

	offset := (page - 1) * pageSize

	var books []models.Book

	if err := config.DB.Order("year DESC").Limit(pageSize).Offset(offset).Find(&books).Error; err != nil {
		config.Log.WithError(err).Error("Error fetching books")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching books"})
		return
	}

	var total int64
	if err := config.DB.Model(&models.Book{}).Count(&total).Error; err != nil {
		config.Log.WithError(err).Error("Error counting books")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error counting books"})
		return
	}

	paginationInfo := gin.H{
		"limit":       pageSize,
		"page":        page,
		"total_count": total,
	}

	c.JSON(http.StatusOK, gin.H{"data": books, "pagination": paginationInfo})
}

func AddBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		config.Log.WithError(err).Error("Invalid input")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if book.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title cannot be empty"})
		return
	}

	if book.Author == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Author cannot be empty"})
		return
	}

	if book.Year == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Year cannot be empty"})
		return
	}

	if err := config.DB.Create(&book).Error; err != nil {
		config.Log.WithError(err).Error("Error adding book")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error adding book"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Book created successfully", "data": book})
}

func GetBookByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		config.Log.WithError(err).Error("Invalid ID")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var book models.Book
	if err := config.DB.First(&book, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			config.Log.WithError(err).Error("Book not found")
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		} else {
			config.Log.WithError(err).Error("Error fetching book")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching book"})
		}
		return
	}
	c.JSON(http.StatusOK, book)
}

func UpdateBookByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		config.Log.WithError(err).Error("Invalid ID")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		config.Log.WithError(err).Error("Invalid input")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var existingBook models.Book
	if err := config.DB.First(&existingBook, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			config.Log.WithError(err).Error("Book not found")
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		} else {
			config.Log.WithError(err).Error("Error fetching book")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching book"})
		}
		return
	}

	if book.Title != "" {
		existingBook.Title = book.Title
	}
	if book.Author != "" {
		existingBook.Author = book.Author
	}
	if book.Year != 0 {
		existingBook.Year = book.Year
	}

	if err := config.DB.Save(&existingBook).Error; err != nil {
		config.Log.WithError(err).Error("Error updating book")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating book"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book successfully updated", "data": book})
}

func DeleteBookByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		config.Log.WithError(err).Error("Invalid ID")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var book models.Book
	if err := config.DB.First(&book, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			config.Log.WithError(err).Error("Book not found")
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		} else {
			config.Log.WithError(err).Error("Error fetching book")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching book"})
		}
		return
	}

	if err := config.DB.Delete(&book).Error; err != nil {
		config.Log.WithError(err).Error("Error deleting book")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting book"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book successfully deleted"})
}
