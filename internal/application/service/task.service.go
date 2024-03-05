package service

import (
	"encoding/json"
	"net/http"

	"http101/internal/application/utility"
	"http101/internal/domain/model"
	"http101/internal/infrastructure/repository"
	base_repository "http101/internal/infrastructure/repository/base"

	"github.com/go-chi/chi/v5"
	"github.com/stroiman/go-automapper"
)

type TaskService struct {
	TaskRepo *repository.TaskRepository
}

func NewTaskService() *TaskService {
	taskRepo := repository.NewTaskRepository()

	return &TaskService{
		TaskRepo: taskRepo,
	}
}

func (s *TaskService) HandleGetTasks(w http.ResponseWriter, r *http.Request) {
	findOptions := base_repository.FindFilterOptions{
		Page:     1,
		PageSize: 10,
		SortBy:   "_id",
		SortDir:  -1,
	}

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

func (s *TaskService) HandleGetTaskById(w http.ResponseWriter, r *http.Request) {
	var requestId = chi.URLParam(r, "taskId")

	task, err := s.TaskRepo.FindByID(requestId)
	if err != nil {
		errResp := utility.NewErrorResult("Failed to find task", err.Error())
		utility.WriteJsonResponse(w, http.StatusBadRequest, errResp)
		return
	}

	response := utility.NewSuccessDataResult(task)
	utility.WriteJsonResponse(w, http.StatusOK, response)
}

func (s *TaskService) HandleUpdateTaskById(w http.ResponseWriter, r *http.Request) {
	var requestId = chi.URLParam(r, "taskId")

	oldTask, err := s.TaskRepo.FindByID(requestId)
	if err != nil {
		errResp := utility.NewErrorResult("Failed to find task", err.Error())
		utility.WriteJsonResponse(w, http.StatusBadRequest, errResp)
		return
	}

	var requestBody model.TaskModel
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		errResp := utility.NewErrorResult("Failed to decode JSON", err.Error())
		utility.WriteJsonResponse(w, http.StatusBadRequest, errResp)
		return
	}

	automapper.MapLoose(requestBody, &oldTask)
	requestBody.SetAuditFieldsBeforeUpdate("testUserId2")

	if err := s.TaskRepo.Update(requestId, requestBody); err != nil {
		errResp := utility.NewErrorResult("Failed to update task", err.Error())
		utility.WriteJsonResponse(w, http.StatusBadRequest, errResp)
		return
	}

	response := utility.NewSuccessDataResult(requestBody)
	utility.WriteJsonResponse(w, http.StatusOK, response)
}
