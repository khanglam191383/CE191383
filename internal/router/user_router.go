package router

import (
	"net/http"

	"ce191383/task_management/internal/handler"
)

func RegisterUserRoutes(userHandler *handler.UserHandler) {

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodGet {
			userHandler.GetAllUsers(w, r)
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

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodPost {
			userHandler.Register(w, r)
			return
		}

		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	})

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodPost {
			userHandler.Login(w, r)
			return
		}

		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	})
}
