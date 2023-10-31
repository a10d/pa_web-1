export type User = {
    id: number;
    name: string;
}

export type TaskType = {
    id: number;
    name: string;
    description: string;
    color: string;
    reminderTime: number;
}

export type Task = {
    id: number;
    type: number;
    title: string;
    description: string | null;
    dueDate: Date;
    assignees: number[];
    completed: boolean;
};


type LoginResponse = {
    success: boolean;
    error?: string;
}

class Connector {


    async initialize() {

    }

    /**
     * Login with the given credentials.
     */
    async login(credentials: { username: string, password: string }): Promise<LoginResponse> {
        try {

            return {
                success: result.status === 200,
            };
        } catch (e) {
            return {
                success: false,
                error: e.message,
            };
        }
    }

    async fetchUser() {
        return //this.http.get("/api/auth/me");
    }

    async fetchUsers() {
        return //this.http.get("/api/auth/users");
    }


    async fetchTasks() {
        return //this.http.get("/api/tasks/all");
    }

    async createTask() {
        return //this.http.post("/api/tasks/create");
    }

    async updateTask(task: Task) {
        return //this.http.patch(`/api/tasks/update/${ task.id }`);
    }

    async deleteTask(task: Task) {
        return //this.http.delete(`/api/tasks/update/${ task.id }`);
    }

    async fetchTaskTypes() {
        return //this.http.get("/api/taskTypes/all");
    }

    async createTaskType() {
        return //this.http.post("/api/taskTypes/create");
    }

    async updateTaskType(taskType: TaskType) {
        return //this.http.patch(`/api/taskTypes/update/${ taskType.id }`);
    }

    async deleteTaskType(taskType: TaskType) {
        return //this.http.delete(`/api/tasks/update/${ taskType.id }`);
    }


}


const client = new Connector();

client.initialize();

export const useBackend = () => client;
