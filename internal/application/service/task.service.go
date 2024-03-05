package service

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	response_utility "http101/internal/application/utility"
	"http101/internal/domain/model"
	"http101/internal/infrastructure/repository"
	base_repository "http101/internal/infrastructure/repository/base"
)

// TaskService struct to encapsulate task handlers
type TaskService struct {
	TaskRepo *repository.TaskRepository
}

// NewTaskService creates a new instance of TaskService
func NewTaskService() *TaskService {
	taskRepo := repository.NewTaskRepository()

	return &TaskService{
		TaskRepo: taskRepo,
	}
}

// HandleGetTasks method for getting tasks
func (s *TaskService) HandleGetTasks(w http.ResponseWriter, r *http.Request) {
	// filter options
	findOptions := base_repository.FindFilterOptions{
		Page:     1,
		PageSize: 10,
		SortBy:   "_id",
		SortDir:  -1,
	}

	// add filter options if provided in query parameters
	queryParams := r.URL.Query()

	for param, value := range queryParams {
		v := value[0]
		switch strings.ToLower(param) {
		case "sortby":
			if v != "" {
				findOptions.SortBy = v
			}
		case "sortdir":
			if v != "" {
				if sortDir, err := strconv.Atoi(v); err == nil {
					findOptions.SortDir = sortDir
				}
			}
		case "page":
			if v != "" {
				if page, err := strconv.Atoi(v); page > 0 && err == nil {
					findOptions.Page = page
				}
			}
		case "pagesize":
			if v != "" {
				if pageSize, err := strconv.Atoi(v); pageSize > 0 && err == nil {
					findOptions.PageSize = pageSize
				}
			}
		}
	}

	if tasks, taskErr := s.TaskRepo.FindAllWithOptions(findOptions); taskErr == nil {
		totalRecords, countErr := s.TaskRepo.Count()
		if countErr != nil {
			errResp := response_utility.NewErrorResult("Failed to count tasks", countErr.Error())
			response_utility.WriteJsonResponse(w, http.StatusBadRequest, errResp)
			return
		}

		response := response_utility.NewPaginatedResultDto(totalRecords, findOptions.Page, findOptions.PageSize, tasks)
		response_utility.WriteJsonResponse(w, http.StatusOK, response)
	} else {
		errResp := response_utility.NewErrorResult("Failed to retrieve tasks", taskErr.Error())
		response_utility.WriteJsonResponse(w, http.StatusBadRequest, errResp)
	}
}

// HandleAddTask method for adding a new task
func (s *TaskService) HandleAddTask(w http.ResponseWriter, r *http.Request) {
	var requestBody model.TaskModel
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		errResp := response_utility.NewErrorResult("Failed to decode JSON", err.Error())
		response_utility.WriteJsonResponse(w, http.StatusBadRequest, errResp)
		return
	}

	requestBody.SetAuditFieldsBeforeCreate("testUserId")
	insertedID, err := s.TaskRepo.Create(requestBody)
	if err != nil {
		errResp := response_utility.NewErrorResult("Failed to insert task", err.Error())
		response_utility.WriteJsonResponse(w, http.StatusBadRequest, errResp)
		return
	}

	requestBody.ID = insertedID
	response := response_utility.NewSuccessDataResult(requestBody)
	response_utility.WriteJsonResponse(w, http.StatusOK, response)
}
