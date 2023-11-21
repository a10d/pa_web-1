import {defineStore} from "pinia";
import {Todo, TodoType, useBackend, User} from "../backend";
import {useEventBus} from "../eventBus";


type TodoStoreState = {
    todos: Todo[];
    todoTypes: TodoType[];
    users: User[];
}

const eventBus = useEventBus();

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
            eventBus.emit('fetchUsers');
        },
        async createUser(user: Partial<User>) {
            await useBackend()
                .createUser(user)
                .then((result) => eventBus.emit('createUser', result))
                .then(this.fetchUsers);
        },
        async updateUser(user: User) {
            await useBackend()
                .updateUser(user)
                .then((result) => eventBus.emit('updateUser', result))
                .then(this.fetchUsers);
        },
        async deleteUser(user: User) {
            await useBackend()
                .deleteUser(user)
                .then(() => eventBus.emit('deleteUser', user))
                .then(this.fetchUsers);
        },

        async fetchTodoTypes() {
            this.todoTypes = await useBackend().fetchTodoTypes();
            eventBus.emit('fetchTodoTypes');
        },
        async createTodoType(todoType: Partial<TodoType>) {
            await useBackend()
                .createTodoType(todoType)
                .then((result) => eventBus.emit('createTodoType', result))
                .then(this.fetchTodoTypes);
        },
        async updateTodoType(todoType: TodoType) {
            await useBackend()
                .updateTodoType(todoType)
                .then((result) => eventBus.emit('updateTodoType', result))
                .then(this.fetchTodoTypes);
        },
        async deleteTodoType(todoType: TodoType) {
            await useBackend()
                .deleteTodoType(todoType)
                .then(() => eventBus.emit('deleteTodoType', todoType))
                .then(this.fetchTodoTypes);
        },
        todoTypeById(id: string): TodoType {
            return this.todoTypes.find((i) => i.id === id) ?? {
                id,
                name: "Unbekannt",
                description: "Unbekannter Aufgabentyp",
                reminderTime: 0,
                color: "#000000",
            }
        },

        async fetchTodos() {
            this.todos = await useBackend().fetchTodos();
            eventBus.emit('fetchTodos');
        },
        async createTodo(todo: Partial<Todo>) {
            await useBackend()
                .createTodo(todo)
                .then((result) => eventBus.emit('createTodo', result))
                .then(this.fetchTodos);
        },
        async updateTodo(todo: Todo) {
            await useBackend()
                .updateTodo(todo)
                .then((result) => eventBus.emit('updateTodo', result))
                .then(this.fetchTodos);
        },
        async deleteTodo(todo: Todo) {
            await useBackend()
                .deleteTodo(todo)
                .then(() => eventBus.emit('deleteTodo', todo))
                .then(this.fetchTodos);
        }
    },
});
