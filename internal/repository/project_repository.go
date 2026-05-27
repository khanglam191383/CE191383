package repository

import (
	"database/sql"
	"errors"

	"ce191383/task_management/internal/entity"
)

type ProjectRepository interface {
	Create(project entity.Project) (entity.Project, error)
	GetAll() ([]entity.Project, error)
	GetByID(id int) (*entity.Project, error)
	Update(id int, project entity.Project) error
	Delete(id int) error
	GetOwnerID(projectID int) (int, error)
	GetAllByOwner(userID int) ([] entity.Project, error)
}

type projectRepository struct {
	db *sql.DB
}

func NewProjectRepository(db *sql.DB) ProjectRepository {
	return &projectRepository{db: db}
}

func (r *projectRepository) Create(p entity.Project) (entity.Project, error) {
	query := `INSERT INTO projects (name, description, owner_id)
	          VALUES ($1, $2, $3) RETURNING id`

	err := r.db.QueryRow(query, p.Name, p.Description, p.OwnerID).Scan(&p.ID)
	return p, err
}

func (r *projectRepository) GetAll() ([]entity.Project, error) {
	query := `SELECT id, name, description, owner_id FROM projects`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []entity.Project

	for rows.Next() {
		var p entity.Project
		rows.Scan(&p.ID, &p.Name, &p.Description, &p.OwnerID)
		projects = append(projects, p)
	}

	return projects, nil
}

func (r *projectRepository) GetByID(id int) (*entity.Project, error) {
	query := `SELECT id, name, description, owner_id FROM projects WHERE id=$1`

	var p entity.Project
	err := r.db.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Description, &p.OwnerID)

	if err != nil {
		return nil, errors.New("project not found")
	}

	return &p, nil
}

func (r *projectRepository) Update(id int, p entity.Project) error {
	query := `UPDATE projects SET name=$1, description=$2, owner_id=$3 WHERE id=$4`

	res, err := r.db.Exec(query, p.Name, p.Description, p.OwnerID, id)
	if err != nil {
		return err
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return errors.New("project not found")
	}

	return nil
}

func (r *projectRepository) Delete(id int) error {
	query := `DELETE FROM projects WHERE id=$1`

	res, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return errors.New("project not found")
	}

	return nil
}

func (r *projectRepository) GetOwnerID(projectID int) (int, error){
	query := "SELECT owner_id FROM projects WHERE id = $1"

	var ownerID int

	err := r.db.QueryRow(query, projectID).Scan(&ownerID)

	if err != nil{
		return 0, errors.New("project not found")
	}

	return ownerID, nil
}

func (r* projectRepository) GetAllByOwner(userID int) ([]entity.Project, error){
	query := `SELECT id, name, description, owner_id FROM projects WHERE owner_id = $1`

	rows, err := r.db.Query(query, userID)

	if err != nil{
		return nil, err
	}

	defer rows.Close()

	var projects []entity.Project

	for rows.Next(){

		var p entity.Project

		rows.Scan(
			&p.ID,
			&p.Name,
			&p.Description,
			&p.OwnerID,
		)

		projects = append(projects, p)
	}

	return projects, nil

}