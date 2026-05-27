package handler

import(
	"encoding/json"
	"net/http"
	"strconv"

	"ce191383/task_management/internal/entity"
	"ce191383/task_management/internal/service"
)

type CommentHandler struct{
	service service.CommentService
}

func NewCommentHandler(
	s service.CommentService,
) *CommentHandler {

	return &CommentHandler{
		service: s,
	}
}

func (h *CommentHandler) CreateComment(
	w http.ResponseWriter,
	r *http.Request,
){

	var comment entity.Comment

	json.NewDecoder(r.Body).Decode(&comment)

	userID := r.Context().Value("user_id").(int)

	comment.UserID = userID

	result, err := h.service.Create(comment)

	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(result)
}

func (h*CommentHandler) GetCommentsByTaskID(
	w http.ResponseWriter,
	r *http.Request,
){
	taskIDStr := r.URL.Query().Get("task_id")

	taskID, err := strconv.Atoi(taskIDStr)

	if err != nil {
		http.Error(w, "invalid task_id", http.StatusBadRequest)
		return
	}

	comments, err := h.service.GetByTaskID(taskID)

	if err != nil {
		http .Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(comments)
}