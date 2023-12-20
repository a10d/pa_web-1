package main

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "net/http"
    "time"
)

type todoInput struct {
    TodoTypeID  string    `json:"type"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    DueDate     time.Time `json:"dueDate"`
    AssigneeIDs []string  `json:"assignees"`
}

func (s *APIServer) handleGetTodos(w http.ResponseWriter, r *http.Request) error {

    todos, err := s.store.GetTodos()

    if err != nil {
        return err
    }

    return writeJSON(w, http.StatusOK, todos)
}

func (s *APIServer) handleCreateTodo(w http.ResponseWriter, r *http.Request) error {

    createTodoReq := new(todoInput)

    if err := json.NewDecoder(r.Body).Decode(createTodoReq); err != nil {
        return err
    }

    todoType, err := s.store.GetTodoTypeById(createTodoReq.TodoTypeID)

    if err != nil {
        return err
    }

    assignees := make([]*User, 0)

    for _, assigneeID := range createTodoReq.AssigneeIDs {
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

func (s *APIServer) handleGetTodoById(w http.ResponseWriter, r *http.Request) error {

    todo, err := s.store.GetTodoById(mux.Vars(r)["id"])

    if err != nil {
        return err
    }

    return writeJSON(w, http.StatusOK, todo)
}

func (s *APIServer) handleUpdateTodo(w http.ResponseWriter, r *http.Request) error {
    return nil
}

func (s *APIServer) handleDeleteTodo(w http.ResponseWriter, r *http.Request) error {

    todo, err := s.store.GetTodoById(mux.Vars(r)["id"])

    if err != nil {
        return err
    }

    err = s.store.DeleteTodo(todo)

    if err != nil {
        return err
    }

    w.WriteHeader(http.StatusNoContent)
    return nil
}
