package endpoint

import (
	task_handler "http101/internal/application/handler"
	"net/http"
)

func RegisterTaskEndpoints(server *http.ServeMux) {
	server.HandleFunc(http.MethodGet+" /", func(w http.ResponseWriter, r *http.Request) {
		task_handler.HandleGetTasks(w, r)
	})

	server.HandleFunc(http.MethodPost+" /add/", func(w http.ResponseWriter, r *http.Request) {
		task_handler.HandleAddTask(w, r)
	})
}
