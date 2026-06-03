package service

import (
	"errors"
	"ce191383/task_management/internal/entity"
    "ce191383/task_management/internal/repository"
	"ce191383/task_management/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	GetAll() ([]entity.User, error)
	GetByID(id int) (*entity.User, error)
	Login(email, password string) (string, error)
	Register(user entity.User) (entity.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService{
	return &userService{repo: repo}
}



func (s *userService) GetAll() ([]entity.User,error){
	return s.repo.GetAll()
}

func (s *userService) GetByID(id int) (*entity.User, error){
	return s.repo.GetByID(id)
}


func (s *userService) Login(email, password string) (string, error){
	user, err := s.repo.GetByEmail(email)
	if err != nil{
		return "", errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte(password),
	)

	if err != nil{
		return "", errors.New("invalid email or password")
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil{
		return "", err
	}

	return token, nil
}

func (s *userService) Register(
	user entity.User,
) (entity.User, error) {

	if user.Email == "" {
		return entity.User{}, errors.New("email is required")
	}

	if user.Password == "" {
		return entity.User{}, errors.New("password is required")
	}

	if user.FullName == "" {
		return entity.User{}, errors.New("full_name is required")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return entity.User{}, err
	}

	user.PasswordHash = string(hashedPassword)

	return s.repo.Register(user)
}



