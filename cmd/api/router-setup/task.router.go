package routersetup

import (
	task_handler "http101/internal/application/handler"
	"net/http"
)

func (rm *RouterMap) RegisterTaskEndpoints() {
	taskRouter := rm.routers.Group("/task")

	// Register handlers with the router
	taskRouter.Handle(http.MethodGet, "/get-all", task_handler.HandleGetTasks)
	taskRouter.Handle(http.MethodPost, "/add-new", task_handler.HandleAddTask)
}
