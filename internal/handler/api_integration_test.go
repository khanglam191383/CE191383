package handler

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	_ "github.com/lib/pq"

	"ce191383/task_management/internal/entity"
	"ce191383/task_management/internal/repository"
	"ce191383/task_management/internal/service"
)

func setupTestDB() *sql.DB {

	fmt.Println("TEST DB CONNECTING TO 5433")
	db, err := sql.Open(
		"postgres",
		"host=localhost port=5433 user=admin password=admin123 dbname=taskdb sslmode=disable",
	)

	if err != nil {
		panic(err)
	}

	return db
}

func setupTestUserHandler(t *testing.T) *UserHandler {

	db := setupTestDB()

	repo := repository.NewUserRepository(db)

	userService := service.NewUserService(repo)

	t.Cleanup(func() {
		db.Close()
	})

	return NewUserHandler(userService)
}

func createTestUser(db *sql.DB) entity.User {

	userRepo := repository.NewUserRepository(db)

	user, err := userRepo.Register(entity.User{
		Email: fmt.Sprintf(
			"user_%d@gmail.com",
			time.Now().UnixNano(),
		),
		PasswordHash: "$2a$10$dummyhash",
		FullName:     "Test User",
	})

	if err != nil {
		panic(err)
	}

	return user
}

func TestRegisterAPI(t *testing.T) {

	handler := setupTestUserHandler(t)

	body, _ := json.Marshal(entity.User{
		Email: fmt.Sprintf(
			"api_%d@gmail.com",
			time.Now().UnixNano(),
		),
		Password: "123456",
		FullName: "APIv2 User",
	})

	req := httptest.NewRequest(
		http.MethodPost,
		"/register",
		bytes.NewBuffer(body),
	)

	rr := httptest.NewRecorder()

	handler.Register(rr, req)

	t.Log("response:", rr.Body.String())

	if rr.Code != http.StatusOK {
		t.Fatalf(
			"expected 200 got %d",
			rr.Code,
		)
	}
}

func TestLoginAPI(t *testing.T) {

	handler := setupTestUserHandler(t)

	email := fmt.Sprintf(
		"login_%d@gmail.com",
		time.Now().UnixNano(),
	)

	registerBody, _ := json.Marshal(entity.User{
		Email:    email,
		Password: "123456",
		FullName: "Login User",
	})

	registerReq := httptest.NewRequest(
		http.MethodPost,
		"/register",
		bytes.NewBuffer(registerBody),
	)

	registerRR := httptest.NewRecorder()

	handler.Register(registerRR, registerReq)

	loginBody, _ := json.Marshal(map[string]string{
		"email":    email,
		"password": "123456",
	})

	req := httptest.NewRequest(
		http.MethodPost,
		"/login",
		bytes.NewBuffer(loginBody),
	)

	rr := httptest.NewRecorder()

	handler.Login(rr, req)

	t.Log("response:", rr.Body.String())

	if rr.Code != http.StatusOK {
		t.Fatalf(
			"expected 200 got %d",
			rr.Code,
		)
	}
}

func TestCreateProjectAPI(t *testing.T) {

	db := setupTestDB()
	defer db.Close()

	user := createTestUser(db)

	projectRepo := repository.NewProjectRepository(db)
	projectService := service.NewProjectService(projectRepo)
	projectHandler := NewProjectHandler(projectService)

	body := `{
		"name":"Project API",
		"description":"demo"
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/projects",
		strings.NewReader(body),
	)

	ctx := context.WithValue(
		req.Context(),
		"user_id",
		user.ID,
	)

	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()

	projectHandler.CreateProject(rr, req)

	t.Log("response:", rr.Body.String())

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200 got %d", rr.Code)
	}
}

func TestCreateTaskAPI(t *testing.T) {

	db := setupTestDB()
	defer db.Close()

	user := createTestUser(db)

	projectRepo := repository.NewProjectRepository(db)

	project, err := projectRepo.Create(entity.Project{
		Name:        "Project Test",
		Description: "Demo",
		OwnerID:     user.ID,
	})

	if err != nil {
		t.Fatal(err)
	}

	taskRepo := repository.NewTaskRepository(db)

	taskService := service.NewTaskService(
		taskRepo,
		nil,
	)

	taskHandler := NewTaskHandler(taskService)

	body := fmt.Sprintf(`{
		"project_id":%d,
		"title":"API Task",
		"description":"API Task Desc",
		"status":"TODO",
		"assignee_id":%d
	}`,
		project.ID,
		user.ID,
	)

	req := httptest.NewRequest(
		http.MethodPost,
		"/tasks",
		strings.NewReader(body),
	)

	rr := httptest.NewRecorder()

	taskHandler.CreateTask(rr, req)

	t.Log("response:", rr.Body.String())

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200 got %d", rr.Code)
	}
}

func TestCreateCommentAPI(t *testing.T) {

	db := setupTestDB()
	defer db.Close()

	user := createTestUser(db)

	projectRepo := repository.NewProjectRepository(db)

	project, _ := projectRepo.Create(entity.Project{
		Name:        "Comment Project",
		Description: "Demo",
		OwnerID:     user.ID,
	})

	taskRepo := repository.NewTaskRepository(db)

	task, _ := taskRepo.Create(entity.Task{
		ProjectID:   project.ID,
		Title:       "Comment Task",
		Description: "Demo",
		Status:      "TODO",
		AssigneeID:  user.ID,
	})

	commentRepo := repository.NewCommentRepository(db)
	commentService := service.NewCommentService(commentRepo)
	commentHandler := NewCommentHandler(commentService)

	body := fmt.Sprintf(`{
		"task_id":%d,
		"content":"hello"
	}`,
		task.ID,
	)

	req := httptest.NewRequest(
		http.MethodPost,
		"/comments",
		strings.NewReader(body),
	)

	ctx := context.WithValue(
		req.Context(),
		"user_id",
		user.ID,
	)

	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()

	commentHandler.CreateComment(rr, req)

	t.Log("response:", rr.Body.String())

	if rr.Code != http.StatusOK {
		t.Fatalf("expected 200 got %d", rr.Code)
	}
}
