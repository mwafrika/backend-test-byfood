package services

import "byfood-test/models"

type ErrorResponse struct {
	Error string `json:"error"`
}

type BookListResponse struct {
	Data       []models.Book `json:"data"`
	Pagination Pagination    `json:"pagination"`
}

type Pagination struct {
	Limit      int   `json:"limit"`
	Page       int   `json:"page"`
	TotalCount int64 `json:"total_count"`
}

type BookResponse struct {
	Message string      `json:"message"`
	Data    models.Book `json:"data"`
}

type SuccessMessage struct {
	Message string `json:"message"`
}

type SuccessProcessURL struct {
	ProcessedUrl string `json:"processed_url"`
}
