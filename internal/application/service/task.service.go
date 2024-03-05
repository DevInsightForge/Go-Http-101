package service

import (
	"encoding/json"
	"net/http"

	"http101/internal/application/utility"
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
	utility.ParseQueryParams(queryParams, &findOptions)

	if tasks, taskErr := s.TaskRepo.FindAllWithOptions(findOptions); taskErr == nil {
		totalRecords, countErr := s.TaskRepo.Count()
		if countErr != nil {
			errResp := utility.NewErrorResult("Failed to count tasks", countErr.Error())
			utility.WriteJsonResponse(w, http.StatusBadRequest, errResp)
			return
		}

		response := utility.NewPaginatedResultDto(totalRecords, findOptions.Page, findOptions.PageSize, tasks)
		utility.WriteJsonResponse(w, http.StatusOK, response)
	} else {
		errResp := utility.NewErrorResult("Failed to retrieve tasks", taskErr.Error())
		utility.WriteJsonResponse(w, http.StatusBadRequest, errResp)
	}
}

// HandleAddTask method for adding a new task
func (s *TaskService) HandleAddTask(w http.ResponseWriter, r *http.Request) {
	var requestBody model.TaskModel
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		errResp := utility.NewErrorResult("Failed to decode JSON", err.Error())
		utility.WriteJsonResponse(w, http.StatusBadRequest, errResp)
		return
	}

	requestBody.SetAuditFieldsBeforeCreate("testUserId")
	insertedID, err := s.TaskRepo.Create(requestBody)
	if err != nil {
		errResp := utility.NewErrorResult("Failed to insert task", err.Error())
		utility.WriteJsonResponse(w, http.StatusBadRequest, errResp)
		return
	}

	requestBody.ID = insertedID
	response := utility.NewSuccessDataResult(requestBody)
	utility.WriteJsonResponse(w, http.StatusOK, response)
}
