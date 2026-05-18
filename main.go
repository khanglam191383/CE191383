package main

import (
	"fmt"
	"net/http"
	"ce191383/task_management/handler"
	"ce191383/task_management/repository"
	"ce191383/task_management/router"
	"ce191383/task_management/service"
	"ce191383/task_management/middleware"
)

func main(){

	taskRepository:= repository.NewTaskRepository()

	taskService := service.NewTaskService(taskRepository)
	taskHandler := handler.NewTaskHandler(taskService)
	router.SetupRoutes(taskHandler)

	fmt.Println("Server is running on port 8080")

	handlerWithMiddleware := middleware.Logger(
		middleware.Recovery(
			middleware.RequestID(http.DefaultServeMux),
		),
	)
	
	http.ListenAndServe(":8080", handlerWithMiddleware)
}