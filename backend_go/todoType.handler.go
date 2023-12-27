package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
    "github.com/thedevsaddam/govalidator"
	"net/http"
    "net/url"
)

type todoTypeInput struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	Color        string `json:"color"`
	ReminderTime int    `json:"reminderTime"`
}

func (input *todoTypeInput) validate() url.Values {

    return govalidator.New(govalidator.Options{
        Data: input,
        Rules: govalidator.MapData{
            "name":         []string{"min:3", "max:255"},
            "description":  []string{"min:3", "max:3000"},
            "color":        []string{"css_color"},
            "reminderTime": []string{"numeric_between:1,60"},
        },
        Messages: govalidator.MapData{
            "name": []string{
                "min:Der Name des Aufgabentyps muss mindestens drei Zeichen lang sein.",
                "max:Der Name des Aufgabentyps darf nicht länger als 255 Zeichen sein.",
            },
            "description": []string{
                "min:Die Beschreibung muss mindestens drei Zeichen lang sein.",
                "max:Die Beschreibung darf nicht länger als 3000 Zeichen sein. (Ja, das reicht wirklich.)",
            },
            "color": []string{
                "css_color:Die Farbe muss ein gültiger CSS-Farbcode sein.",
            },
            "reminderTime": []string{
                "numeric_between:Die Erinnerungszeit muss zwischen 1 und 60 Tagen liegen.",
            },
        },
        RequiredDefault: true,
    }).ValidateStruct()
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

    e := input.validate()

    if len(e) > 0 {
        return writeValidationError(w, e)
    }

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

    e := input.validate()

    if len(e) > 0 {
        return writeValidationError(w, e)
    }

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
