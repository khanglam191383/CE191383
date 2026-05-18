package service

import (
	"testing"
	"ce191383/task_management/entity"
	"ce191383/task_management/repository"
)

func TestCreateTask(t *testing.T) {

	repo := repository.NewTaskRepository()

	service := NewTaskService(repo)

	task := entity.Task{
		ID:     1,
		Title:  "Learn Go",
		Status: "TODO",
	}

	result, err := service.Create(task)

	if err != nil {
		t.Errorf("expected no error")
	}

	if result.Title != "Learn Go" {
		t.Errorf("expected Learn Go")
	}
}