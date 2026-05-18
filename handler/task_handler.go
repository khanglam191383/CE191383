package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"ce191383/task_management/utils"

	"ce191383/task_management/entity"
	"ce191383/task_management/service"
)

type TaskHandler struct {
	service service.TaskService
}

func NewTaskHandler(service service.TaskService) *TaskHandler {
	return &TaskHandler{
		service: service,
	}
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {

	var task entity.Task

	err := json.NewDecoder(r.Body).Decode(&task)

	if err != nil {

		utils.WriteError(w, http.StatusBadRequest, err.Error())

		return
	}

	result, err := h.service.Create(task)

	if err != nil {

		utils.WriteError(w, http.StatusBadRequest, err.Error())

		return
	}

	utils.WriteSuccess(w, result)
}

func (h *TaskHandler) GetAllTasks(w http.ResponseWriter, r *http.Request) {

	tasks := h.service.GetAll()

	json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandler) GetTaskByID(w http.ResponseWriter, r *http.Request) {

	idParam := strings.TrimPrefix(r.URL.Path, "/tasks/")

	id, err := strconv.Atoi(idParam)

	if err != nil {

		utils.WriteError(w, http.StatusBadRequest, "invalid id")

		return
	}

	task, err := h.service.GetByID(id)

	if err != nil {

		http.Error(w, err.Error(), http.StatusNotFound)

		return
	}

	json.NewEncoder(w).Encode(task)
}

func (h *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {

	idParam := strings.TrimPrefix(r.URL.Path, "/tasks/")

	id, err := strconv.Atoi(idParam)

	if err != nil {

		utils.WriteError(w, http.StatusBadRequest, "invalid id")

		return
	}

	var task entity.Task

	err = json.NewDecoder(r.Body).Decode(&task)

	if err != nil {

		utils.WriteError(w, http.StatusBadRequest, err.Error())

		return
	}

	err = h.service.Update(id, task)

	if err != nil {

		utils.WriteError(w, http.StatusBadRequest, err.Error())

		return
	}

	w.Write([]byte("updated successfully"))
}

func (h *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {

	idParam := strings.TrimPrefix(r.URL.Path, "/tasks/")

	id, err := strconv.Atoi(idParam)

	if err != nil {

		utils.WriteError(w, http.StatusBadRequest, "invalid id")

		return
	}

	err = h.service.Delete(id)

	if err != nil {

		utils.WriteError(w, http.StatusBadRequest, err.Error())

		return
	}

	w.Write([]byte("deleted successfully"))
}