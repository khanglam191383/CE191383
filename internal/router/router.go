package router

import (
	"net/http"
	"ce191383/task_management/internal/handler"
	"ce191383/task_management/internal/middleware"
	"ce191383/task_management/internal/websocket"
)

func SetupRoutes(taskHandler *handler.TaskHandler, projectHandler *handler.ProjectHandler, userHandler *handler.UserHandler, commentHandler *handler.CommentHandler) {

	
	http.HandleFunc("/tasks", middleware.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodGet {
			taskHandler.GetAllTasks(w, r)
			return
		}

		if r.Method == http.MethodPost {
			taskHandler.CreateTask(w, r)
			return
		}

		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}))

	http.HandleFunc("/tasks/", middleware.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {

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
	}))

	
	http.HandleFunc("/projects", middleware.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodGet {
			projectHandler.GetAllProjects(w, r)
			return
		}

		if r.Method == http.MethodPost {
			projectHandler.CreateProject(w, r)
			return
		}

		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}))

	http.HandleFunc("/projects/", middleware.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {

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
	}))

http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		userHandler.GetAllUsers(w, r)
		return
	}

	if r.Method == http.MethodPost {
		userHandler.CreateUser(w, r)
		return
	}

	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
})

http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		userHandler.GetUserByID(w, r)
		return
	}

	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
})

http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request){

	if r.Method == http.MethodPost {
		userHandler.Login(w,r)
		return
	}

	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
})


http.HandleFunc("/ws", websocket.HandleConnections)


http.HandleFunc("/comments",middleware.AuthMiddleware(func(
	w http.ResponseWriter,
	r *http.Request,
) {
	if r.Method == http.MethodPost {
		commentHandler.CreateComment(w, r)
		return
	}

	if r.Method == http.MethodGet {
		commentHandler.GetCommentsByTaskID(w, r)
		return
	}

	http.Error(
		w,
		"method not allowed",
		http.StatusMethodNotAllowed,
	)
}))
}