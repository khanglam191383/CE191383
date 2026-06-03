package service

import (
	"testing"

	"ce191383/task_management/internal/entity"
)

type mockProjectRepo struct{}

func (m *mockProjectRepo) Create(
	p entity.Project,
) (entity.Project, error) {

	p.ID = 1

	return p, nil
}

func (m *mockProjectRepo) GetAll() ([]entity.Project, error) {

	return []entity.Project{}, nil
}

func (m *mockProjectRepo) GetAllByOwner(
	userID int,
) ([]entity.Project, error) {

	return []entity.Project{}, nil
}

func (m *mockProjectRepo) GetByID(
	id int,
) (*entity.Project, error) {

	return &entity.Project{
		ID: id,
	}, nil
}

func (m *mockProjectRepo) GetOwnerID(
	id int,
) (int, error) {

	return 1, nil
}

func (m *mockProjectRepo) Update(
	id int,
	p entity.Project,
) error {

	return nil
}

func (m *mockProjectRepo) Delete(
	id int,
) error {

	return nil
}

func TestCreateProjectSuccess(t *testing.T) {

	repo := &mockProjectRepo{}

	service := NewProjectService(repo)

	_, err := service.Create(
		entity.Project{
			Name: "Demo Project",
		},
	)

	if err != nil {
		t.Errorf("expected success")
	}
}

func TestCreateProjectEmptyName(t *testing.T) {

	repo := &mockProjectRepo{}

	service := NewProjectService(repo)

	_, err := service.Create(
		entity.Project{},
	)

	if err == nil {
		t.Errorf("expected error")
	}
}
