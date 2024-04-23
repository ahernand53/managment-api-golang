package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type TaskService struct {
	store Store
}

func NewTasksService(s Store) *TaskService {
	return &TaskService{store: s}
}

func (s *TaskService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/tasks", WithJWTAuth(s.handleCreateTask, s.store)).Methods("POST")
	r.HandleFunc("/tasks/{id}", WithJWTAuth(s.handleGetTask, s.store)).Methods("GET")
}

func (s *TaskService) handleCreateTask(w http.ResponseWriter, r *http.Request) {
	// create task
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, ErrorResponse{Error: "Error reading request body"})
		return
	}

	defer r.Body.Close()

	var task *Task
	err = json.Unmarshal(body, &task)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid request payload"})
		return
	}

	if err := validateTaskPayload(task); err != nil {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	t, err := s.store.CreateTask(task)

	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, ErrorResponse{Error: "Error creating task"})
		return
	}

	WriteJSON(w, http.StatusCreated, t)

}

func (s *TaskService) handleGetTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if id == "" {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid task id"})
		return
	}

	t, err := s.store.GetTask(id)

	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, ErrorResponse{Error: "task not found"})
		return
	}

	WriteJSON(w, http.StatusOK, t)
}

func validateTaskPayload(task *Task) error {
	if task.Name == "" {
		return errors.New("task name is required")
	}

	if task.ProjectID == 0 {
		return errors.New("project id is required")
	}

	if task.AssignedToId == 0 {
		return errors.New("assigned to id is required")
	}

	return nil
}
