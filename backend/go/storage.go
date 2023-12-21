package main

type Storage interface {
	GetTodos() ([]*Todo, error)
	GetTodoById(id string) (*Todo, error)
	CreateTodo(*Todo) (*Todo, error)
	UpdateTodo(*Todo) (*Todo, error)
	DeleteTodo(*Todo) error

	GetTodoTypes() ([]*TodoType, error)
	GetTodoTypeById(id string) (*TodoType, error)
	CreateTodoType(*TodoType) (*TodoType, error)
	UpdateTodoType(*TodoType) (*TodoType, error)
	DeleteTodoType(*TodoType) error

	GetUsers() ([]*User, error)
	GetUserById(id string) (*User, error)
	GetUsersByIds(ids []string) ([]*User, error)
	CreateUser(*User) (*User, error)
	UpdateUser(*User) (*User, error)
	DeleteUser(*User) error
}
