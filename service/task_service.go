package service

import (
	"errors"
	"ce191383/task_management/entity"
	"ce191383/task_management/repository"
)

type TaskService interface {
	Create(task entity.Task) (entity.Task, error)
	GetAll() []entity.Task
	GetByID(id int) (*entity.Task, error)
	Update(id int, task entity.Task) error
	Delete(id int) error
}

type taskService struct {
	repo repository.TaskRepository
}

func NewTaskService(repo repository.TaskRepository) TaskService {
	return &taskService{
		repo: repo,
	}
}

func (s *taskService) Create(task entity.Task) (entity.Task, error) {

	if task.Title == "" {
		return entity.Task{}, errors.New("title is required")
	}

	validStatus := map[string]bool{
		"TODO":        true,
		"IN_PROGRESS": true,
		"DONE":        true,
	}

	if !validStatus[task.Status] {
		return entity.Task{}, errors.New("invalid status")
	}

	return s.repo.Create(task), nil
}

func (s *taskService) GetAll() []entity.Task {
	return s.repo.GetAll()
}

func (s *taskService) GetByID(id int) (*entity.Task, error) {
	return s.repo.GetByID(id)
}

func (s *taskService) Update(id int, task entity.Task) error {

	if task.Title == "" {
		return errors.New("title is required")
	}

	return s.repo.Update(id, task)
}

func (s *taskService) Delete(id int) error {
	return s.repo.Delete(id)
}