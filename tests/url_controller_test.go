package tests

import (
	"byfood-test-backend/controllers"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/process_url", controllers.ProcessURL)
	return router
}

func TestProcessURLInvalidInput(t *testing.T) {
	router := setupRouter()

	requestBody := map[string]string{
		"url":       "",
		"operation": "all",
	}
	requestJSON, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("POST", "/process_url", bytes.NewBuffer(requestJSON))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code)
	var responseBody map[string]string
	err := json.Unmarshal(resp.Body.Bytes(), &responseBody)
	assert.NoError(t, err)
	assert.Equal(t, "Invalid input", responseBody["error"])
}

func TestProcessURL(t *testing.T) {
	router := setupRouter()

	requestBody := map[string]string{
		"url":       "https://BYFOOD.com/food-EXPeriences?query=abc/",
		"operation": "all",
	}
	requestJSON, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("POST", "/process_url", bytes.NewBuffer(requestJSON))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	var responseBody map[string]string
	err := json.Unmarshal(resp.Body.Bytes(), &responseBody)
	assert.NoError(t, err)
	assert.Equal(t, "https://www.byfood.com/food-experiences", responseBody["processed_url"])
}

func TestProcessURLCanonical(t *testing.T) {
	router := setupRouter()

	requestBody := map[string]string{
		"url":       "https://BYFOOD.com/food-EXPeriences?query=abc/",
		"operation": "canonical",
	}
	requestJSON, _ := json.Marshal(requestBody)
	req, _ := http.NewRequest("POST", "/process_url", bytes.NewBuffer(requestJSON))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	var responseBody map[string]string
	err := json.Unmarshal(resp.Body.Bytes(), &responseBody)
	assert.NoError(t, err)
	assert.Equal(t, "https://BYFOOD.com/food-EXPeriences", responseBody["processed_url"])
}
