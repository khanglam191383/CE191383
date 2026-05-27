package service

import (
	"errors"

	"ce191383/task_management/internal/entity"
	"ce191383/task_management/internal/repository"

	"ce191383/task_management/internal/websocket"
	"encoding/json"
)

type CommentService interface {
	Create(comment entity.Comment) (entity.Comment, error)
	GetByTaskID(taskID int) ([]entity.Comment, error)
}

type commentService struct {
	repo repository.CommentRepository
}

func NewCommentService(
	repo repository.CommentRepository,
) CommentService {
	return &commentService{
		repo: repo,
	}
}

func (s *commentService) Create(
	comment entity.Comment,
) (entity.Comment, error) {

	if comment.Content == "" {
		return entity.Comment{}, errors.New("content is required")
	}

	createdComment, err := s.repo.Create(comment)

	if err != nil {
		return entity.Comment{}, err
	}

	event := map[string]interface{}{
		"even":    "comment_created",
		"task_id": createdComment.TaskID,
		"content": createdComment.Content,
	}

	data, _ := json.MarshalIndent(event, "", "  ")

	websocket.Broadcast(data)

	return s.repo.Create(comment)
}

func (s *commentService) GetByTaskID(taskID int) ([]entity.Comment, error) {
	return s.repo.GetByTaskID(taskID)
}
