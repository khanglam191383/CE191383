package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"ce191383/task_management/internal/entity"
	"ce191383/task_management/internal/service"
)

type ProjectHandler struct {
	service service.ProjectService
}

func NewProjectHandler(s service.ProjectService) *ProjectHandler {
	return &ProjectHandler{service: s}
}

func (h *ProjectHandler) CreateProject(w http.ResponseWriter, r *http.Request) {
	var p entity.Project

	json.NewDecoder(r.Body).Decode(&p)

	userID:= r.Context().Value("user_id").(int)

	p.OwnerID = userID

	result, err := h.service.Create(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	encoder.Encode(result)
}

func (h *ProjectHandler) GetAllProjects(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("user_id").(int)
	projects, _ := h.service.GetAll(userID)
	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	encoder.Encode(projects)
}

func (h *ProjectHandler) GetProjectByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Path[len("/projects/"):])

	project, err := h.service.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	encoder.Encode(project)
}

func (h *ProjectHandler) UpdateProject(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Path[len("/projects/"):])

	var p entity.Project
	json.NewDecoder(r.Body).Decode(&p)

	userID := r.Context().Value("user_id").(int)
	err := h.service.Update(id, userID, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte("updated"))
}

func (h *ProjectHandler) DeleteProject(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Path[len("/projects/"):])

	userID := r.Context().Value("user_id").(int)

	err := h.service.Delete(id, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte("deleted"))
}
