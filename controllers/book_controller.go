package controllers

import (
	"byfood-test/config"
	"byfood-test/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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

	rows, err := config.DB.Query("SELECT * FROM books ORDER BY year DESC LIMIT $1 OFFSET $2", pageSize, offset)
	if err != nil {
		config.Log.WithError(err).Error("Error fetching books")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching books"})
		return
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year, &book.CreatedAt, &book.UpdatedAt); err != nil {
			config.Log.WithError(err).Error("Error scanning book")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning book"})
			return
		}
		books = append(books, book)
	}

	var total int
	err = config.DB.QueryRow("SELECT COUNT(*) FROM books").Scan(&total)
	if err != nil {
		config.Log.WithError(err).Error("Error counting books")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error counting books"})
		return
	}

	paginationInfo := struct {
		Limit      int `json:"limit"`
		Page       int `json:"page"`
		TotalCount int `json:"total_count"`
	}{
		Limit:      pageSize,
		Page:       page,
		TotalCount: total,
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

	query := "INSERT INTO books (title, author, year) VALUES ($1, $2, $3) RETURNING id"
	err := config.DB.QueryRow(query, book.Title, book.Author, book.Year).Scan(&book.ID)
	if err != nil {
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
	query := "SELECT id, title, author, year FROM books WHERE id = $1"
	err = config.DB.QueryRow(query, id).Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	if err != nil {
		config.Log.WithError(err).Error("Book not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
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

	query := "UPDATE books SET title = $1, author = $2, year = $3 WHERE id = $4"
	_, err = config.DB.Exec(query, book.Title, book.Author, book.Year, id)
	if err != nil {
		config.Log.WithError(err).Error("Error updating book")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating book"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book successfully updated", "data": book})
}
