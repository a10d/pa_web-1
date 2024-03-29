import {ServerConnector} from "./Connectors/ServerConnector.ts";
import {ArrayConnector} from "./Connectors/ArrayConnector.ts";
import {LocalStorageConnector} from "./Connectors/LocalStorageConnector.ts";

export type User = {
    id: string;
    firstName: string;
    lastName: string;
    email: string;
};

export type TodoType = {
    id: string;
    name: string;
    description: string;
    color: string;
    reminderTime: number;
}

export type Todo = {
    id: string;
    type: TodoType;
    title: string;
    description: string | null;
    dueDate: Date;
    assignees: User[];
    completed: boolean;
};

export class ValidationError extends Error {

    name = "ValidationError";

    messages: {
        [key: string]: string[];
    };

    constructor(messages: { [key: string]: string[] } | undefined) {
        super("Validation error occurred.");
        this.messages = messages ?? {};
    }

    has(key: string): boolean {
        return this.messages.hasOwnProperty(key);
    }

    get(key: string): string[] {
        return this.messages[key] ?? [];
    }

    all(): { [key: string]: string[] } {
        return this.messages;
    }
}

export interface Connector {
    createUser(user: Partial<User>): Promise<User>;

    updateUser(user: User): Promise<User>;

    deleteUser(user: User): Promise<boolean>;

    fetchUsers(): Promise<User[]>;

    fetchTodos(): Promise<Todo[]>;

    createTodo(todo: Partial<Todo>): Promise<Todo>;

    updateTodo(todo: Todo): Promise<Todo>;

    deleteTodo(todo: Todo): Promise<boolean>;

    fetchTodoTypes(): Promise<TodoType[]>;

    createTodoType(todoType: Partial<TodoType>): Promise<TodoType>;

    updateTodoType(todoType: TodoType): Promise<TodoType>;

    deleteTodoType(todoType: TodoType): Promise<boolean>;

    createExport(): Promise<string>;

    importData(data: string): Promise<boolean>;
}


const client = function (): Connector {
    switch (import.meta.env.VITE_BACKEND_CONNECTOR ?? "array") {
        case "localstorage":
            return new LocalStorageConnector();
        case "server":
            return new ServerConnector(import.meta.env.VITE_BACKEND_URL ?? "http://localhost:8080");
        case "array":
        default:
            return new ArrayConnector();
    }
}();

export const useBackend = (): Connector => client;
