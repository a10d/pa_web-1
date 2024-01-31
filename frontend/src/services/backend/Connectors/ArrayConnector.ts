import {v4 as uuid} from "uuid";
import {Connector, Todo, TodoType, User} from "../index.ts";

export class ArrayConnector implements Connector {

    protected users: User[] = [];
    protected todos: Todo[] = [];
    protected todoTypes: TodoType[] = [];

    protected static IMPORT_VERSION = 2;

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

    createExport(): Promise<string> {
        return Promise.resolve(JSON.stringify({
            version: ArrayConnector.IMPORT_VERSION,
            users: this.users,
            todos: this.todos,
            todoTypes: this.todoTypes,
        }, null, 2));
    }

    importData(data: string): Promise<boolean> {
        try {
            const imported = JSON.parse(data) as Partial<{
                version: number,
                users: User[],
                todos: Todo[],
                todoTypes: TodoType[]
            }>;

            if (imported.version !== ArrayConnector.IMPORT_VERSION) {
                return Promise.reject(new Error("Invalid import signature"));
            }

            if (!imported.users || !imported.todos || !imported.todoTypes) {
                return Promise.reject(new Error("Invalid import contents"));
            }

            this.users = imported.users;
            this.todos = imported.todos;
            this.todoTypes = imported.todoTypes;

            this.writeChanges();
        } catch (e) {
            this.loadChanges();
            return Promise.reject(e);
        }

        return Promise.resolve(true);
    }

    protected newId(): string {
        return uuid();
    }

    protected writeChanges() {
        //
    }

    protected loadChanges() {
        this.users = [];
        this.todos = [];
        this.todoTypes = [];
    }


}
