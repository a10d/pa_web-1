package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
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

    router.HandleFunc("/todos", makeServeHandleFunc(s.handleGetTodos)).Methods(http.MethodGet)
    router.HandleFunc("/todos", makeServeHandleFunc(s.handleCreateTodo)).Methods(http.MethodPost)
    router.HandleFunc("/todos/{id}", makeServeHandleFunc(s.handleGetTodoById)).Methods(http.MethodGet)
    router.HandleFunc("/todos/{id}", makeServeHandleFunc(s.handleUpdateTodo)).Methods(http.MethodPut)
    router.HandleFunc("/todos/{id}", makeServeHandleFunc(s.handleDeleteTodo)).Methods(http.MethodDelete)

    router.HandleFunc("/todoTypes", makeServeHandleFunc(s.handleGetTodoTypes)).Methods(http.MethodGet)
    router.HandleFunc("/todoTypes", makeServeHandleFunc(s.handleCreateTodoType)).Methods(http.MethodPost)
    router.HandleFunc("/todoTypes/{id}", makeServeHandleFunc(s.handleGetTodoTypeById)).Methods(http.MethodGet)
    router.HandleFunc("/todoTypes/{id}", makeServeHandleFunc(s.handleUpdateTodoType)).Methods(http.MethodPut)
    router.HandleFunc("/todoTypes/{id}", makeServeHandleFunc(s.handleDeleteTodoType)).Methods(http.MethodDelete)

    router.HandleFunc("/users", makeServeHandleFunc(s.handleGetUsers)).Methods(http.MethodGet)
    router.HandleFunc("/users", makeServeHandleFunc(s.handleCreateUser)).Methods(http.MethodPost)
    router.HandleFunc("/users/{id}", makeServeHandleFunc(s.handleGetUserById)).Methods(http.MethodGet)
    router.HandleFunc("/users/{id}", makeServeHandleFunc(s.handleUpdateUser)).Methods(http.MethodPut)
    router.HandleFunc("/users/{id}", makeServeHandleFunc(s.handleDeleteUser)).Methods(http.MethodDelete)

	log.Println("Listening on", s.listenAddr)

	return http.ListenAndServe(s.listenAddr, router)
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
