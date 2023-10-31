import {defineStore} from "pinia";
import {Task, TaskType, useBackend} from "../backend";
import {useAuthStore} from "./auth";


type TaskStoreState = {
    tasks: Task[];
    taskTypes: TaskType[];
}

export const useTaskStore = defineStore("task", {
    state: (): TaskStoreState => ({
        tasks: [],
        taskTypes: [],
    }),
    actions: {
        async initialize() {
            // Cannot continue if not authenticated
            if (!useAuthStore().isAuthenticated) {
                return;
            }

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
