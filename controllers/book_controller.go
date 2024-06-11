package controllers

import (
	"byfood-test/config"
	"byfood-test/models"
	"byfood-test/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetBooks handles the retrieval of books with pagination and ordering by year
// @Summary Get all books with pagination and ordering
// @Description Get details of all books with pagination, ordered by year in descending order
// @Tags Books
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param pageSize query int false "Number of items per page" default(10)
// @Success 200 {object} services.BookListResponse
// @Failure 400 {object} services.ErrorResponse
// @Failure 500 {object} services.ErrorResponse
// @Router /books [get]
func GetBooks(c *gin.Context) {

	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("pageSize", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, services.ErrorResponse{Error: "Invalid page parameter. Page must be a positive integer"})
		return
	}
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		c.JSON(http.StatusBadRequest, services.ErrorResponse{Error: "Invalid pageSize parameter. Page size must be a positive integer"})
		return
	}

	offset := (page - 1) * pageSize

	var books []models.Book

	if err := config.DB.Order("year DESC").Limit(pageSize).Offset(offset).Find(&books).Error; err != nil {
		config.Log.WithError(err).Error("Error fetching books")
		c.JSON(http.StatusInternalServerError, services.ErrorResponse{Error: "Error fetching books"})
		return
	}

	var total int64
	if err := config.DB.Model(&models.Book{}).Count(&total).Error; err != nil {
		config.Log.WithError(err).Error("Error counting books")
		c.JSON(http.StatusInternalServerError, services.ErrorResponse{Error: "Error counting books"})
		return
	}

	paginationInfo := services.Pagination{
		Limit:      pageSize,
		Page:       page,
		TotalCount: total,
	}

	c.JSON(http.StatusOK, services.BookListResponse{Data: books, Pagination: paginationInfo})
}

// AddBook handles adding a new book to the database
// @Summary Add a new book
// @Description Add a new book to the database
// @Tags Books
// @Accept json
// @Produce json
// @Param book body models.Book true "Book to add"
// @Success 201 {object} services.BookResponse
// @Failure 400 {object} services.ErrorResponse
// @Failure 500 {object} services.ErrorResponse
// @Router /books [post]
func AddBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		config.Log.WithError(err).Error("Invalid input")
		c.JSON(http.StatusBadRequest, services.ErrorResponse{Error: "Invalid input"})
		return
	}

	if book.Title == "" {
		c.JSON(http.StatusBadRequest, services.ErrorResponse{Error: "Title cannot be empty"})
		return
	}

	if book.Author == "" {
		c.JSON(http.StatusBadRequest, services.ErrorResponse{Error: "Author cannot be empty"})
		return
	}

	if book.Year == 0 {
		c.JSON(http.StatusBadRequest, services.ErrorResponse{Error: "Year cannot be empty"})
		return
	}

	if err := config.DB.Create(&book).Error; err != nil {
		config.Log.WithError(err).Error("Error adding book")
		c.JSON(http.StatusInternalServerError, services.ErrorResponse{Error: "Error adding book"})
		return
	}
	c.JSON(http.StatusCreated, services.BookResponse{Message: "Book created successfully", Data: book})
}

// GetBookByID handles retrieving a book by its ID
// @Summary Get a book by ID
// @Description Get details of a specific book by its ID
// @Tags Books
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} models.Book
// @Failure 400 {object} services.ErrorResponse
// @Failure 404 {object} services.ErrorResponse
// @Router /books/{id} [get]
func GetBookByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		config.Log.WithError(err).Error("Invalid ID")
		c.JSON(http.StatusBadRequest, services.ErrorResponse{Error: "Invalid ID"})
		return
	}

	var book models.Book
	if err := config.DB.First(&book, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			config.Log.WithError(err).Error("Book not found")
			c.JSON(http.StatusNotFound, services.ErrorResponse{Error: "Book not found"})
		} else {
			config.Log.WithError(err).Error("Error fetching book")
			c.JSON(http.StatusInternalServerError, services.ErrorResponse{Error: "Error fetching book"})
		}
		return
	}
	c.JSON(http.StatusOK, book)
}

// UpdateBookByID handles updating a book by its ID
// @Summary Update a book by ID
// @Description Update the details of a specific book by its ID
// @Tags Books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Param book body models.Book true "Book data to update"
// @Success 200 {object} services.BookResponse
// @Failure 400 {object} services.ErrorResponse
// @Failure 404 {object} services.ErrorResponse
// @Failure 500 {object} services.ErrorResponse
// @Router /books/{id} [put]
func UpdateBookByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		config.Log.WithError(err).Error("Invalid ID")
		c.JSON(http.StatusBadRequest, services.ErrorResponse{Error: "Invalid ID"})
		return
	}

	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		config.Log.WithError(err).Error("Invalid input")
		c.JSON(http.StatusBadRequest, services.ErrorResponse{Error: "Invalid input"})
		return
	}

	var existingBook models.Book
	if err := config.DB.First(&existingBook, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			config.Log.WithError(err).Error("Book not found")
			c.JSON(http.StatusNotFound, services.ErrorResponse{Error: "Book not found"})
		} else {
			config.Log.WithError(err).Error("Error fetching book")
			c.JSON(http.StatusInternalServerError, services.ErrorResponse{Error: "Error fetching book"})
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
		c.JSON(http.StatusInternalServerError, services.ErrorResponse{Error: "Error updating book"})
		return
	}
	c.JSON(http.StatusOK, services.BookResponse{Message: "Book successfully updated", Data: book})
}

// DeleteBookByID handles deleting a book by its ID
// @Summary Delete a book by ID
// @Description Delete a specific book by its ID
// @Tags Books
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} services.SuccessMessage
// @Failure 400 {object} services.ErrorResponse
// @Failure 404 {object} services.ErrorResponse
// @Router /books/{id} [delete]
func DeleteBookByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		config.Log.WithError(err).Error("Invalid ID")
		c.JSON(http.StatusBadRequest, services.ErrorResponse{Error: "Invalid ID"})
		return
	}

	var book models.Book
	if err := config.DB.First(&book, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			config.Log.WithError(err).Error("Book not found")
			c.JSON(http.StatusNotFound, services.ErrorResponse{Error: "Book not found"})
		} else {
			config.Log.WithError(err).Error("Error fetching book")
			c.JSON(http.StatusInternalServerError, services.ErrorResponse{Error: "Error fetching book"})
		}
		return
	}

	if err := config.DB.Delete(&book).Error; err != nil {
		config.Log.WithError(err).Error("Error deleting book")
		c.JSON(http.StatusInternalServerError, services.ErrorResponse{Error: "Error deleting book"})
		return
	}

	c.JSON(http.StatusOK, services.SuccessMessage{Message: "Book successfully deleted"})
}
