import mitt from 'mitt'
import {Todo, TodoType, User} from "../backend";

type EventMap = {
    'reload': void;

    'fetchTodos': void;
    'createTodo': Todo;
    'updateTodo': Todo;
    'deleteTodo': Todo;

    'fetchTodoTypes': void;
    'createTodoType': TodoType;
    'updateTodoType': TodoType;
    'deleteTodoType': TodoType;

    'fetchUsers': void;
    'createUser': User;
    'updateUser': User;
    'deleteUser': User;

    'playSound': string;
}

const eventBus = mitt<EventMap>()

export const useEventBus = () => eventBus
