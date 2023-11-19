<script lang="ts" setup>

import {Todo, TodoType} from "../../services/backend";
import {computed} from "vue";
import {useTodoStore} from "../../services/store/todo.ts";
import * as moment from "moment/moment";

type ListTodoProps = {
    todo: Todo,
};

const props = defineProps<ListTodoProps>();

const store = useTodoStore();

const checkBoxProxy = computed({
    get: () => props.todo.completed,
    set: (completed: boolean) => store.updateTodo({
        ...props.todo,
        completed,
    })
});

const todoType = computed<TodoType>(() => store.todoTypeById(props.todo.type));

const dueDateRelative = computed(() => moment(props.todo.dueDate).fromNow());

const isDueWarning = computed(() => {
    if (props.todo.completed) return false;

    const diffTime = new Date(props.todo.dueDate).getTime() - new Date().getTime();
    const diffDays = Math.ceil(diffTime / (1000 * 60 * 60 * 24));

    return diffDays >= todoType.value.reminderTime;
});

</script>

<template>


    <div class="flex items-stretch border-b">
        <!-- checkbox -->
        <div class="w-8 border-r p-2">
            <input v-model="checkBoxProxy"
                   class="border-transparent aspect-square w-4"
                   type="checkbox"/>
        </div>

        <!-- content -->
        <div class="flex-grow">

            <!-- title -->
            <div class="p-1.5 border-b border-dotted">
                <h1 :class="{
                'line-through': todo.completed,
                'text-gray-600': todo.completed,
            }"
                    class="text-lg font-medium font-handwriting"
                    v-text="todo.title"
                />
            </div>

            <!-- description -->
            <div class="p-2">
                <p :class="{
                'line-through': todo.completed,
                'text-gray-600': todo.completed,
            }"
                   class="font-handwriting"
                   v-text="todo.description"/>
            </div>

            <!-- due date -->
            <div class="p-2">
                <p class="text-sm text-gray-600">
                    Fällig {{ dueDateRelative }}
                </p>
            </div>


            <!-- Assignees -->
            <div class="p-2">


                <pre>{{ todo }}</pre>
                <pre>{{ {isDueWarning} }}</pre>
                <pre>{{ {todoType} }}</pre>

            </div>
        </div>
    </div>
</template>
