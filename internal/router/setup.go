package router

import (
	"ce191383/task_management/internal/handler"
)

func SetupRoutes(
	taskHandler *handler.TaskHandler,
	projectHandler *handler.ProjectHandler,
	userHandler *handler.UserHandler,
	commentHandler *handler.CommentHandler,
) {

	RegisterTaskRoutes(taskHandler)

	RegisterProjectRoutes(projectHandler)

	RegisterUserRoutes(userHandler)

	RegisterCommentRoutes(commentHandler)

	RegisterHealthRouters()
}