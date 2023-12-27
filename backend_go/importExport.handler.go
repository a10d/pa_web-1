package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type todoTypeTransfer struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Color        string `json:"color"`
	ReminderTime int    `json:"reminderTime"`
}

type userTransfer struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type todoTransfer struct {
	ID          string   `json:"id"`
	TodoTypeID  string   `json:"type"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	DueDate     string   `json:"dueDate"`
	Completed   bool     `json:"completed"`
	AssigneeIDs []string `json:"assignees"`
}

type transferDataset struct {
	Version   int                 `json:"version"`
	TodoTypes []*todoTypeTransfer `json:"todoTypes"`
	Todos     []*todoTransfer     `json:"todos"`
	Users     []*userTransfer     `json:"users"`
}

const importVersion = 2

func makeExportData(tti []*TodoType, ti []*Todo, ui []*User) *transferDataset {

	tt := make([]*todoTypeTransfer, len(tti))
	t := make([]*todoTransfer, len(ti))
	u := make([]*userTransfer, len(ui))

	for i, m := range tti {
		tt[i] = &todoTypeTransfer{
			ID:           m.ID,
			Name:         m.Name,
			Description:  m.Description,
			Color:        m.Color,
			ReminderTime: m.ReminderTime,
		}
	}

	for i, m := range ti {

		aIDs := make([]string, len(m.Assignees))

		for i, a := range m.Assignees {
			aIDs[i] = a.ID
		}

		t[i] = &todoTransfer{
			ID:          m.ID,
			TodoTypeID:  m.Type.ID,
			Title:       m.Title,
			Description: m.Description,
			DueDate:     m.DueDate.Format(time.RFC3339),
			Completed:   m.Completed,
			AssigneeIDs: aIDs,
		}
	}

	for i, m := range ui {
		u[i] = &userTransfer{
			ID:        m.ID,
			FirstName: m.FirstName,
			LastName:  m.LastName,
			Email:     m.Email,
		}
	}

    return &transferDataset{
		Version:   importVersion,
		TodoTypes: tt,
		Todos:     t,
		Users:     u,
	}
}

func (s *APIServer) handleImportData(w http.ResponseWriter, r *http.Request) error {

    input := new(transferDataset)

	if err := json.NewDecoder(r.Body).Decode(input); err != nil {
		return err
	}

	if input.Version != importVersion {
		return writeJSON(w, http.StatusBadRequest, map[string]interface{}{
			"message": "Die Daten konnten nicht importiert werden, da sie von einer anderen Version stammen.",
			"error":   "invalid version",
		})
	}

	log.Println("Importing data, clearing database.")

	if err := s.store.Clear(); err != nil {
		return err
	}

	users := make(map[string]*User)
	for _, m := range input.Users {
		u, err := s.store.CreateUser(NewUser(m.FirstName, m.LastName, m.Email))

		if err != nil {
			return err
		}

		users[m.ID] = u
	}

	todoTypes := make(map[string]*TodoType)
	for _, m := range input.TodoTypes {
		tt, err := s.store.CreateTodoType(NewTodoType(m.Name, m.Description, m.Color, m.ReminderTime))

		if err != nil {
			return err
		}

		todoTypes[m.ID] = tt
	}

	for _, m := range input.Todos {

		dueDate, err := time.Parse(time.RFC3339, m.DueDate)

		if err != nil {
			return err
		}

		fmt.Println("Due date is", dueDate)

		assignees := make([]*User, len(m.AssigneeIDs))

		for i, aID := range m.AssigneeIDs {
			assignees[i] = users[aID]
		}

		_, err = s.store.CreateTodo(NewTodo(
			todoTypes[m.TodoTypeID],
			m.Title,
			m.Description,
			dueDate,
			m.Completed,
			assignees,
		))

		if err != nil {
			return err
		}
	}

	return writeJSON(w, http.StatusOK, map[string]interface{}{
		"message": "Daten wurden erfolgreich importiert.",
		"error":   nil,
	})
}

func (s *APIServer) handleExportData(w http.ResponseWriter, r *http.Request) error {

	todoTypes, err := s.store.GetTodoTypes()
	todos, err := s.store.GetTodos()
	users, err := s.store.GetUsers()

	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, makeExportData(todoTypes, todos, users))
}
