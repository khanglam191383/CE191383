package router

import (
	"net/http"

	"ce191383/task_management/internal/handler"
	"ce191383/task_management/internal/middleware"
	"ce191383/task_management/internal/websocket"
)

func RegisterCommentRoutes(commentHandler *handler.CommentHandler) {

	http.HandleFunc(
		"/comments",
		middleware.AuthMiddleware(func(
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
		}),
	)

	http.HandleFunc("/ws", websocket.HandleConnections)
}