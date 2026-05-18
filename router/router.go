package router

import (
	"net/http"

	"ce191383/task_management/handler"
)

func SetupRoutes(taskHandler *handler.TaskHandler) {

	http.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodGet {

			taskHandler.GetAllTasks(w, r)

			return
		}

		if r.Method == http.MethodPost {

			taskHandler.CreateTask(w, r)

			return
		}

		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	})

	http.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodGet {

			taskHandler.GetTaskByID(w, r)

			return
		}

		if r.Method == http.MethodPut {

			taskHandler.UpdateTask(w, r)

			return
		}

		if r.Method == http.MethodDelete {

			taskHandler.DeleteTask(w, r)

			return
		}

		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	})
}