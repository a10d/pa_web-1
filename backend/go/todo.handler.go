package main

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "github.com/thedevsaddam/govalidator"
    "net/http"
    "net/url"
    "time"
)

type todoInput struct {
    TodoTypeID  string   `json:"type"`
    Title       string   `json:"title"`
    Description string   `json:"description"`
    DueDate     string   `json:"dueDate"`
    Completed   bool     `json:"completed"`
    AssigneeIDs []string `json:"assignees"`
}

func (input *todoInput) validateOnCreate() url.Values {

    return govalidator.New(govalidator.Options{
        Data: input,
        Rules: govalidator.MapData{
            "type":        []string{"isTodoType"},
            "title":       []string{"min:1", "max:255"},
            "description": []string{"min:1", "max:3000"},
            "dueDate":     []string{"datetime", "afterToday"},
            "assignees":   []string{"isListOfUsers"},
        },
        RequiredDefault: true,
    }).ValidateStruct()
}

func (input *todoInput) validateOnUpdate() url.Values {

    return govalidator.New(govalidator.Options{
        Data: input,
        Rules: govalidator.MapData{
            "type":        []string{"isTodoType"},
            "title":       []string{"min:1", "max:255"},
            "description": []string{"min:1", "max:3000"},
            "dueDate":     []string{"datetime"},
            "assignees":   []string{"isListOfUsers"},
            "completed":   []string{"bool"},
        },
        RequiredDefault: true,
    }).ValidateStruct()
}

func (s *APIServer) handleGetTodos(w http.ResponseWriter, r *http.Request) error {

    todos, err := s.store.GetTodos()

    if err != nil {
        return err
    }

    return writeJSON(w, http.StatusOK, todos)
}

func (s *APIServer) handleCreateTodo(w http.ResponseWriter, r *http.Request) error {

    input := new(todoInput)

    if err := json.NewDecoder(r.Body).Decode(input); err != nil {
        return err
    }

    e := input.validateOnCreate()

    if len(e) > 0 {
        return writeValidationError(w, e)
    }

    todoType, err := s.store.GetTodoTypeById(input.TodoTypeID)
    assignees, err := s.store.GetUsersByIds(input.AssigneeIDs)

    if err != nil {
        return err
    }

    dueDate, err := time.Parse(time.RFC3339, input.DueDate)

    createTodo, err := s.store.CreateTodo(NewTodo(
        todoType,
        input.Title,
        input.Description,
        dueDate,
        assignees,
    ))

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

    todo, err := s.store.GetTodoById(mux.Vars(r)["id"])

    if err != nil {
        return err
    }

    input := new(todoInput)

    if err = json.NewDecoder(r.Body).Decode(input); err != nil {
        return err
    }

    e := input.validateOnUpdate()

    if len(e) > 0 {
        return writeValidationError(w, e)
    }

    if input.TodoTypeID != todo.Type.ID {
        todoType, err := s.store.GetTodoTypeById(input.TodoTypeID)

        if err != nil {
            return err
        }

        todo.Type = todoType
    }

    dueDate, err := time.Parse(time.RFC3339, input.DueDate)
    if err != nil {
        return err
    }

    assignees, err := s.store.GetUsersByIds(input.AssigneeIDs)
    if err != nil {
        return err
    }

    todo.Assignees = assignees
    todo.Title = input.Title
    todo.Description = input.Description
    todo.DueDate = dueDate
    todo.Completed = input.Completed

    updateTodo, err := s.store.UpdateTodo(todo)

    if err != nil {
        return err
    }

    return writeJSON(w, http.StatusOK, updateTodo)
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
