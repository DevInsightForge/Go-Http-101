package task_handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	common_dto "http101/internal/application/dto/common"
	"http101/internal/application/model"
	"http101/internal/application/repository"
	base_repository "http101/internal/application/repository/base"
)

func HandleGetTasks(w http.ResponseWriter, r *http.Request) {
	taskRepo, repoErr := repository.NewTaskRepository()
	if repoErr != nil {
		sendErrorResponse(w, repoErr, "Failed to initialize task repository", http.StatusInternalServerError)
		return
	}

	// filter options
	findOptions := base_repository.FindFilterOptions{
		Page:     1,
		PageSize: 1,
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

	if tasks, err := taskRepo.FindAllWithOptions(findOptions); err == nil {
		totalRecords, countErr := taskRepo.Count()
		if countErr != nil {
			sendErrorResponse(w, countErr, "Failed to count tasks", http.StatusInternalServerError)
			return
		}

		interfaceSlice := make([]interface{}, len(tasks))
		for i, task := range tasks {
			interfaceSlice[i] = task
		}
		response := common_dto.NewPaginatedResultDto(totalRecords, findOptions.Page, findOptions.PageSize, interfaceSlice)
		sendJSONResponse(w, http.StatusOK, response)
	} else {
		sendErrorResponse(w, err, "Failed to retrieve tasks", http.StatusInternalServerError)
	}
}

func HandleAddTask(w http.ResponseWriter, r *http.Request) {
	var requestBody model.TaskModel
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		sendErrorResponse(w, err, "Failed to decode JSON", http.StatusBadRequest)
		return
	}

	requestBody.CreatedAt = time.Now()
	requestBody.UpdatedAt = time.Now()

	taskRepo, repoErr := repository.NewTaskRepository()
	if repoErr != nil {
		sendErrorResponse(w, repoErr, "Failed to initialize task repository", http.StatusInternalServerError)
		return
	}

	insertedID, err := taskRepo.Create(requestBody)
	if err != nil {
		sendErrorResponse(w, err, "Failed to insert task", http.StatusInternalServerError)
		return
	}

	requestBody.ID = insertedID

	sendJSONResponse(w, http.StatusCreated, requestBody)
}

func sendJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func sendErrorResponse(w http.ResponseWriter, err error, message string, statusCode int) {
	log.Println(err.Error())
	errorReponse := common_dto.NewErrorResult(message, err.Error())
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(errorReponse)
}
