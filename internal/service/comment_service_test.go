package service

import (
	"testing"

	"ce191383/task_management/internal/entity"
)

type mockCommentRepo struct{}

func (m *mockCommentRepo) Create(
	c entity.Comment,
) (entity.Comment, error) {

	c.ID = 1

	return c, nil
}

func (m *mockCommentRepo) GetByTaskID(
	taskID int,
) ([]entity.Comment, error) {

	return []entity.Comment{
		{
			ID: 1,
			TaskID: taskID,
			Content: "hello",
		},
	}, nil
}

func TestCreateCommentSuccess(t *testing.T) {

	repo := &mockCommentRepo{}

	service := NewCommentService(repo)

	_, err := service.Create(
		entity.Comment{
			TaskID: 1,
			UserID: 1,
			Content: "Test",
		},
	)

	if err != nil {
		t.Errorf("expected success")
	}
}

func TestCreateCommentEmptyContent(t *testing.T) {

	repo := &mockCommentRepo{}

	service := NewCommentService(repo)

	_, err := service.Create(
		entity.Comment{},
	)

	if err == nil {
		t.Errorf("expected error")
	}
}

func TestGetCommentsByTaskID(t *testing.T) {

	repo := &mockCommentRepo{}

	service := NewCommentService(repo)

	comments, err := service.GetByTaskID(1)

	if err != nil {
		t.Errorf("unexpected error")
	}

	if len(comments) != 1 {
		t.Errorf("expected 1 comment")
	}
}