package service

import (
	"errors"

	"ce191383/task_management/internal/entity"
	"ce191383/task_management/internal/repository"
)

type ProjectService interface {
	Create(p entity.Project) (entity.Project, error)
	GetAll(userID int) ([]entity.Project, error)
	GetByID(id int) (*entity.Project, error)
	Update(id int, userID int, p entity.Project) error
	Delete(id int, userID int) error
}

type projectService struct {
	repo repository.ProjectRepository
}

func NewProjectService(repo repository.ProjectRepository) ProjectService {
	return &projectService{repo: repo}
}

func (s *projectService) Create(p entity.Project) (entity.Project, error) {
	if p.Name == "" {
		return entity.Project{}, errors.New("name is required")
	}

	return s.repo.Create(p)
}

func (s *projectService) GetAll(userID int) ([]entity.Project, error) {
	return s.repo.GetAllByOwner(userID)
}

func (s *projectService) GetByID(id int) (*entity.Project, error) {
	return s.repo.GetByID(id)
}

func (s *projectService) Update(id int, userID int, p entity.Project) error {

	if p.Name == "" {
		return errors.New("name is required")
	}

	ownerID, err := s.repo.GetOwnerID(id)
	if err != nil {
		return err
	}

	if ownerID != userID {
		return errors.New("forbidden")
	}

	return s.repo.Update(id, p)
}

func (s *projectService) Delete(id int, userID int) error {

	ownerID, err := s.repo.GetOwnerID(id)
	if err != nil {
		return err
	}

	if ownerID != userID {
		return errors.New("forbidden")
	}

	return s.repo.Delete(id)
}
