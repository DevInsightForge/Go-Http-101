package endpoint

import (
	"github.com/go-chi/chi/v5"

	"http101/internal/application/service"
)

type TaskEndpoint struct{}

func (rs TaskEndpoint) Routes() chi.Router {
	router := chi.NewRouter()
	taskService := service.NewTaskService()

	router.Get("/get-all", taskService.HandleGetTasks)
	router.Post("/create", taskService.HandleAddTask)
	router.Route("/{taskId}", func(subRouter chi.Router) {
		subRouter.Get("/get", taskService.HandleGetTaskById)
		subRouter.Put("/update", taskService.HandleUpdateTaskById)
	})

	return router
}
