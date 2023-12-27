import {Connector} from "./Connectors/Connector.ts";
import {ServerConnector} from "./Connectors/ServerConnector.ts";

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

const client = new ServerConnector();

export const useBackend = (): Connector => client;
