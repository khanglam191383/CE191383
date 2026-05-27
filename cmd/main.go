package main

import (
	"ce191383/task_management/internal/config"
	"ce191383/task_management/internal/handler"
	"ce191383/task_management/internal/middleware"
	"ce191383/task_management/internal/repository"
	"ce191383/task_management/internal/router"
	"ce191383/task_management/internal/service"
	"ce191383/task_management/internal/worker"
	"fmt"
	"net/http"
)

func main() {

	db, err := config.ConnectDB()

	redisClient := config.ConnectRedis()

	_ = redisClient

	if err != nil {
		panic(err)
	}

	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	projectRepository := repository.NewProjectRepository(db)
	projectService := service.NewProjectService(projectRepository)
	projectHandler := handler.NewProjectHandler(projectService)

	taskRepository := repository.NewTaskRepository(db)
	taskService := service.NewTaskService(taskRepository, redisClient)
	taskHandler := handler.NewTaskHandler(taskService)

	commentRepo := repository.NewCommentRepository(db)

	commentService := service.NewCommentService(commentRepo)

	commentHandler := handler.NewCommentHandler(commentService)

	router.SetupRoutes(taskHandler, projectHandler, userHandler, commentHandler)

	go worker.StartNotificationWorker(redisClient)

	fmt.Println("Server is running on port 8080...")

	handlerWithMiddleware := middleware.Logger(
		middleware.Recovery(
			middleware.RequestID(http.DefaultServeMux),
		),
	)

	http.ListenAndServe(":8080", handlerWithMiddleware)
}
