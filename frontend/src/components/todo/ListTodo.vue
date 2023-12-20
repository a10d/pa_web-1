<script lang="ts" setup>

import {Todo, TodoType, User} from "../../services/backend";
import {computed} from "vue";
import {useTodoStore} from "../../services/store/todo.ts";
import {useEventBus} from "../../services/eventBus";
import UserAvatarList from "../ui/UserAvatarList.vue";
import {useMoment} from "../../services/moment";

type ListTodoProps = {
    todo: Todo,
};

const props = defineProps<ListTodoProps>();

const emit = defineEmits<{
    editTodo: [todo: Todo],
}>();

const moment = useMoment();
const store = useTodoStore();
const eventBus = useEventBus();

const checkBoxProxy = computed({
    get: () => props.todo.completed,
    set: (completed: boolean) => {
        if (completed === props.todo.completed) return;

        eventBus.emit('playSound', completed ? 'todoCompleted' : 'todoUncompleted')

        store.updateTodo({
            ...props.todo,
            completed,
        });
    }
});

const todoType = computed<TodoType>(() => store.todoTypeById(props.todo.type));

const assignees = computed<User[]>((): User[] => {
    return props.todo.assignees
        .map(id => store.users.find(user => user.id === id))
        .filter(user => typeof user === 'object') as User[]
});

const dueDateReadable = computed<string>(() => moment(props.todo.dueDate).format('dddd, DD. MMMM YYYY'));
const dueDateRelative = computed<string>(() => moment(props.todo.dueDate).fromNow());
const isDueWarning = computed<boolean>(() => !props.todo.completed && moment(props.todo.dueDate)
    .subtract(todoType.value.reminderTime, 'days').isBefore(moment()));

function editTodo() {
    emit('editTodo', props.todo);
}

</script>

<template>

    <article :class="{'border-warning': isDueWarning,}" class="todo">

        <!-- Checkbox -->
        <div class="w-8 p-2">
            <input v-model="checkBoxProxy"
                   class="border-transparent aspect-square w-4 cursor-pointer"
                   type="checkbox"/>
        </div>

        <div class="cursor-pointer select-none" @click="editTodo">

            <!-- Title -->
            <div class="px-2 py-1">
                <h1 :class="{
                'line-through': todo.completed,
                'text-gray-600': todo.completed,
                'text-warning': isDueWarning,
                'font-bold': isDueWarning,
            }"
                    class="text-lg font-medium"
                    v-text="todo.title"
                />
            </div>


            <!-- Details -->
            <div class="pr-2 flex items-center gap-3">
                <!-- Due Date -->
                <span :class="{'bg-warning text-white font-bold shadow-sm': isDueWarning}"
                      :title="dueDateReadable"
                      class="inline-block text-sm text-gray-600 rounded-full py-1 px-2">
                    Fällig {{ dueDateRelative }}
                </span>

                <!-- Type -->
                <span class="inline-block text-sm">
                      <span :style="{color: todoType.color}" class="mr-1">■</span>
                     <span class="font-medium" v-text="todoType.name"/>
                </span>

                <!-- Assignees -->
                <UserAvatarList :users="assignees"/>
            </div>


            <!-- Description -->
            <div class="p-2">
                <p :class="{
                'line-through': todo.completed,
                'text-gray-600': todo.completed,
            }"
                   v-text="todo.description"/>
            </div>

        </div>
    </article>

</template>

<style scoped>

.todo {
    @apply flex border rounded-lg mb-4 shadow-sm;
}


.border-warning {
    border-color: v-bind('todoType.color');
}

.bg-warning {
    background-color: v-bind('todoType.color');
}

.text-warning {
    color: v-bind('todoType.color');
}

</style>
