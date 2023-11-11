import {defineStore} from "pinia";
import {Todo, TodoType, useBackend, User} from "../backend";


type TodoStoreState = {
    todos: Todo[];
    todoTypes: TodoType[];
    users: User[];
}

export const useTodoStore = defineStore("todo", {
    state: (): TodoStoreState => ({
        todos: [],
        todoTypes: [],
        users: [],
    }),
    actions: {
        async initialize() {
            await Promise.all([
                this.fetchTodos(),
                this.fetchTodoTypes(),
                this.fetchUsers(),
            ]);
        },

        async fetchUsers() {
            this.users = await useBackend().fetchUsers();
        },
        async createUser(user: Partial<User>) {
            await useBackend().createUser(user).then(this.fetchUsers);
        },
        async updateUser(user: User) {
            await useBackend().updateUser(user).then(this.fetchUsers);
        },
        async deleteUser(user: User) {
            await useBackend().deleteUser(user).then(this.fetchUsers);
        },

        async fetchTodoTypes() {
            this.todoTypes = await useBackend().fetchTodoTypes();
        },
        async createTodoType(todoType: Partial<TodoType>) {
            await useBackend().createTodoType(todoType).then(this.fetchTodoTypes);
        },
        async updateTodoType(todoType: TodoType) {
            await useBackend().updateTodoType(todoType).then(this.fetchTodoTypes);
        },
        async deleteTodoType(todoType: TodoType) {
            await useBackend().deleteTodoType(todoType).then(this.fetchTodoTypes);
        },

        async fetchTodos() {
            this.todos = await useBackend().fetchTodos();
        },
        async createTodo(todo: Partial<Todo>) {
            await useBackend().createTodo(todo).then(this.fetchTodos);
        },
        async updateTodo(todo: Todo) {
            await useBackend().updateTodo(todo).then(this.fetchTodos);
        },
        async deleteTodo(todo: Todo) {
            await useBackend().deleteTodo(todo).then(this.fetchTodos);
        }
    },
});
