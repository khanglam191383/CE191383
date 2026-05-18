package repository

import (
	"errors"
	"ce191383/task_management/entity"
)

type TaskRepository interface {
	Create(task entity.Task) entity.Task
	GetAll() []entity.Task
	GetByID(id int) (*entity.Task, error)
	Update(id int, updatedTask entity.Task) error
	Delete(id int) error
}

type taskRepository struct {
	tasks []entity.Task
}

func NewTaskRepository() TaskRepository {
	return &taskRepository{
		tasks: []entity.Task{},
	}
}

func (r *taskRepository) Create(task entity.Task) entity.Task {

	r.tasks = append(r.tasks, task)

	return task
}

func (r *taskRepository) GetAll() []entity.Task {
	return r.tasks
}

func (r *taskRepository) GetByID(id int) (*entity.Task, error) {

	for _, task := range r.tasks {

		if task.ID == id {
			return &task, nil
		}
	}

	return nil, errors.New("task not found")
}

func (r *taskRepository) Update(id int, updatedTask entity.Task) error {

	for i, task := range r.tasks {

		if task.ID == id {

			r.tasks[i] = updatedTask

			return nil
		}
	}

	return errors.New("task not found")
}

func (r *taskRepository) Delete(id int) error {

	for i, task := range r.tasks {

		if task.ID == id {

			r.tasks = append(r.tasks[:i], r.tasks[i+1:]...)

			return nil
		}
	}

	return errors.New("task not found")
}