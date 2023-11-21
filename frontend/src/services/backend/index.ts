import {LocalStoreConnector} from "./Connectors/LocalStoreConnector.ts";
import {Connector} from "./Connectors/Connector.ts";

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
    type: string;
    title: string;
    description: string | null;
    dueDate: Date;
    assignees: string[];
    completed: boolean;
};


const client = new LocalStoreConnector();

export const useBackend = (): Connector => client;
