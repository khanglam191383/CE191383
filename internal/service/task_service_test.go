package service

import (
	"testing"

	"ce191383/task_management/internal/entity"
)

type mockTaskRepo struct{}

func (m *mockTaskRepo) Create(
	task entity.Task,
) (entity.Task, error) {

	task.ID = 1

	return task, nil
}

func (m *mockTaskRepo) GetAll() (
	[]entity.Task,
	error,
) {

	return []entity.Task{
		{
			ID:        1,
			ProjectID: 1,
			Title:     "Test Task",
			Status:    "TODO",
		},
	}, nil
}

func (m *mockTaskRepo) GetByID(
	id int,
) (*entity.Task, error) {

	return &entity.Task{
		ID:     id,
		Title:  "Test Task",
		Status: "TODO",
	}, nil
}

func (m *mockTaskRepo) Update(
	id int,
	task entity.Task,
) error {

	return nil
}

func (m *mockTaskRepo) Delete(
	id int,
) error {

	return nil
}

func TestTaskCreateSuccess(t *testing.T) {

	repo := &mockTaskRepo{}

	service := NewTaskService(
		repo,
		nil,
	)

	_, err := service.Create(
		entity.Task{
			ProjectID:  1,
			Title:      "Learn Go",
			Status:     "TODO",
			AssigneeID: 1,
		},
	)

	if err != nil {
		t.Errorf("expected success")
	}
}

func TestTaskCreateEmptyTitle(t *testing.T) {

	repo := &mockTaskRepo{}

	service := NewTaskService(
		repo,
		nil,
	)

	_, err := service.Create(
		entity.Task{
			ProjectID:  1,
			Title:      "",
			Status:     "TODO",
			AssigneeID: 1,
		},
	)

	if err == nil {
		t.Errorf("expected error")
	}
}

func TestTaskCreateInvalidStatus(t *testing.T) {

	repo := &mockTaskRepo{}

	service := NewTaskService(
		repo,
		nil,
	)

	_, err := service.Create(
		entity.Task{
			ProjectID:  1,
			Title:      "Test Task",
			Status:     "INVALID",
			AssigneeID: 1,
		},
	)

	if err == nil {
		t.Errorf("expected error")
	}
}

func TestTaskGetAll(t *testing.T) {

	repo := &mockTaskRepo{}

	service := NewTaskService(
		repo,
		nil,
	)

	tasks, err := service.GetAll()

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if len(tasks) != 1 {
		t.Errorf("expected 1 task, got %d", len(tasks))
	}
}

func TestTaskGetByID(t *testing.T) {

	repo := &mockTaskRepo{}

	service := NewTaskService(
		repo,
		nil,
	)

	task, err := service.GetByID(1)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if task.ID != 1 {
		t.Errorf("expected ID 1, got %d", task.ID)
	}
}

func TestTaskUpdateSuccess(t *testing.T) {

	repo := &mockTaskRepo{}

	service := NewTaskService(
		repo,
		nil,
	)

	err := service.Update(
		1,
		entity.Task{
			Title:  "Updated Task",
			Status: "IN_PROGRESS",
		},
	)

	if err != nil {
		t.Errorf("expected success")
	}
}

func TestTaskDeleteSuccess(t *testing.T) {

	repo := &mockTaskRepo{}

	service := NewTaskService(
		repo,
		nil,
	)

	err := service.Delete(1)

	if err != nil {
		t.Errorf("expected success")
	}
}

func TestTaskUpdateEmptyTitle(t *testing.T) {

	repo := &mockTaskRepo{}

	service := NewTaskService(
		repo,
		nil,
	)

	err := service.Update(
		1,
		entity.Task{
			Title:  "",
			Status: "TODO",
		},
	)

	if err == nil {
		t.Errorf("expected error")
	}
}
