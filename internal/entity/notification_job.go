package entity

type NotificationJob struct{
	TaskID int `json:"task_id"`
	UserID int `json:"user_id"`
	Message string `json:"message"`
	Retry int `json:"retry"`
}