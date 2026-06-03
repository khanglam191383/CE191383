package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"ce191383/task_management/internal/dto"
	"ce191383/task_management/internal/entity"
	"ce191383/task_management/internal/service"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(s service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	var response []dto.UserResponse

	for _, user := range users {
		response = append(response, dto.UserResponse{
			ID:       user.ID,
			Email:    user.Email,
			FullName: user.FullName,
		})
	}

	encoder.Encode(response)
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Path[len("/users/"):])

	user, err := h.service.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	response := dto.UserResponse{
		ID:       user.ID,
		Email:    user.Email,
		FullName: user.FullName,
	}

	encoder.Encode(response)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginRequest

	json.NewDecoder(r.Body).Decode(&req)

	token, err := h.service.Login(req.Email, req.Password)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")

	encoder.Encode(map[string]string{
		"token": token,
	})
}

func (h *UserHandler) Register(
	w http.ResponseWriter,
	r *http.Request,
) {

	var user entity.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	registeredUser, err := h.service.Register(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := dto.UserResponse{
		ID:       registeredUser.ID,
		Email:    registeredUser.Email,
		FullName: registeredUser.FullName,
	}

	json.NewEncoder(w).Encode(response)

}
