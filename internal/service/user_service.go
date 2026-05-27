package service

import
(
	"errors"
	"ce191383/task_management/internal/entity"
    "ce191383/task_management/internal/repository"
	"ce191383/task_management/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface{
	Create(user entity.User)(entity.User,error)
	GetAll()([]entity.User,error)
	GetByID(id int)(*entity.User,error)
	Login(email, password string) (string, error)
}

type userService struct{
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService{
	return &userService{repo: repo}
}

func (s *userService)Create(user entity.User)(entity.User,error){
	if user.Email==""{
		return entity.User{}, errors.New("email is required")
	}
	if user.PasswordHash==""{
		return entity.User{},errors.New("password is required")
	}

	if user.FullName == ""{
		return entity.User{}, errors.New("full_name is required")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), 10)
	if err != nil{
		return entity.User{}, err
	}

	user.PasswordHash = string(hashed)
	return s.repo.Create(user)
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



