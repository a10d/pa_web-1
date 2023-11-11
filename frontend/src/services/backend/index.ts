import {LocalStoreConnector} from "./Connectors/LocalStoreConnector.ts";

export type User = {
    id: number;
    name: string;
};

export type TodoType = {
    id: number;
    name: string;
    description: string;
    color: string;
    reminderTime: number;
}

export type Todo = {
    id: number;
    type: number;
    title: string;
    description: string | null;
    dueDate: Date;
    assignees: number[];
    completed: boolean;
};

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
}

const client = new LocalStoreConnector();

export const useBackend = (): Connector => client;
