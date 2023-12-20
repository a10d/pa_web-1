package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type APIServer struct {
	listenAddr string
	store      Storage
}

func NewAPIServer(listenAddr string, store Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()

	router.HandleFunc("/todos", makeServeHandleFunc(s.handleTodo))
	router.HandleFunc("/users", makeServeHandleFunc(s.handleUsers))
	router.HandleFunc("/todoTypes", makeServeHandleFunc(s.handleTodoTypes))

	log.Println("Listening on", s.listenAddr)

	return http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleTodo(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case http.MethodGet:
		return s.handleGetTodo(w, r)
	case http.MethodPost:
		return s.handleCreateTodo(w, r)
	case http.MethodPut:
		return s.handleUpdateTodo(w, r)
	case http.MethodDelete:
		return s.handleDeleteTodo(w, r)
	}
	return fmt.Errorf("method %s not allowed on this endpoint", r.Method)
}

func (s *APIServer) handleGetTodo(w http.ResponseWriter, r *http.Request) error {

	todo := NewTodo(
		NewTodoType(
			"Wichtige Aufgabe",
			"Wichtige Aufgaben, die sofort erledigt werden müssen",
			"#ff0000",
			2,
		),
		"Testaufgabe",
		"Das ist eine Testaufgabe",
		time.Now(),
		[]*User{
			NewUser("Peter", "Tester", "test@example.com"),
		},
	)

	return writeJSON(w, http.StatusOK, todo)
}

func (s *APIServer) handleCreateTodo(w http.ResponseWriter, r *http.Request) error {

	createTodoReq := new(struct {
		TodoTypeID  string    `json:"type"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		DueDate     time.Time `json:"dueDate"`
		Assignees   []string  `json:"assignees"`
	})

	if err := json.NewDecoder(r.Body).Decode(createTodoReq); err != nil {
		return err
	}

	todoType, err := s.store.GetTodoTypeById(createTodoReq.TodoTypeID)

	if err != nil {
		return err
	}

	assignees := make([]*User, 0)

	for _, assigneeID := range createTodoReq.Assignees {
		assignee, err := s.store.GetUserById(assigneeID)

		if err != nil {
			return err
		}

		assignees = append(assignees, assignee)
	}

	todo := NewTodo(
		todoType,
		createTodoReq.Title,
		createTodoReq.Description,
		createTodoReq.DueDate,
		assignees,
	)

	createTodo, err := s.store.CreateTodo(todo)

	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, createTodo)
}

func (s *APIServer) handleUpdateTodo(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteTodo(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleUsers(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case http.MethodGet:
		return s.handleGetUsers(w, r)
	case http.MethodPost:
		return s.handleCreateUser(w, r)
	case http.MethodPut:
	case http.MethodDelete:
	}
	return fmt.Errorf("method %s not allowed on this endpoint", r.Method)
}

func (s *APIServer) handleGetUsers(w http.ResponseWriter, r *http.Request) error {
	users, err := s.store.GetUsers()

	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, users)
}

func (s *APIServer) handleCreateUser(w http.ResponseWriter, r *http.Request) error {

	input := new(struct {
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Email     string `json:"email"`
	})

	if err := json.NewDecoder(r.Body).Decode(input); err != nil {
		return err
	}

	user, err := s.store.CreateUser(NewUser(
		input.FirstName,
		input.LastName,
		input.Email))

	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, user)
}

func (s *APIServer) handleTodoTypes(w http.ResponseWriter, r *http.Request) error {
	switch r.Method {
	case http.MethodGet:
		return s.handleGetTodoTypes(w, r)
	case http.MethodPost:
	case http.MethodPut:
	case http.MethodDelete:
	}
	return fmt.Errorf("method %s not allowed on this endpoint", r.Method)
}

func (s *APIServer) handleGetTodoTypes(w http.ResponseWriter, r *http.Request) error {
	todoTypes, err := s.store.GetTodoTypes()

	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, todoTypes)
}

func writeJSON(w http.ResponseWriter, status int, v interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type apiError struct {
	Error string `json:"error"`
}

func makeServeHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			err := writeJSON(w, http.StatusInternalServerError, apiError{Error: err.Error()})
			if err != nil {
				return
			}
		}
	}
}
