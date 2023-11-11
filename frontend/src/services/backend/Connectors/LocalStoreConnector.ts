import {Connector, Todo, TodoType, User} from "../index.ts";


export class LocalStoreConnector implements Connector {

    private users: User[] = [];
    private todos: Todo[] = [];
    private todoTypes: TodoType[] = [];

    private newId(): number {
        return Math.floor(Math.random() * 100000);
    }

    private writeChanges() {
        localStorage.setItem("users", JSON.stringify(this.users));
        localStorage.setItem("todos", JSON.stringify(this.todos));
        localStorage.setItem("todoTypes", JSON.stringify(this.todoTypes));
        this.loadChanges();
    }

    private loadChanges() {
        this.users = [];
        this.todos = [];
        this.todoTypes = [];

        const users = localStorage.getItem("users");
        if (users) {
            this.users = JSON.parse(users);
        }

        const todos = localStorage.getItem("todos");
        if (todos) {
            this.todos = JSON.parse(todos);
        }

        const todoTypes = localStorage.getItem("todoTypes");
        if (todoTypes) {
            this.todoTypes = JSON.parse(todoTypes);
        }
    }

    constructor() {
        this.loadChanges();
    }

    createTodo(todo: Partial<Todo>): Promise<Todo> {
        todo.id = this.newId();

        this.todos.push(todo as Todo);
        this.writeChanges();

        return Promise.resolve(todo as Todo);
    }

    createTodoType(todoType: Partial<TodoType>): Promise<TodoType> {
        todoType.id = this.newId();
        this.todoTypes.push(todoType as TodoType);
        this.writeChanges();

        return Promise.resolve(todoType as TodoType);
    }

    createUser(user: Partial<User>): Promise<User> {
        user.id = this.newId();
        this.users.push(user as User);
        this.writeChanges();

        return Promise.resolve(user as User);
    }

    deleteTodo(todo: Todo): Promise<boolean> {
        const index = this.todos.findIndex((t) => t.id === todo.id);
        if (index === -1) {
            return Promise.reject(new Error("Todo not found"));
        }

        this.todos.splice(index, 1);
        this.writeChanges();

        return Promise.resolve(true);
    }

    deleteTodoType(todoType: TodoType): Promise<boolean> {
        const index = this.todoTypes.findIndex((t) => t.id === todoType.id);

        if (index === -1) {
            return Promise.reject(new Error("TodoType not found"));
        }

        this.todoTypes.splice(index, 1);
        this.writeChanges();

        return Promise.resolve(true);
    }

    deleteUser(user: User): Promise<boolean> {
        const index = this.users.findIndex((u) => u.id === user.id);

        if (index === -1) {
            return Promise.reject(new Error("User not found"));
        }

        this.users.splice(index, 1);
        this.writeChanges();

        return Promise.resolve(true);
    }

    fetchTodoTypes(): Promise<TodoType[]> {
        return Promise.resolve(this.todoTypes);
    }

    fetchTodos(): Promise<Todo[]> {
        return Promise.resolve(this.todos);
    }

    fetchUsers(): Promise<User[]> {
        return Promise.resolve(this.users);
    }

    updateTodo(todo: Todo): Promise<Todo> {
        const index = this.todos.findIndex((t) => t.id === todo.id);
        if (index === -1) {
            return Promise.reject(new Error("Todo not found"));
        }

        this.todos[index] = todo;
        this.writeChanges();

        return Promise.resolve(todo);
    }

    updateTodoType(todoType: TodoType): Promise<TodoType> {
        const index = this.todoTypes.findIndex((t) => t.id === todoType.id);
        if (index === -1) {
            return Promise.reject(new Error("TodoType not found"));
        }

        this.todoTypes[index] = todoType;
        this.writeChanges();

        return Promise.resolve(todoType);
    }

    updateUser(user: User): Promise<User> {
        const index = this.users.findIndex((u) => u.id === user.id);
        if (index === -1) {
            return Promise.reject(new Error("User not found"));
        }

        this.users[index] = user;
        this.writeChanges();

        return Promise.resolve(user);
    }


}
