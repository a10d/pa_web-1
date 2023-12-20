package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type todoTypeInput struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	Color        string `json:"color"`
	ReminderTime int    `json:"reminderTime"`
}

func (s *APIServer) handleGetTodoTypes(w http.ResponseWriter, r *http.Request) error {

	todoTypes, err := s.store.GetTodoTypes()

	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, todoTypes)
}

func (s *APIServer) handleCreateTodoType(w http.ResponseWriter, r *http.Request) error {

	input := new(todoTypeInput)

	if err := json.NewDecoder(r.Body).Decode(input); err != nil {
		return err
	}

	// TODO: Handle validation

	todoType, err := s.store.CreateTodoType(NewTodoType(
		input.Name,
		input.Description,
		input.Color,
		input.ReminderTime,
	))

	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, todoType)
}

func (s *APIServer) handleGetTodoTypeById(w http.ResponseWriter, r *http.Request) error {

	todoType, err := s.store.GetTodoById(mux.Vars(r)["id"])

	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, todoType)
}

func (s *APIServer) handleUpdateTodoType(w http.ResponseWriter, r *http.Request) error {

	todoType, err := s.store.GetTodoTypeById(mux.Vars(r)["id"])

	if err != nil {
		return err
	}

	input := new(todoTypeInput)

	if err := json.NewDecoder(r.Body).Decode(input); err != nil {
		return err
	}

	// TODO: Handle validation

	todoType.Name = input.Name
	todoType.Description = input.Description
	todoType.Color = input.Color
	todoType.ReminderTime = input.ReminderTime

	updatedTodoType, err := s.store.UpdateTodoType(todoType)

	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, updatedTodoType)
}

func (s *APIServer) handleDeleteTodoType(w http.ResponseWriter, r *http.Request) error {

	todoType, err := s.store.GetTodoTypeById(mux.Vars(r)["id"])

	if err != nil {
		return err
	}

	err = s.store.DeleteTodoType(todoType)

	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}
