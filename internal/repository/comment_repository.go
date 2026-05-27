package repository

import (
	"database/sql"

	"ce191383/task_management/internal/entity"
)

type CommentRepository interface{
	Create(comment entity.Comment) (entity.Comment, error)
	GetByTaskID(taskID int) ([]entity.Comment, error)
}

type commentRepository struct{
	db *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository{
	return &commentRepository{
		db :db,
	}
}

func (r *commentRepository) Create(comment entity.Comment) (entity.Comment, error){
	query:= "INSERT INTO comments (task_id, user_id, content) VALUES ($1, $2, $3) RETURNING id, created_at"


	err := r.db.QueryRow(
		query,
		comment.TaskID,
		comment.UserID,
		comment.Content,
	).Scan(
		&comment.ID,
		&comment.CreatedAt,
	)

	if err != nil{
		return entity.Comment{}, err
	}

	return comment, nil
}

func (r *commentRepository) GetByTaskID(taskID int) ([]entity.Comment, error){

	query := "SELECT id, task_id, user_id, content, created_at FROM comments WHERE task_id = $1 ORDER BY created_at ASC"

	rows, err := r.db.Query(query, taskID)

	if err != nil {

		return nil, err
	}
		defer rows.Close()

		var comments []entity.Comment

		for rows.Next() {
			var comment entity.Comment

			err := rows.Scan(
				&comment.ID,
				&comment.TaskID,
				&comment.UserID,
				&comment.Content,
				&comment.CreatedAt,
			)

			if err != nil{
				return nil, err
			}

			comments = append(comments, comment)
		}

		return comments, nil
	}

