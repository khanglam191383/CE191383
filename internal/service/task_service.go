package service

import (
	"ce191383/task_management/internal/entity"
	"ce191383/task_management/internal/repository"
	"encoding/json"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"

	"ce191383/task_management/internal/config"
	"ce191383/task_management/internal/websocket"
	"context"
	"fmt"
)

type TaskService interface {
	Create(task entity.Task) (entity.Task, error)
	GetAll() ([]entity.Task, error)
	GetByID(id int) (*entity.Task, error)
	Update(id int, task entity.Task) error
	Delete(id int) error
}

type taskService struct {
	repo  repository.TaskRepository
	redis *redis.Client
}

func NewTaskService(repo repository.TaskRepository, redis *redis.Client) TaskService {
	return &taskService{
		repo:  repo,
		redis: redis,
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

	created, err := s.repo.Create(task)
	if err != nil {
		return entity.Task{}, err
	}

	job := entity.NotificationJob{
		TaskID:  created.ID,
		UserID:  1,
		Message: "Task created successfully",
		Retry:   0,
	}

	err = s.PushNotification(job)
	if err != nil {
		fmt.Println("push job error:", err)
	}

	if s.redis != nil {
		s.redis.Del(context.Background(), "tasks")
	}

	return created, nil
}

func (s *taskService) GetAll() ([]entity.Task, error) {
	if s.redis == nil {
		return s.repo.GetAll()
	}

	ctx := context.Background()
	cacheKey := "tasks"

	cacheTasks, err := s.redis.Get(ctx, cacheKey).Result()

	if err == nil {
		var tasks []entity.Task

		if err := json.Unmarshal([]byte(cacheTasks), &tasks); err != nil {
			fmt.Println("CACHE UNMARSHAL ERROR:", err)
		} else {
			fmt.Println("CACHE HIT")

			return tasks, nil
		}

	}

	fmt.Println("CACHE MISS")

	tasks, err := s.repo.GetAll()

	if err != nil {
		return nil, err
	}

	jsonData, err := json.Marshal(tasks)

	if err != nil {
		return nil, err
	}
	err = s.redis.Set(
		config.Ctx,
		cacheKey,
		jsonData,
		1*time.Minute,
	).Err()
	if err != nil {
		fmt.Println("REDIS SET ERROR:", err)
	}

	return tasks, nil
}

func (s *taskService) GetByID(id int) (*entity.Task, error) {
	if s.redis == nil {
		return s.repo.GetByID(id)
	}

	ctx := context.Background()
	cacheKey := fmt.Sprintf("task:%d", id)

	cached, err := s.redis.Get(ctx, cacheKey).Result()
	if err == nil {
		var task entity.Task

		if err := json.Unmarshal([]byte(cached), &task); err != nil {
			fmt.Println("CACHE UNMARSHAL ERROR:", err)
		} else {
			return &task, nil
		}

	}

	task, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(task)

	if err != nil {
		return nil, err
	}
	err = s.redis.Set(
		ctx,
		cacheKey,
		data,
		10*time.Minute).Err()
	if err != nil {
		fmt.Println("REDIS SET ERROR:", err)
	}

	return task, nil
}

func (s *taskService) Update(id int, task entity.Task) error {

	if task.Title == "" {
		return errors.New("title is required")
	}

	err := s.repo.Update(id, task)
	if err != nil {
		return err
	}
	s.clearTaskCache(id)

	event := map[string]interface{}{
		"event":   "task_updated",
		"task_id": id,
	}

	data, _ := json.MarshalIndent(event, "", "  ")

	websocket.Broadcast(data)

	return nil
}

func (s *taskService) Delete(id int) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}

	s.clearTaskCache(id)

	return nil
}

func (s *taskService) clearTaskCache(id int) {
	if s.redis == nil {
		return
	}

	ctx := context.Background()

	s.redis.Del(ctx, "tasks")
	s.redis.Del(ctx, fmt.Sprintf("task:%d", id))
}

func (s *taskService) PushNotification(job entity.NotificationJob) error {
	if s.redis == nil {
		return nil
	}

	ctx := context.Background()

	data, err := json.Marshal(job)
	if err != nil {
		return err
	}

	return s.redis.LPush(ctx, "notification_queue", data).Err()
}
