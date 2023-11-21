import {ArrayConnector} from "./ArrayConnector.ts";


export class LocalStoreConnector extends ArrayConnector {
    protected writeChanges() {
        localStorage.setItem("users", JSON.stringify(this.users));
        localStorage.setItem("todos", JSON.stringify(this.todos));
        localStorage.setItem("todoTypes", JSON.stringify(this.todoTypes));
        this.loadChanges();
    }

    protected loadChanges() {
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
}
