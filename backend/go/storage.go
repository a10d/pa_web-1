package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

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

type SqliteStorage struct {
	db *sql.DB
}

func NewSqliteStorage(databaseName string) (*SqliteStorage, error) {

	db, err := sql.Open("sqlite3", databaseName)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Println("Failed connecting to database: ", err)
		return nil, err
	}

	s := &SqliteStorage{
		db: db,
	}

	err = s.init()

	return s, err
}

func (s *SqliteStorage) init() error {
	return s.migrate()
}

func (s *SqliteStorage) migrate() error {

	migrationQuery, err := os.ReadFile("migration.sql")

	if err != nil {
		return err
	}

	_, err = s.db.Exec(string(migrationQuery))

	return err
}

func (s *SqliteStorage) GetTodos() ([]*Todo, error) {

	rows, err := s.db.Query(`
    select
        todos.id,
        todos.title,
        todos.description,
        todos.dueDate,
        todos.completed,
        todoTypes.id as typeId,
        todoTypes.name as typeName,
        todoTypes.description as typeDescription,
        todoTypes.color as typeColor,
        todoTypes.reminderTime as typeReminderTime
    from todos
        left join todoTypes on todoTypes.id = todos.todoTypeId
        join todoAssignees on todoAssignees.todoId = todos.id
        join users on users.id = todoAssignees.userId
    order by todos.dueDate `)

	if err != nil {
		return nil, err
	}

	todos := make([]*Todo, 0)

	for rows.Next() {

		var todo Todo

		if err = rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.Description,
			&todo.DueDate,
			&todo.Completed,
			&todo.Type.ID,
			&todo.Type.Name,
			&todo.Type.Description,
			&todo.Type.Color,
			&todo.Type.ReminderTime,
		); err != nil {
			return nil, err
		}

		todos = append(todos, &todo)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return todos, rows.Close()
}

func (s *SqliteStorage) GetTodoById(id string) (*Todo, error) {

	return nil, nil
}

func (s *SqliteStorage) CreateTodo(todo *Todo) (*Todo, error) {
	return todo, nil
}

func (s *SqliteStorage) UpdateTodo(todo *Todo) (*Todo, error) {
	return todo, nil
}

func (s *SqliteStorage) DeleteTodo(todo *Todo) error {
	return nil
}

func rowsToTodoTypes(rows *sql.Rows) ([]*TodoType, error) {
    todoTypes := make([]*TodoType, 0)

    for rows.Next() {
        var todoType TodoType

        if err := rows.Scan(
            &todoType.ID,
            &todoType.Name,
            &todoType.Description,
            &todoType.Color,
            &todoType.ReminderTime,
            &todoType.CreatedAt,
            &todoType.UpdatedAt,
        ); err != nil {
            return nil, err
        }

        todoTypes = append(todoTypes, &todoType)
    }

    return todoTypes, rows.Err()
}

func (s *SqliteStorage) GetTodoTypes() ([]*TodoType, error) {
	return nil, nil
}

func (s *SqliteStorage) GetTodoTypeById(id string) (*TodoType, error) {
	return nil, nil
}

func (s *SqliteStorage) CreateTodoType(todoType *TodoType) (*TodoType, error) {
	return todoType, nil
}

func (s *SqliteStorage) UpdateTodoType(todoType *TodoType) (*TodoType, error) {
	return todoType, nil
}

func (s *SqliteStorage) DeleteTodoType(todoType *TodoType) error {
	_, err := s.db.Exec(`delete from todoTypes where id = ?`, todoType.ID)
	return err
}

func rowsToUsers(rows *sql.Rows) ([]*User, error) {
	users := make([]*User, 0)

	for rows.Next() {
		var user User

		if err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, rows.Err()
}

func (s *SqliteStorage) GetUsers() ([]*User, error) {
	rows, err := s.db.Query(`
        select
            id,
            firstName,
            lastName,
            email,
            createdAt,
            updatedAt
        from users
        order by lastName, firstName
    `)

	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println("Error closing rows: ", err)
		}
	}(rows)

	return rowsToUsers(rows)
}

func (s *SqliteStorage) GetUserById(id string) (*User, error) {

	result := s.db.QueryRow(`select * from users where id = ?`, id)

	var user User

	if err := result.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *SqliteStorage) GetUsersByIds(ids []string) ([]*User, error) {

	rows, err := s.db.Query(`
        select
            id,
            firstName,
            lastName,
            email,
            createdAt,
            updatedAt
        from users
        where users.id in (?)
    `, ids)

	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println("Error closing rows: ", err)
		}
	}(rows)

	return rowsToUsers(rows)
}

func (s *SqliteStorage) CreateUser(user *User) (*User, error) {

	_, err := s.db.Exec(`
        insert into users (id, firstName, lastName, email)
        values (?, ?, ?, ?)`,
		user.ID,
		user.FirstName,
		user.LastName,
		user.Email)

	if err != nil {
		return nil, err
	}

	return s.GetUserById(user.ID)
}

func (s *SqliteStorage) UpdateUser(user *User) (*User, error) {

	_, err := s.db.Exec(`
        update users set
            firstName = ?,
            lastName = ?,
            email = ?
        where id = ?`,
		user.FirstName,
		user.LastName,
		user.Email,
		user.ID)

	if err != nil {
		return nil, err
	}

	return s.GetUserById(user.ID)
}

func (s *SqliteStorage) DeleteUser(user *User) error {

	_, err := s.db.Exec(`
        delete from users
        where id = ? `,
		user.ID)

	return err
}
