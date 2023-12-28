package main

import (
	"github.com/google/uuid"
	"time"
)

func newID() string {
	return uuid.New().String()
}

type Todo struct {
	ID          string    `json:"id"`
	Type        *TodoType `json:"type"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"dueDate"`
	Assignees   []*User   `json:"assignees"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func NewTodo(todoType *TodoType, title string, description string, dueDate time.Time, completed bool, assignees []*User) *Todo {
	return &Todo{
		ID:          newID(),
		Type:        todoType,
		Title:       title,
		Description: description,
		DueDate:     dueDate,
        Completed: completed,
		Assignees:   assignees,
	}
}

type TodoType struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Color        string    `json:"color"`
	ReminderTime int       `json:"reminderTime"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

func NewTodoType(name string, description string, color string, reminderTime int) *TodoType {
	return &TodoType{
		ID:           newID(),
		Name:         name,
		Description:  description,
		Color:        color,
		ReminderTime: reminderTime,
	}
}

type User struct {
	ID        string    `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewUser(firstName string, lastName string, email string) *User {
	return &User{
		ID:        newID(),
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}
}
