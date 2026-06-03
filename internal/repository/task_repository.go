package repository

import (
	"errors"
	"database/sql"
	"ce191383/task_management/internal/entity"
)

type TaskRepository interface {
	Create(task entity.Task) (entity.Task, error)
	GetAll() ([]entity.Task, error)
	GetByID(id int) (*entity.Task, error)
	Update(id int, updatedTask entity.Task) error
	Delete(id int) error
}

type taskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) TaskRepository {
	return &taskRepository{
		db: db,
	}
}

func (r *taskRepository) Create(task entity.Task) (entity.Task, error) {

	query := "INSERT INTO tasks (project_id, title, description, status, assignee_id) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	
	err := r.db.QueryRow(query, task.ProjectID, task.Title, task.Description, task.Status, task.AssigneeID).Scan(&task.ID)
	if err !=nil {
		return entity.Task{}, err
	}

	return task, nil
}

func (r *taskRepository) GetAll() ([]entity.Task, error) {
	query := "SELECT id, project_id, title, description, status, assignee_id FROM tasks"

	rows, err := r.db.Query(query)
	if err != nil {
		return nil , err
	}
	defer rows.Close()

	var tasks []entity.Task

	for rows.Next(){
		var task entity.Task
		err:=rows.Scan(&task.ID, &task.ProjectID, &task.Title, &task.Description, &task.Status, &task.AssigneeID)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (r *taskRepository) GetByID(id int) (*entity.Task, error) {

	query := "SELECT id, project_id, title, description, status, assignee_id FROM tasks WHERE id =$1"
	var task entity.Task

	err := r.db.QueryRow(query, id).Scan(
		&task.ID,
		&task.ProjectID,
		&task.Title,
		&task.Description,
		&task.Status,
		&task.AssigneeID,
	)

	if err != nil {
		return nil, errors.New("task not found")
	}

	return &task, nil
}

func (r *taskRepository) Update(id int, updatedTask entity.Task) error {

	query := "UPDATE tasks SET project_id = $1, title = $2, description = $3, status = $4, assignee_id = $5 WHERE id=$6"

	result, err := r.db.Exec(
		query,
		updatedTask.ProjectID,
		updatedTask.Title,
		updatedTask.Description,
		updatedTask.Status,
		updatedTask.AssigneeID,
		id,
	)

	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("task not found")
	}

	return nil
}

func (r *taskRepository) Delete(id int) error {

	query :="DELETE FROM tasks WHERE id = $1"

	result, err := r.db.Exec(query, id)
	if err!= nil{
		return err
	}

	rows,_ :=result.RowsAffected()
	if rows == 0{
		return errors.New("task not found")
	}

	return nil
}