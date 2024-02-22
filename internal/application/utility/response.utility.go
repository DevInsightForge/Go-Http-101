package response_utility

import (
	"encoding/json"
	"math"
	"net/http"
)

type ResultDto[T any] struct {
	Success         bool   `json:"success"`
	Message         string `json:"message,omitempty"`
	TotalRecords    int64  `json:"totalRecords,omitempty"`
	TotalPages      int    `json:"totalPages,omitempty"`
	CurrentPage     int    `json:"currentPage,omitempty"`
	PageSize        int    `json:"pageSize,omitempty"`
	HasPreviousPage bool   `json:"hasPreviousPage,omitempty"`
	HasNextPage     bool   `json:"hasNextPage,omitempty"`
	Error           string `json:"error,omitempty"`
	Data            T      `json:"data,omitempty"`
}

func NewSuccessDataResult[T any](data T) *ResultDto[T] {
	return &ResultDto[T]{Success: true, Data: data}
}

func NewSuccessMessageResult[T any](message string) *ResultDto[T] {
	return &ResultDto[T]{Success: true, Message: message}
}

func NewErrorResult(message, err string) *ResultDto[interface{}] {
	return &ResultDto[interface{}]{Success: false, Message: message, Error: err}
}

func NewPaginatedResultDto[T any](totalRecords int64, currentPage, pageSize int, data T) *ResultDto[T] {
	totalPages := int(math.Ceil(float64(totalRecords) / float64(pageSize)))
	hasPreviousPage := currentPage > 1
	hasNextPage := currentPage*pageSize < int(totalRecords)

	return &ResultDto[T]{
		Success:         true,
		Data:            data,
		TotalRecords:    totalRecords,
		TotalPages:      totalPages,
		CurrentPage:     currentPage,
		PageSize:        pageSize,
		HasPreviousPage: hasPreviousPage,
		HasNextPage:     hasNextPage,
	}
}

func WriteJsonResponse(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
