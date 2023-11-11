import {defineStore} from "pinia";
import {Todo, TodoType, useBackend} from "../backend";
import {useAuthStore} from "./auth";


type TaskStoreState = {
    tasks: Todo[];
    taskTypes: TodoType[];
}

export const useTaskStore = defineStore("task", {
    state: (): TaskStoreState => ({
        tasks: [],
        taskTypes: [],
    }),
    actions: {
        async initialize() {
            await this.fetchTasks();
            await this.fetchTaskTypes();
        },
        async fetchTasks() {
            this.tasks = await useBackend().fetchTasks();
        },
        async fetchTaskTypes() {
            this.taskTypes = await useBackend().fetchTaskTypes();
        },
    },
});
