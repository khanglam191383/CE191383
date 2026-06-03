package router

import (
	"net/http"

	"ce191383/task_management/internal/handler"
	"ce191383/task_management/internal/middleware"
)

func RegisterProjectRoutes(projectHandler *handler.ProjectHandler) {

	http.HandleFunc(
		"/projects",
		middleware.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {

			if r.Method == http.MethodGet {
				projectHandler.GetAllProjects(w, r)
				return
			}

			if r.Method == http.MethodPost {
				projectHandler.CreateProject(w, r)
				return
			}

			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}),
	)

	http.HandleFunc(
		"/projects/",
		middleware.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {

			if r.Method == http.MethodGet {
				projectHandler.GetProjectByID(w, r)
				return
			}

			if r.Method == http.MethodPut {
				projectHandler.UpdateProject(w, r)
				return
			}

			if r.Method == http.MethodDelete {
				projectHandler.DeleteProject(w, r)
				return
			}

			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}),
	)
}