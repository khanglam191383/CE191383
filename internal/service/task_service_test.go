package service

import (
	"database/sql"
	"testing"

	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"

	"ce191383/task_management/internal/entity"
	"ce191383/task_management/internal/repository"
)

func setupTestDB() *sql.DB {
	connStr := "postgres://postgres:password@localhost:5432/task_management_test?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err)
	}

	return db
}

func cleanDB(t *testing.T, db *sql.DB) {
	if _, err := db.Exec("DELETE FROM tasks"); err != nil {
		t.Fatalf("failed to clean db: %v", err)
	}
}

func TestCreateTask_Success(t *testing.T) {
	db := setupTestDB()
	defer db.Close()

	cleanDB(t, db)

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	defer rdb.Close()

	repo := repository.NewTaskRepository(db)
	service := NewTaskService(repo, rdb)

	task := entity.Task{
		ProjectID:   1,
		Title:       "Learn Go",
		Description: "test",
		Status:      "TODO",
		AssigneeID:  1,
	}

	result, err := service.Create(task)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result.Title != "Learn Go" {
		t.Errorf("expected Learn Go, got %s", result.Title)
	}
}

func TestCreateTask_InvalidTitle(t *testing.T) {
	db := setupTestDB()
	defer db.Close()

	cleanDB(t, db)

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	defer rdb.Close()

	repo := repository.NewTaskRepository(db)
	service := NewTaskService(repo, rdb)

	task := entity.Task{
		ProjectID:  1,
		Title:      "",
		Status:     "TODO",
		AssigneeID: 1,
	}

	_, err := service.Create(task)

	if err == nil {
		t.Errorf("expected error for empty title")
	}
}

func TestCreateTask_InvalidStatus(t *testing.T) {
	db := setupTestDB()
	defer db.Close()

	cleanDB(t, db)

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	defer rdb.Close()

	repo := repository.NewTaskRepository(db)
	service := NewTaskService(repo, rdb)

	task := entity.Task{
		ProjectID:  1,
		Title:      "Test",
		Status:     "INVALID",
		AssigneeID: 1,
	}

	_, err := service.Create(task)

	if err == nil {
		t.Errorf("expected error for invalid status")
	}
}

func TestUpdateTask_Success(t *testing.T) {
	db := setupTestDB()
	defer db.Close()

	cleanDB(t, db)

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	defer rdb.Close()

	repo := repository.NewTaskRepository(db)
	service := NewTaskService(repo, rdb)

	created, err := service.Create(entity.Task{
		ProjectID:  1,
		Title:      "Old",
		Status:     "TODO",
		AssigneeID: 1,
	})

	if err != nil {
		t.Fatalf("setup failed: %v", err)
	}

	err = service.Update(created.ID, entity.Task{
		ProjectID:  1,
		Title:      "New",
		Status:     "DONE",
		AssigneeID: 1,
	})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestUpdateTask_InvalidTitle(t *testing.T) {
	db := setupTestDB()
	defer db.Close()

	cleanDB(t, db)

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	defer rdb.Close()

	repo := repository.NewTaskRepository(db)
	service := NewTaskService(repo, rdb)

	
	created, err := service.Create(entity.Task{
		ProjectID:  1,
		Title:      "Temp",
		Status:     "TODO",
		AssigneeID: 1,
	})

	if err != nil {
		t.Fatalf("setup failed: %v", err)
	}

	err = service.Update(created.ID, entity.Task{
		ProjectID:  1,
		Title:      "",
		Status:     "TODO",
		AssigneeID: 1,
	})

	if err == nil {
		t.Errorf("expected error for empty title")
	}
}

func TestDeleteTask(t *testing.T) {
	db := setupTestDB()
	defer db.Close()

	cleanDB(t, db)

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	defer rdb.Close()

	repo := repository.NewTaskRepository(db)
	service := NewTaskService(repo, rdb)

	created, err := service.Create(entity.Task{
		ProjectID:  1,
		Title:      "To be deleted",
		Status:     "TODO",
		AssigneeID: 1,
	})

	if err != nil {
		t.Fatalf("setup failed: %v", err)
	}

	err = service.Delete(created.ID)

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}
