package router

import(
	"net/http"

	"ce191383/task_management/internal/handler"
)

func RegisterHealthRouters() {
	http.HandleFunc(
		"/health",
		handler.HealthCheck,
	)
}