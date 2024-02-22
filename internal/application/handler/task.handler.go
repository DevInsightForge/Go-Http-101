package task_handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"http101/internal/application/model"
	"http101/internal/application/repository"
	base_repository "http101/internal/application/repository/base"
	response_utility "http101/internal/application/utility"
)

func HandleGetTasks(w http.ResponseWriter, r *http.Request) {
	taskRepo, repoErr := repository.NewTaskRepository()
	if repoErr != nil {
		errResp := response_utility.NewErrorResult("Failed to initialize task repository", repoErr.Error())
		response_utility.WriteJsonResponse(w, http.StatusBadRequest, errResp)
		return
	}

	// filter options
	findOptions := base_repository.FindFilterOptions{
		Page:     1,
		PageSize: 10,
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

	if tasks, taskErr := taskRepo.FindAllWithOptions(findOptions); taskErr == nil {
		totalRecords, countErr := taskRepo.Count()
		if countErr != nil {
			errResp := response_utility.NewErrorResult("Failed to count tasks", countErr.Error())
			response_utility.WriteJsonResponse(w, http.StatusBadRequest, errResp)
			return
		}

		response := response_utility.NewPaginatedResultDto[[]model.TaskModel](totalRecords, findOptions.Page, findOptions.PageSize, tasks)
		response_utility.WriteJsonResponse(w, http.StatusOK, response)
	} else {
		errResp := response_utility.NewErrorResult("Failed to retrieve tasks", taskErr.Error())
		response_utility.WriteJsonResponse(w, http.StatusBadRequest, errResp)
	}
}

func HandleAddTask(w http.ResponseWriter, r *http.Request) {
	var requestBody model.TaskModel
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		errResp := response_utility.NewErrorResult("Failed to decode JSON", err.Error())
		response_utility.WriteJsonResponse(w, http.StatusBadRequest, errResp)
		return
	}

	requestBody.CreatedAt = time.Now()
	requestBody.UpdatedAt = time.Now()

	taskRepo, repoErr := repository.NewTaskRepository()
	if repoErr != nil {
		errResp := response_utility.NewErrorResult("Failed to initialize task repository", repoErr.Error())
		response_utility.WriteJsonResponse(w, http.StatusBadRequest, errResp)
		return
	}

	insertedID, err := taskRepo.Create(requestBody)
	if err != nil {
		errResp := response_utility.NewErrorResult("Failed to insert task", err.Error())
		response_utility.WriteJsonResponse(w, http.StatusBadRequest, errResp)
		return
	}

	requestBody.ID = insertedID
	response := response_utility.NewSuccessDataResult[model.TaskModel](requestBody)
	response_utility.WriteJsonResponse(w, http.StatusOK, response)
}
