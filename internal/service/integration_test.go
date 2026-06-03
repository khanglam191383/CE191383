package service

import (
	"testing"

	"ce191383/task_management/internal/entity"
	"ce191383/task_management/internal/repository"
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

func TestFullFlow(t *testing.T) {

	db, err := sql.Open(
		"postgres",
		"host=localhost port=5433 user=admin password=admin123 dbname=taskdb sslmode=disable",
	)

	if err != nil {
		t.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		t.Fatal(err)
	}
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6378",
	})
	defer rdb.Close()

	userRepo := repository.NewUserRepository(db)
	projectRepo := repository.NewProjectRepository(db)
	taskRepo := repository.NewTaskRepository(db)
	commentRepo := repository.NewCommentRepository(db)

	userService := NewUserService(userRepo)
	projectService := NewProjectService(projectRepo)
	taskService := NewTaskService(taskRepo, rdb)
	commentService := NewCommentService(commentRepo)

	// cleanup
	db.Exec("DELETE FROM comments")
	db.Exec("DELETE FROM tasks")
	db.Exec("DELETE FROM projects")
	db.Exec("DELETE FROM users")

	//----------------------------------
	// REGISTER
	//----------------------------------

	user, err := userService.Register(entity.User{
		Email:    "integration@test.com",
		Password: "123456",
		FullName: "Integration User",
	})

	if err != nil {
		t.Fatalf("register failed: %v", err)
	}

	if user.ID == 0 {
		t.Fatal("user id not generated")
	}

	//----------------------------------
	// LOGIN
	//----------------------------------

	token, err := userService.Login(
		"integration@test.com",
		"123456",
	)

	if err != nil {
		t.Fatalf("login failed: %v", err)
	}

	if token == "" {
		t.Fatal("token is empty")
	}

	//----------------------------------
	// CREATE PROJECT
	//----------------------------------

	project, err := projectService.Create(
		entity.Project{
			Name:        "Demo Project",
			Description: "Integration Test",
			OwnerID:     user.ID,
		},
	)

	if err != nil {
		t.Fatalf("create project failed: %v", err)
	}

	if project.ID == 0 {
		t.Fatal("project id not generated")
	}

	//----------------------------------
	// CREATE TASK
	//----------------------------------

	task, err := taskService.Create(
		entity.Task{
			ProjectID:   project.ID,
			Title:       "Learn Go",
			Description: "Integration task",
			Status:      "TODO",
			AssigneeID:  user.ID,
		},
	)

	if err != nil {
		t.Fatalf("create task failed: %v", err)
	}

	if task.ID == 0 {
		t.Fatal("task id not generated")
	}

	//----------------------------------
	// CREATE COMMENT
	//----------------------------------

	comment, err := commentService.Create(
		entity.Comment{
			TaskID:  task.ID,
			UserID:  user.ID,
			Content: "First comment",
		},
	)

	if err != nil {
		t.Fatalf("create comment failed: %v", err)
	}

	if comment.ID == 0 {
		t.Fatal("comment id not generated")
	}

	t.Log("Full integration flow passed")
}
