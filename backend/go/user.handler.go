package main

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "net/http"
)

type userInput struct {
    FirstName string `json:"firstName"`
    LastName  string `json:"lastName"`
    Email     string `json:"email"`
}

func (s *APIServer) handleGetUsers(w http.ResponseWriter, r *http.Request) error {

    users, err := s.store.GetUsers()

    if err != nil {
        return err
    }

    return writeJSON(w, http.StatusOK, users)
}

func (s *APIServer) handleCreateUser(w http.ResponseWriter, r *http.Request) error {

    input := new(userInput)

    if err := json.NewDecoder(r.Body).Decode(input); err != nil {
        return err
    }

    // TODO: Handle validation

    user, err := s.store.CreateUser(NewUser(
        input.FirstName,
        input.LastName,
        input.Email))

    if err != nil {
        return err
    }

    return writeJSON(w, http.StatusOK, user)
}

func (s *APIServer) handleGetUserById(w http.ResponseWriter, r *http.Request) error {

    user, err := s.store.GetUserById(mux.Vars(r)["id"])

    if err != nil {
        return err
    }

    return writeJSON(w, http.StatusOK, user)
}

func (s *APIServer) handleUpdateUser(w http.ResponseWriter, r *http.Request) error {

    user, err := s.store.GetUserById(mux.Vars(r)["id"])

    if err != nil {
        return err
    }

    input := new(userInput)

    if err := json.NewDecoder(r.Body).Decode(input); err != nil {
        return err
    }

    // TODO: Handle validation

    user.FirstName = input.FirstName
    user.LastName = input.LastName
    user.Email = input.Email

    updatedUser, err := s.store.UpdateUser(user)

    if err != nil {
        return err
    }

    return writeJSON(w, http.StatusOK, updatedUser)
}

func (s *APIServer) handleDeleteUser(w http.ResponseWriter, r *http.Request) error {

    user, err := s.store.GetUserById(mux.Vars(r)["id"])

    if err != nil {
        return err
    }

    err = s.store.DeleteUser(user)

    if err != nil {
        return err
    }

    w.WriteHeader(http.StatusNoContent)
    return nil
}
