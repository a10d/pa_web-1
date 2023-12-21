package main

import (
	"encoding/json"
    "errors"
    "fmt"
	"github.com/gorilla/mux"
    "github.com/thedevsaddam/govalidator"
	"log"
	"net/http"
    "net/url"
    "time"
)

type APIServer struct {
    router *mux.Router
	listenAddr string
	store      Storage
}

func NewAPIServer(listenAddr string, store Storage) *APIServer {
	return &APIServer{
        router: mux.NewRouter(),
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) Run() error {
    s.registerRoutes()
    s.registerCustomValidationRules()

	log.Println("Listening on", s.listenAddr)

    return http.ListenAndServe(s.listenAddr, s.router)
}

func (s *APIServer) registerRoutes() {
    s.router.HandleFunc("/todos", makeServeHandleFunc(s.handleGetTodos)).Methods(http.MethodGet)
    s.router.HandleFunc("/todos", makeServeHandleFunc(s.handleCreateTodo)).Methods(http.MethodPost)
    s.router.HandleFunc("/todos/{id}", makeServeHandleFunc(s.handleGetTodoById)).Methods(http.MethodGet)
    s.router.HandleFunc("/todos/{id}", makeServeHandleFunc(s.handleUpdateTodo)).Methods(http.MethodPut)
    s.router.HandleFunc("/todos/{id}", makeServeHandleFunc(s.handleDeleteTodo)).Methods(http.MethodDelete)

    s.router.HandleFunc("/todoTypes", makeServeHandleFunc(s.handleGetTodoTypes)).Methods(http.MethodGet)
    s.router.HandleFunc("/todoTypes", makeServeHandleFunc(s.handleCreateTodoType)).Methods(http.MethodPost)
    s.router.HandleFunc("/todoTypes/{id}", makeServeHandleFunc(s.handleGetTodoTypeById)).Methods(http.MethodGet)
    s.router.HandleFunc("/todoTypes/{id}", makeServeHandleFunc(s.handleUpdateTodoType)).Methods(http.MethodPut)
    s.router.HandleFunc("/todoTypes/{id}", makeServeHandleFunc(s.handleDeleteTodoType)).Methods(http.MethodDelete)

    s.router.HandleFunc("/users", makeServeHandleFunc(s.handleGetUsers)).Methods(http.MethodGet)
    s.router.HandleFunc("/users", makeServeHandleFunc(s.handleCreateUser)).Methods(http.MethodPost)
    s.router.HandleFunc("/users/{id}", makeServeHandleFunc(s.handleGetUserById)).Methods(http.MethodGet)
    s.router.HandleFunc("/users/{id}", makeServeHandleFunc(s.handleUpdateUser)).Methods(http.MethodPut)
    s.router.HandleFunc("/users/{id}", makeServeHandleFunc(s.handleDeleteUser)).Methods(http.MethodDelete)
}

func (s *APIServer) registerCustomValidationRules() {

    govalidator.AddCustomRule("datetime", func(field string, rule string, message string, value interface{}) error {

        if value == nil {
            if message != "" {
                return errors.New(message)
            }
            return fmt.Errorf("The %s field must be a valid date and time", field)
        }

        _, err := time.Parse(time.RFC3339, value.(string))
        if err != nil {
            if message != "" {
                return errors.New(message)
            }
            return fmt.Errorf("The %s field must be a valid date and time", field)
        }

        return nil
    })

    govalidator.AddCustomRule("afterToday", func(field string, rule string, message string, value interface{}) error {

        date, err := time.Parse(time.RFC3339, value.(string))
        if err != nil {
            if message != "" {
                return errors.New(message)
            }
            return fmt.Errorf("The %s field must be a valid date and time", field)
        }

        if date.Before(time.Now()) {
            if message != "" {
                return errors.New(message)
            }
            return fmt.Errorf("The %s field must be after today", field)
        }

        return nil
    })

    govalidator.AddCustomRule("isTodoType", func(field string, rule string, message string, value interface{}) error {

        if value == nil {
            if message != "" {
                return errors.New(message)
            }
            return fmt.Errorf("The %s field must be a valid todo type", field)
        }

        todoType, err := s.store.GetTodoTypeById(value.(string))

        if err != nil || todoType == nil {
            if message != "" {
                return errors.New(message)
            }
            return fmt.Errorf("The %s field must be a valid todo type", field)
        }

        return nil
    })

    govalidator.AddCustomRule("isListOfUsers", func(field string, rule string, message string, value interface{}) error {

        ids := value.([]string)

        if value == nil || len(ids) == 0 {
            if message != "" {
                return errors.New(message)
            }
            return fmt.Errorf("The %s field must be a list of valid user ids", field)
        }

        users, err := s.store.GetUsersByIds(ids)

        if err != nil || len(users) != len(ids) {
            if message != "" {
                return errors.New(message)
            }
            return fmt.Errorf("The %s field must be a list of valid user ids", field)
        }

        return nil
    })
}

func writeJSON(w http.ResponseWriter, status int, v interface{}) error {

    w.Header().Add("Content-Type", "application/json")

    w.WriteHeader(status)

    return json.NewEncoder(w).Encode(v)
}

func writeValidationError(w http.ResponseWriter, e url.Values) error {
    return writeJSON(w, http.StatusBadRequest, map[string]interface{}{
        "error":      "validation error",
        "validation": e,
    })
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
