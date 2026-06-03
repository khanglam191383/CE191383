package service

import (
	"ce191383/task_management/internal/entity"
	"testing"
)

type mockUserRepo struct{}

func (m *mockUserRepo) Create(
	user entity.User,
) (entity.User, error) {

	user.ID = 1
	return user, nil
}

func (m *mockUserRepo) GetAll() (
	[]entity.User,
	error,
) {
	return []entity.User{}, nil
}

func (m *mockUserRepo) GetByID(
	id int,
) (*entity.User, error) {
	return &entity.User{}, nil
}

func (m *mockUserRepo) GetByEmail(
	email string,
) (*entity.User, error) {
	return &entity.User{}, nil
}


func (m *mockUserRepo) Register(
	user entity.User,
) (entity.User, error) {

	user.ID = 1
	return user, nil
}


func TestRegisterSuccess(
	t *testing.T,
) {

	repo := &mockUserRepo{}

	service := NewUserService(repo)

	user := entity.User{
		Email: "test@gmail.com",
		Password: "123456",
		FullName: "Test User",
	}

	result, err := service.Register(user)

	if err != nil {
		t.Errorf(
			"expected no error, got %v",
			err,
		)
	}

	if result.ID != 1 {
		t.Errorf(
			"expected ID 1, got %d",
			result.ID,
		)
	}
}

func TestRegisterMissingEmail(t *testing.T) {

	repo := &mockUserRepo{}

	service := NewUserService(repo)

	_, err := service.Register(
		entity.User{
			Password: "123456",
			FullName: "Khang",
		},
	)

	if err == nil {
		t.Errorf("expected error")
	}
}

func TestRegisterMissingPassword(t *testing.T) {

	repo := &mockUserRepo{}

	service := NewUserService(repo)

	_, err := service.Register(
		entity.User{
			Email: "test@gmail.com",
			FullName: "Khang",
		},
	)

	if err == nil {
		t.Errorf("expected error")
	}
}