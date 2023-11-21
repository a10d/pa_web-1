import {Todo, TodoType, User} from "../index.ts";

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
