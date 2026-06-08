package service

import (
	"fmt"
	"testing"
	"time"

	"ce191383/task_management/internal/entity"
	"ce191383/task_management/internal/repository"
)

func TestProjectUpdatePermission(t *testing.T) {

	db := setupTestDB()
	defer db.Close()

	projectRepo := repository.NewProjectRepository(db)

	projectService := NewProjectService(projectRepo)

	// User A
	userRepo := repository.NewUserRepository(db)

	userA, err := userRepo.Register(entity.User{
		Email: fmt.Sprintf(
			"userA_%d@gmail.com",
			time.Now().UnixNano(),
		),
		PasswordHash: "hash",
		FullName:     "User A",
	})

	if err != nil {
		t.Fatal(err)
	}

	// User B
	userB, err := userRepo.Register(entity.User{
		Email: fmt.Sprintf(
			"userB_%d@gmail.com",
			time.Now().UnixNano(),
		),
		PasswordHash: "hash",
		FullName:     "User B",
	})

	if err != nil {
		t.Fatal(err)
	}

	// User A tạo project
	project, err := projectService.Create(entity.Project{
		Name:        "Secret Project",
		Description: "Demo",
		OwnerID:     userA.ID,
	})

	if err != nil {
		t.Fatal(err)
	}

	// User B sửa project của User A
	err = projectService.Update(
		project.ID,
		userB.ID,
		entity.Project{
			Name:        "Hacked",
			Description: "Hacked",
		},
	)

	if err == nil {
		t.Fatal("expected forbidden error")
	}

	if err.Error() != "forbidden" {
		t.Fatalf(
			"expected forbidden got %v",
			err,
		)
	}
}