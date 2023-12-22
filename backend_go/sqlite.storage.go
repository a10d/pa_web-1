package main

import (
    "database/sql"
    "fmt"
    "log"
    _ "modernc.org/sqlite"
    "os"
    "strings"
)

type SqliteStorage struct {
	db *sql.DB
}

func NewSqliteStorage(databaseName string) (*SqliteStorage, error) {

    db, err := sql.Open("sqlite", databaseName)

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
            todos.createdAt,
            todos.updatedAt,
            todoTypes.id as type_id,
            todoTypes.name as type_name,
            todoTypes.description as type_description,
            todoTypes.color as type_color,
            todoTypes.reminderTime as type_reminderTime,
            todoTypes.createdAt as type_createdAt,
            todoTypes.updatedAt as type_updatedAt,

            case when (select count(*) from todoAssignees where todoAssignees.todoId = todos.id) > 0
                    then
                (select group_concat(tA.userId) from todoAssignees tA where tA.todoId = todos.id)
                else
                    ''
                end
                as assigneeIds

        from todos

            left join todoTypes on todoTypes.id = todos.todoTypeId

        order by completed desc, todos.dueDate`)

	if err != nil {

        fmt.Println("ERROR:", err)

		return nil, err
	}

	todos := make([]*Todo, 0)

	for rows.Next() {

		todo := new(Todo)
		todo.Type = new(TodoType)
		todo.Assignees = make([]*User, 0)

		assigneeIds := ""

		if err = rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.Description,
			&todo.DueDate,
			&todo.Completed,
			&todo.CreatedAt,
			&todo.UpdatedAt,

			&todo.Type.ID,
			&todo.Type.Name,
			&todo.Type.Description,
			&todo.Type.Color,
			&todo.Type.ReminderTime,
			&todo.Type.CreatedAt,
			&todo.Type.UpdatedAt,

			&assigneeIds,
		); err != nil {
			return nil, err
		}

		todo.Assignees, err = s.GetUsersByIds(strings.Split(assigneeIds, ","))

		if err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return todos, rows.Close()
}

func (s *SqliteStorage) GetTodoById(id string) (*Todo, error) {

	todo := new(Todo)
	todo.Type = new(TodoType)

	var assigneeIds string

	result := s.db.QueryRow(`
        select
            todos.id,
            todos.title,
            todos.description,
            todos.dueDate,
            todos.completed,
            todos.createdAt,
            todos.updatedAt,

            todoTypes.id as type_id,
            todoTypes.name as type_name,
            todoTypes.description as type_description,
            todoTypes.color as type_color,
            todoTypes.reminderTime as type_reminderTime,
            todoTypes.createdAt as type_createdAt,
            todoTypes.updatedAt as type_updatedAt,

            case when (select count(*) from todoAssignees where todoAssignees.todoId = todos.id) > 0
                    then
                (select group_concat(tA.userId) from todoAssignees tA where tA.todoId = todos.id)
                else
                    ''
                end
                as assigneeIds

        from todos

            left join todoTypes on todoTypes.id = todos.todoTypeId

        where todos.id = ?;`,
		id,
	)

	if err := result.Scan(
		&todo.ID,
		&todo.Title,
		&todo.Description,
		&todo.DueDate,
		&todo.Completed,
		&todo.CreatedAt,
		&todo.UpdatedAt,

		&todo.Type.ID,
		&todo.Type.Name,
		&todo.Type.Description,
		&todo.Type.Color,
		&todo.Type.ReminderTime,
		&todo.Type.CreatedAt,
		&todo.Type.UpdatedAt,

		&assigneeIds,
	); err != nil {
		return nil, err
	}

	assignees, err := s.GetUsersByIds(strings.Split(assigneeIds, ","))

	if err != nil {
		return nil, err
	}

	todo.Assignees = assignees

	return todo, nil

}

func (s *SqliteStorage) CreateTodo(todo *Todo) (*Todo, error) {

	tx, err := s.db.Begin()

	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare(`
        insert into todos (id, todoTypeId, title, description, dueDate, completed)
        values (?, ?, ?, ?, ?, ?);`,
	)

	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(todo.ID, todo.Type.ID, todo.Title, todo.Description, todo.DueDate, todo.Completed)

	if err != nil {
		return nil, err
	}

	stmt, err = tx.Prepare(`
        insert into todoAssignees (todoId, userId)
        values (?, ?);`,
	)

	if err != nil {
		return nil, err
	}

	for _, assignee := range todo.Assignees {
		_, err = stmt.Exec(todo.ID, assignee.ID)

		if err != nil {
			return nil, err
		}
	}

	err = tx.Commit()

	if err != nil {
		return nil, err
	}

	return s.GetTodoById(todo.ID)
}

func (s *SqliteStorage) UpdateTodo(todo *Todo) (*Todo, error) {

	tx, err := s.db.Begin()

	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare(`
        update todos set
            todoTypeId = ?,
            title = ?,
            description = ?,
            dueDate = ?,
            completed = ?
        where id = ?;`,
	)

	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(
		todo.Type.ID,
		todo.Title,
		todo.Description,
		todo.DueDate,
		todo.Completed,
		todo.ID,
	)

	if err != nil {
		return nil, err
	}

	stmt, err = tx.Prepare(`
        delete from todoAssignees
        where todoId = ?;`,
	)

	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(todo.ID)

	if err != nil {
		return nil, err
	}

	stmt, err = tx.Prepare(`
        insert into todoAssignees (todoId, userId)
        values (?, ?);`,
	)

	if err != nil {
		return nil, err
	}

	for _, assignee := range todo.Assignees {
		_, err = stmt.Exec(todo.ID, assignee.ID)

		if err != nil {
			return nil, err
		}
	}

	err = tx.Commit()

	if err != nil {
		return nil, err
	}

	return s.GetTodoById(todo.ID)
}

func (s *SqliteStorage) DeleteTodo(todo *Todo) error {

	_, err := s.db.Exec(`delete from todos where id = ?`, todo.ID)

	return err
}

func (s *SqliteStorage) GetTodoTypes() ([]*TodoType, error) {

	rows, err := s.db.Query(`
        select
            id,
            name,
            description,
            color,
            reminderTime,
            createdAt,
            updatedAt
        from todoTypes
        order by name;
    `)

	if err != nil {
		return nil, err
	}

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

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return todoTypes, rows.Close()
}

func (s *SqliteStorage) GetTodoTypeById(id string) (*TodoType, error) {

	var todoType TodoType

	result := s.db.QueryRow(`
        select
            id,
            name,
            description,
            color,
            reminderTime,
            createdAt,
            updatedAt
        from todoTypes
        where id = ?;`,
		id,
	)

	if err := result.Scan(
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

	return &todoType, nil
}

func (s *SqliteStorage) CreateTodoType(todoType *TodoType) (*TodoType, error) {

	_, err := s.db.Exec(`
        insert into todoTypes (id, name, description, color, reminderTime)
        values (?, ?, ?, ?, ?);`,
		todoType.ID,
		todoType.Name,
		todoType.Description,
		todoType.Color,
		todoType.ReminderTime,
	)

	if err != nil {
		return nil, err
	}

	return s.GetTodoTypeById(todoType.ID)
}

func (s *SqliteStorage) UpdateTodoType(todoType *TodoType) (*TodoType, error) {

	_, err := s.db.Exec(`
        update todoTypes set
            name = ?,
            description = ?,
            color = ?,
            reminderTime = ?
        where id = ?;`,
		todoType.Name,
		todoType.Description,
		todoType.Color,
		todoType.ReminderTime,
		todoType.ID,
	)

	if err != nil {
		return nil, err
	}

	return s.GetTodoTypeById(todoType.ID)
}

func (s *SqliteStorage) DeleteTodoType(todoType *TodoType) error {

	_, err := s.db.Exec(`delete from todoTypes where id = ?;`, todoType.ID)

	return err
}

func rowsToUsers(rows *sql.Rows) ([]*User, error) {

	users := make([]*User, 0)

	for rows.Next() {

		user := new(User)

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

		users = append(users, user)
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
        order by lastName, firstName;
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

	result := s.db.QueryRow(`select * from users where id = ?;`, id)

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

	if len(ids) == 0 {
		return nil, nil
	}

	query := fmt.Sprintf(`select
            id,
            firstName,
            lastName,
            email,
            createdAt,
            updatedAt
        from users
        where users.id in ( %s );`,
		"?"+strings.Repeat(", ?", len(ids)-1),
	)

	args := make([]interface{}, len(ids))

	for i, id := range ids {
		args[i] = id
	}

	rows, err := s.db.Query(query, args...)

	if err != nil {
		return nil, err
	}

	users, err := rowsToUsers(rows)

	if err != nil {
		return nil, err
	}

	return users, rows.Close()
}

func (s *SqliteStorage) CreateUser(user *User) (*User, error) {

	_, err := s.db.Exec(`
        insert into users (id, firstName, lastName, email)
        values (?, ?, ?, ?);`,
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email)

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
        where id = ?;`,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.ID)

	if err != nil {
		return nil, err
	}

	return s.GetUserById(user.ID)
}

func (s *SqliteStorage) DeleteUser(user *User) error {

	_, err := s.db.Exec(`
        delete from users
        where id = ?;`,
		&user.ID)

	return err
}

func (s *SqliteStorage) Clear() error {

    _, err := s.db.Exec(`
        delete from todos;
        delete from todoTypes;
        delete from users;
        delete from todoAssignees;
    `)

    return err

}
