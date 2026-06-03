package repository

import(
	"database/sql"
	"errors"
	"ce191383/task_management/internal/entity"
)

type UserRepository interface{
	Register(user entity.User)(entity.User, error)
	GetAll()([]entity.User,error)
	GetByID(id int)(*entity.User, error)
	GetByEmail(email string) (*entity.User,error)
}

type userRepository struct{
	db*sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r*userRepository) Register(u entity.User) (entity.User,error){
	query :="INSERT INTO users (email, password_hash, full_name) VALUES ($1, $2, $3) RETURNING id"

	err:=r.db.QueryRow(query, u.Email, u.PasswordHash, u.FullName).Scan(&u.ID)
	return u, err
}


func (r *userRepository) GetAll()([]entity.User,error){
	query := "SELECT id, email, password_hash, full_name FROM users"

	rows, err := r.db.Query(query)
	if err != nil{
		return nil, err
	}
	defer rows.Close()

	var users []entity.User

	for rows.Next() {
		var u entity.User

		err := rows.Scan(
			&u.ID,
			&u.Email,
			&u.PasswordHash,
			&u.FullName,
		)

		if err != nil{
			return nil, err

		}

		users = append(users, u)
	}

	return users, nil
}


func (r *userRepository) GetByID(id int) (*entity.User,error){
	query := "SELECT id , email, password_hash, full_name FROM users WHERE id=$1"

	var u entity.User
	err:=r.db.QueryRow(query, id).Scan(
		&u.ID,
		&u.Email,
		&u.PasswordHash,
		&u.FullName,
	)

	if err != nil{
		return nil, errors.New("user not found")
	}
	return &u, nil
}


func (r *userRepository) GetByEmail(email string) (*entity.User,error){
	query := "SELECT id, email, password_hash, full_name FROM users WHERE email = $1"

	var u entity.User

	err := r.db.QueryRow(query, email).Scan(
		&u.ID,
		&u.Email,
		&u.PasswordHash,
		&u.FullName,
	)

	if err != nil{
		return nil, errors.New("user not found")
	}
	return &u, nil

}


