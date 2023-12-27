import {Connector} from "./Connector.ts";
import {Todo, TodoType, User, ValidationError} from "../index.ts";


type ValidationErrorResponse = {
    error: string;
    validation?: {
        [key: string]: string[];
    }
}

export class ServerConnector implements Connector {

    private endpoint: string = "http://localhost:8080/";

    private async get<T>(uri: string): Promise<T> {
        const response = await this.fetch(uri);
        return response.json();
    }

    private async post<T>(uri: string, body: any): Promise<T> {
        return this.fetchJSON<T>(uri, {
            method: "POST",
        }, body);
    }

    private async put<T>(uri: string, body: any): Promise<T> {
        return this.fetchJSON<T>(uri, {
            method: "PUT",
        }, body);
    }

    private async delete(uri: string, options: any = {}): Promise<Response> {
        return this.fetch(uri, {
            method: "delete",
            ...options,
        });
    }

    private async fetch(uri: string, options: any = {}) {
        const response = await fetch(this.endpoint + uri, options);

        if (response.status === 400) {
            await this.recoverValidationError(response);
        }

        if (!response.ok) {
            throw response;
        }

        return response;
    }

    private async recoverValidationError(response: Response) {
        const error: ValidationErrorResponse = await response.json();

        if (error.error == "validation error") {
            throw new ValidationError(error.validation);
        }
    }

    private async fetchJSON<T>(uri: string, options: any = {}, body: any): Promise<T> {
        const response = await this.fetch(uri, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                ...(options.headers || {})
            },
            body: JSON.stringify(body),
            ...options,
        });

        return response.json();
    }


    async createExport(): Promise<string> {
        return JSON.stringify(await this.get<any>("export"), null, 2);
    }

    createTodo(todo: Partial<Todo>): Promise<Todo> {
        return this.post<Todo>("todos", {
            type: todo.type ? todo.type.id : null,
            title: todo.title,
            description: todo.description,
            dueDate: new Date(todo.dueDate ?? new Date()).toISOString(),
            completed: todo.completed,
            assignees: todo.assignees ? todo.assignees.map((user) => user.id) : [],
        });
    }

    createTodoType(todoType: Partial<TodoType>): Promise<TodoType> {
        return this.post<TodoType>("todoTypes", todoType);
    }

    createUser(user: Partial<User>): Promise<User> {
        return this.post<User>("users", user);
    }

    async deleteTodo(todo: Todo): Promise<boolean> {
        return (await this.delete(`todo/${todo.id}`)).status === 204;
    }

    async deleteTodoType(todoType: TodoType): Promise<boolean> {
        return (await this.delete(`todoType/${todoType.id}`)).status === 204;
    }

    async deleteUser(user: User): Promise<boolean> {
        return (await this.delete(`user/${user.id}`)).status === 204;
    }

    fetchTodoTypes(): Promise<TodoType[]> {
        return this.get<TodoType[]>("todoTypes");
    }

    fetchTodos(): Promise<Todo[]> {
        return this.get<Todo[]>("todos");
    }

    fetchUsers(): Promise<User[]> {
        return this.get<User[]>("users");
    }

    async importData(data: string): Promise<boolean> {
        const response = await this.post<{
            error: string | null
            success?: boolean
        }>("import", JSON.parse(data));

        return !!response.success && response.error === null;
    }

    updateTodo(todo: Todo): Promise<Todo> {
        return this.put<Todo>(`todo/${todo.id}`, {
            type: todo.type.id,
            title: todo.title,
            description: todo.description,
            dueDate: todo.dueDate,
            completed: todo.completed,
            assignees: todo.assignees.map((user) => user.id),
        });
    }

    updateTodoType(todoType: TodoType): Promise<TodoType> {
        return this.put<TodoType>(`todoType/${todoType.id}`, todoType);
    }

    updateUser(user: User): Promise<User> {
        return this.put<User>(`user/${user.id}`, user);
    }

}
