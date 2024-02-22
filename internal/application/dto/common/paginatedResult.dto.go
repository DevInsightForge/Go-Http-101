package common_dto

import "math"

// PaginatedResultDto represents the paginated result structure.
type PaginatedResultDto struct {
	TotalRecords    int64       `json:"totalRecords"`
	TotalPages      int         `json:"totalPages"`
	CurrentPage     int         `json:"currentPage"`
	PageSize        int         `json:"pageSize"`
	HasPreviousPage bool        `json:"hasPreviousPage"`
	HasNextPage     bool        `json:"hasNextPage"`
	Data            interface{} `json:"data"`
}

func NewPaginatedResultDto(totalRecords int64, currentPage int, pageSize int, data []interface{}) *PaginatedResultDto {
	totalPages := int(math.Ceil(float64(totalRecords) / float64(pageSize)))
	hasPreviousPage := currentPage > 1
	hasNextPage := currentPage*pageSize < int(totalRecords)

	return &PaginatedResultDto{
		TotalRecords:    totalRecords,
		TotalPages:      totalPages,
		CurrentPage:     currentPage,
		PageSize:        pageSize,
		HasPreviousPage: hasPreviousPage,
		HasNextPage:     hasNextPage,
		Data:            data,
	}
}
