package endpoint

import (
	"github.com/go-chi/chi/v5"

	"http101/internal/application/service"
	"http101/internal/infrastructure/repository"
)

type TaskEndpoint struct{}

func (rs TaskEndpoint) Routes() chi.Router {
	taskRepo := repository.NewTaskRepository()
	taskService := service.NewTaskService(taskRepo)

	router := chi.NewRouter()
	router.Get("/get-all", taskService.HandleGetTasks)
	router.Post("/add-new", taskService.HandleAddTask)

	return router
}
