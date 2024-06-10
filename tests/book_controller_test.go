package tests

import (
	"byfood-test/config"
	"byfood-test/controllers"
	"byfood-test/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupBookRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/books", controllers.AddBook)
	router.GET("/books", controllers.GetBooks)
	router.GET("/books/:id", controllers.GetBookByID)
	router.PUT("/books/:id", controllers.UpdateBookByID)
	router.DELETE("/books/:id", controllers.DeleteBookByID)
	return router
}

func initializeTestData() {
	config.DB.Exec("DELETE FROM books")
	config.DB.Exec("ALTER SEQUENCE books_id_seq RESTART WITH 1")

	books := []models.Book{
		{Title: "Book One", Author: "Author One", Year: 2001},
		{Title: "Book Two", Author: "Author Two", Year: 2002},
		{Title: "Book Three", Author: "Author Three", Year: 2003},
	}

	for _, book := range books {
		config.DB.Create(&book)
	}
}

func TestGetBooks(t *testing.T) {
	initializeTestData()
	router := setupBookRouter()

	req, _ := http.NewRequest("GET", "/books", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	var responseBody map[string]interface{}
	err := json.Unmarshal(resp.Body.Bytes(), &responseBody)
	assert.NoError(t, err)
	assert.Contains(t, responseBody, "data")
	assert.Len(t, responseBody["data"], 3)
}

func TestAddBook(t *testing.T) {
	router := setupBookRouter()

	book := models.Book{Title: "New Book", Author: "New Author", Year: 2024}
	requestJSON, _ := json.Marshal(book)
	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(requestJSON))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusCreated, resp.Code)
	var responseBody map[string]interface{}
	err := json.Unmarshal(resp.Body.Bytes(), &responseBody)
	assert.NoError(t, err)
	assert.Equal(t, "Book created successfully", responseBody["message"])
}

func TestEmptyBookTitle(t *testing.T) {
	router := setupBookRouter()

	book := models.Book{Title: "", Author: "New Author", Year: 2024}
	requestJSON, _ := json.Marshal(book)
	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(requestJSON))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusBadRequest, resp.Code)
	var responseBody map[string]interface{}
	err := json.Unmarshal(resp.Body.Bytes(), &responseBody)
	assert.NoError(t, err)
	assert.Equal(t, "Title cannot be empty", responseBody["error"])
}

func TestEmptyBookAuthor(t *testing.T) {
	router := setupBookRouter()

	book := models.Book{Title: "New Book", Author: "", Year: 2024}
	requestJSON, _ := json.Marshal(book)
	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(requestJSON))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusBadRequest, resp.Code)
	var responseBody map[string]interface{}
	err := json.Unmarshal(resp.Body.Bytes(), &responseBody)
	assert.NoError(t, err)
	assert.Equal(t, "Author cannot be empty", responseBody["error"])
}

func TestEmptyBookYear(t *testing.T) {
	router := setupBookRouter()

	book := models.Book{Title: "New Book", Author: "Josh", Year: 0}
	requestJSON, _ := json.Marshal(book)
	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(requestJSON))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusBadRequest, resp.Code)
	var responseBody map[string]interface{}
	err := json.Unmarshal(resp.Body.Bytes(), &responseBody)
	assert.NoError(t, err)
	assert.Equal(t, "Year cannot be empty", responseBody["error"])
}

func TestGetBookByID(t *testing.T) {
	initializeTestData()
	router := setupBookRouter()

	req, _ := http.NewRequest("GET", "/books/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	var book models.Book
	err := json.Unmarshal(resp.Body.Bytes(), &book)
	assert.NoError(t, err)
	assert.Equal(t, "Book One", book.Title)
	assert.Equal(t, "Author One", book.Author)
	assert.Equal(t, 2001, book.Year)
}

func TestGetBookNotFound(t *testing.T) {
	initializeTestData()
	router := setupBookRouter()

	req, _ := http.NewRequest("GET", "/books/1555", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)

	var responseBody map[string]interface{}
	err := json.Unmarshal(resp.Body.Bytes(), &responseBody)
	assert.NoError(t, err)

	assert.Equal(t, "Book not found", responseBody["error"])
}

func TestUpdateBookByID(t *testing.T) {
	initializeTestData()
	router := setupBookRouter()

	book := models.Book{Title: "Updated Book", Author: "Updated Author", Year: 2025}
	requestJSON, _ := json.Marshal(book)
	req, _ := http.NewRequest("PUT", "/books/1", bytes.NewBuffer(requestJSON))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	var responseBody map[string]interface{}
	err := json.Unmarshal(resp.Body.Bytes(), &responseBody)
	assert.NoError(t, err)
	assert.Equal(t, "Book successfully updated", responseBody["message"])
}

func TestDeleteBook(t *testing.T) {
	initializeTestData()
	router := setupBookRouter()

	req, _ := http.NewRequest("DELETE", "/books/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	var responseBody map[string]interface{}
	err := json.Unmarshal(resp.Body.Bytes(), &responseBody)
	assert.NoError(t, err)
	assert.Equal(t, "Book successfully deleted", responseBody["message"])
}

func TestMain(m *testing.M) {
	os.Setenv("APP_ENV", "test")
	config.LoadEnvVariables()
	config.SetupTestDB()
	code := m.Run()
	os.Exit(code)
}
